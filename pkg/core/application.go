package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PUSH bool = false

// The main entry point for the application. There should only be one instance of this being handled at a time!
type Application struct {
	name     string
	modules  []Module
	databass *Databass
	session  *discordgo.Session
	commands map[string]CommandType
	Debug    bool
}

var trevor *Application

// Getter for the currently running application.
func App() *Application {
	if trevor == nil {
		panic("No application has been initialized!")
	}
	return trevor
}

// Initialize the application and all its submodules. Be aware that submoudles can fail their initialization without causing the application init to fail.
func InitApplication(name string, modules []Module, debug bool) (*Application, []error) {
	if trevor != nil {
		return nil, []error{fmt.Errorf("Application already initialized")}
	}

	var errs []error = make([]error, 0)

	// Base modules + modules passed in
	modules = append([]Module{
		{
			Name:        "core",
			Description: "Core commands for Trevor.",
			Commands: []CommandType{
				{
					Name:        "ping",
					Description: "Ping the bot to make sure it's alive.",
					Callback:    func(c Command) {},
				},
			},
		},
	}, modules...)

	// Initialize database connection using enviromenent variables
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		// This is probably a fatal error- return and let the caller handle it
		errs = append(errs, err)
		if debug {
			Log("Failed to initialize database:", err)
			Log("Application failed to initialize- returning to caller with error(s).")
		}
		return nil, errs
	}

	// Initialize discord session using environment variables
	session, err := discordgo.New("Bot " + os.Getenv("TREVOR"))
	if err != nil {
		errs = append(errs, err)
		if debug {
			Log("Failed to initialize discord:", err)
			Log("Application failed to initialize- returning to caller with error(s).")
		}
		return nil, errs
	}

	// Register discord events, such as 'ready' and message handlers.
	session.AddHandler(clientReady)
	session.AddHandler(messageEvent)
	session.AddHandler(commandEvent)
	session.AddHandler(guildEvent)

	// Build the application
	trevor = &Application{
		name:     name,
		modules:  modules,
		databass: &Databass{client: mongoClient},
		session:  session,
		Debug:    debug,
		commands: make(map[string]CommandType),
	}

	// Init modules
	for _, m := range trevor.modules {
		if m.Init != nil {
			Logf("Initializing module \"%s\"... ", m.Name)
			err := m.Init()
			if err != nil {
				errs = append(errs, err)
				Logp("Failed: \"%s\": %s\n", m.Name, err)
			} else {
				Logp("Ok.\n")
			}
		}
	}

	// Return
	return trevor, errs
}

// Opens all application connections, including database and discord.
// Note: This does NOT excute start functions for modules, that is handled when discord triggers the ready event.
func StartApplication() error {
	if trevor == nil {
		return fmt.Errorf("Application not initialized")
	}

	// Connect to database
	err := trevor.databass.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
	}

	// Connect to discord
	err = trevor.session.Open()
	if err != nil {
		return fmt.Errorf("failed to connect to discord: %s", err)
	}
	return nil
}

// Stops the application and free's all resources. Returns true if successful.
func StopApplication() (bool, []error) {
	if trevor == nil {
		return false, []error{fmt.Errorf("Application not initialized")}
	}

	Log("Stopping application...")
	var errs []error = make([]error, 0)

	// Close database connection
	err := trevor.databass.Disconnect()
	if err != nil {
		errs = append(errs, fmt.Errorf("failed to disconnect from database: %s", err))
		Log("Failed to disconnect from database:", err)
	}

	// Close discord session
	err = trevor.session.Close()
	if err != nil {
		errs = append(errs, fmt.Errorf("failed to close discord session: %s", err))
		Log("Failed to close discord session:", err)
	}

	// Stop modules
	for _, m := range trevor.modules {
		if m.Stop != nil {
			Logf("Stopping module \"%s\"... ", m.Name)
			err := m.Stop()
			if err != nil {
				errs = append(errs, err)
				Logp("Failed: \"%s\": %s\n", m.Name, err)
			} else {
				Logp("ok.\n")
			}
		}
	}
	if len(errs) > 0 {
		Logf("Application stopped with %d errors.", len(errs))
	} else {
		Log("Application stopped safely!")
	}
	trevor = nil
	return true, errs
}

func (a *Application) Name() string {
	return a.name
}

//==============================================================================================================================
//================================================ Discord event handlers ======================================================
//==============================================================================================================================

// Called when the discord client is ready.
func clientReady(s *discordgo.Session, r *discordgo.Ready) {
	Log("Discord client connected")

	trevor := App()

	// Collect all registered discord commands so we can use them later
	commands, _ := trevor.session.ApplicationCommands(trevor.session.State.User.ID, "")

	if len(commands) > 0 {
		Log("Discord commands:")
		for _, c := range commands {
			Logf("\t%s: %s\n", c.Name, c.Description)
			if len(c.Options) > 0 {
				Log("\t\tSub-commands:")
				for _, o := range c.Options {
					Logf("\t\t\t%s: %s\n", o.Name, o.Description)
				}
			}
		}
	}

	// Build all commands and push to discord if needed
	for _, m := range trevor.modules {
		if len(m.Commands) > 0 {
			// Build the root command for this module
			rootCommand := discordgo.ApplicationCommand{
				Name:        m.Name,
				Description: m.Description,
				Options:     []*discordgo.ApplicationCommandOption{},
			}

			for _, c := range m.Commands {
				// Register commands locally for callback lookup
				trevor.commands[c.Name] = c
				// Add command to root command (For exporting to discord)
				cmd := c.IntoDiscordCommand()
				rootCommand.Options = append(rootCommand.Options, &cmd)
			}

			// Push!
			if PUSH {
				_, err := trevor.session.ApplicationCommandCreate(trevor.session.State.User.ID, "", &rootCommand)
				if err != nil {
					Log("Failed to push commands to discord:", err)
				} else {
					Logf("Pushed module group \"%s\" to discord!\n", rootCommand.Name)
				}
			}
		}
	}

	// Execute start functions for modules
	for _, m := range trevor.modules {
		if m.Ready != nil {
			Logf("Starting module \"%s\"... ", m.Name)
			err := m.Ready()
			//TODO: Handle errors
			if err != nil {
				Logp("Failed: \"%s\": %s\n", m.Name, err)
			} else {
				Logp("Ok.\n")
			}
		}
	}
}

func commandEvent(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

// This function will be called every time a new message is created on any channel that the authenticated bot has access to.
func messageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(strings.ToLower(m.Content), BangTag(s.State.User.ID)) {
		taggedMessageEvent(s, m)
	}
}

// Handles messages that start with the bot's mention.
func taggedMessageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Find the channel that the message came from.
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	// Find the guild for that channel.
	_, err = s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}
	msg := fmt.Sprintf("Hello %s, your discord username is %s\n", UserTag(m.Author.ID), m.Author.Username)
	s.ChannelMessageSend(m.ChannelID, msg)
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildEvent(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}
}
