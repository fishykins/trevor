package core

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/internal/core/lang"
)

const updateDiscordCommands bool = false

var Modules []Module = []Module{
	{
		Name:        "core",
		Description: "Core commands for Trevor.",
		Commands: []DiscordCommand{
			{
				Command: &discordgo.ApplicationCommandOption{
					Name:        "ping",
					Description: "Ping the bot to make sure it's alive.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				Call: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:   1 << 6,
							Content: "pong!",
						},
					})
				},
			},
		},
	},
}

// Starts the application with the given modules.
func Start(modules []Module) {
	// First, setup the language dictionary.
	err := lang.Open()
	if err != nil {
		fmt.Println("dictionary fatal error:", err)
		os.Exit(1)
	}

	// Setup the discord client.
	DiscordStart()

	// Add modules- These are initiated once the discord client is ready.
	Modules = append(Modules, modules...)

	// Start the event loop
	go Update()
}

func PostStart() {
	// Check commands that discord already has registered.
	cmds, err := GetRegisteredDiscordCommands()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, cmd := range cmds {
			fmt.Printf("Pre-registered slash command found: %v\n", cmd)
		}
	}

	// Register all modules.
	for _, module := range Modules {
		name := module.Name
		fmt.Printf("Registering module commands for '%s'... ", name)
		errs := RegisterModuleCommands(module, updateDiscordCommands)
		if len(errs) > 0 {
			fmt.Printf(" failed: \n")
			for _, err := range errs {
				fmt.Println(err)
			}
		} else {
			fmt.Printf(" ok.\n")
		}
		if module.Start != nil {
			fmt.Printf("Starting module '%s'...", module.Name)
			err := module.Start()
			if err != nil {
				fmt.Printf(" failed: %v\n", err)
			} else {
				fmt.Printf(" ok.\n")
			}
		}
	}
}

func Stop() {
	fmt.Println("Stopping Trevor...")
	lang.Close()
	DiscordStop()
	for _, module := range Modules {
		if module.Stop != nil {
			fmt.Printf("Stopping module '%s'...", module.Name)
			err := module.Stop()
			if err != nil {
				fmt.Printf(" failed: %v\n", err)
			} else {
				fmt.Printf(" ok.\n")
			}
		}
	}
	os.Exit(1)
}

func Update() {
	for {
		Tick()
	}
}

func Tick() {

}
