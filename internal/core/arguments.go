package core

import (
	"github.com/bwmarrin/discordgo"
)

func GetArgument(i *discordgo.InteractionCreate, command string, arg string) *discordgo.ApplicationCommandInteractionDataOption {
	data := i.ApplicationCommandData()
	if data.Name == command {
		for _, option := range data.Options {
			if option.Name == arg {
				return option
			}
		}
	}

	for _, option := range data.Options {
		if option.Type == discordgo.ApplicationCommandOptionSubCommand {
			for _, subOption := range option.Options {
				if subOption.Name == arg {
					return subOption
				}
			}
		}
	}
	return nil
}
