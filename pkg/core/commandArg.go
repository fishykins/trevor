package core

import "github.com/bwmarrin/discordgo"

type CommandArgType struct {
	Name        string
	Description string
	Required    bool
	Default     interface{}
	Type        discordgo.ApplicationCommandOptionType
}

type CommandArg struct {
	Name  string
	Value interface{}
	Type  discordgo.ApplicationCommandOptionType
}

func ArgFromDiscordOption(a *discordgo.ApplicationCommandInteractionDataOption) *CommandArg {
	return &CommandArg{
		Name:  a.Name,
		Value: a.Value,
		Type:  a.Type,
	}
}

func (a *CommandArg) IntoDiscordOption() discordgo.ApplicationCommandInteractionDataOption {
	opt := discordgo.ApplicationCommandInteractionDataOption{
		Name:  a.Name,
		Value: a.Value,
		Type:  a.Type,
	}
	return opt
}
