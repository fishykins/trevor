package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/internal/core/lang"
)

var Client *discordgo.Session
var CommandCallbacks map[string]CallbackFunc

type CallbackFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

func init() {
	var err error
	token := os.Getenv("TREVOR")
	Client, err = discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	CommandCallbacks = make(map[string]CallbackFunc)
}

func DiscordStart() {
	registerEventhandlers()

	err := Client.Open()
	if err != nil {
		panic(err)
	}
}

func DiscordStop() {
	Client.Close()
}

func GetRegisteredDiscordCommands() ([]*discordgo.ApplicationCommand, error) {
	commands, err := Client.ApplicationCommands(Client.State.User.ID, "")
	if err != nil {
		return nil, err
	}
	return commands, nil
}

func RegisterModuleCommands(module Module, push bool) []error {
	errors := make([]error, 0)

	// Build the master command, or "group"
	group := discordgo.ApplicationCommand{
		Name:        module.Name,
		Description: module.Description,
		Options:     []*discordgo.ApplicationCommandOption{},
	}

	// Iter through all commands and add them as options to the group.
	for _, command := range module.Commands {
		name := module.Name + "_" + command.Command.Name
		cmd := command.Command
		group.Options = append(group.Options, cmd)
		CommandCallbacks[name] = command.Call
		fmt.Printf("registered callback for command '%s'. ", name)
	}

	if push {
		_, err := Client.ApplicationCommandCreate(Client.State.User.ID, "", &group)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func registerEventhandlers() {
	Client.AddHandler(commandEvent)
	Client.AddHandler(clientReady)
	Client.AddHandler(messageEvent)
	Client.AddHandler(guildEvent)
}

// Called when the client connects to Discord.
func clientReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Discord client connected")
	PostStart()
}

func commandEvent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var commandId string
	var callback CallbackFunc

	// Find the callback for the command and resulting arguments.
	commandId = i.ApplicationCommandData().Name
	if cb, ok := CommandCallbacks[commandId]; ok {
		callback = cb
	} else {
		commandId = i.ApplicationCommandData().Name + "_" + i.ApplicationCommandData().Options[0].Name
		if cb, ok := CommandCallbacks[commandId]; ok {
			callback = cb
		}
	}

	if callback != nil {
		fmt.Println("Executing callback for", commandId, "...")
		callback(s, i)
		return
	} else {
		// Respond with an error message.
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "I seem to have failed to register a callback for this command. Please shout at Fishy if you want this fixed.",
			},
		})

		// Lets dump all the data to the console for debugging purposes.
		fmt.Printf("Command \"%s\" failed to parse...\n", i.ApplicationCommandData().Name)
		for _, opt := range i.ApplicationCommandData().Options[0].Options {
			fmt.Printf("\t%s: %s\n", opt.Name, opt.Value)
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println(m.Content)

	// check if the message is "!airhorn"
	if strings.HasPrefix(m.Content, "insult") {

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

		// Respond
		msg, err := lang.Insult("you")
		if err != nil {
			fmt.Println(err)
		} else {
			s.ChannelMessageSend(m.ChannelID, msg)
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildEvent(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}
}

type DiscordCommand struct {
	Command *discordgo.ApplicationCommandOption
	Call    CallbackFunc
}
