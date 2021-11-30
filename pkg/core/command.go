package core

import "github.com/bwmarrin/discordgo"

type Command struct {
	Type           string
	Args           []CommandArg
	DiscordSession *discordgo.Session
	DiscordMessage *discordgo.MessageCreate
}
