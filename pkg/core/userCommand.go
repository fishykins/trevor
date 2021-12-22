package core

import (
	"github.com/bwmarrin/discordgo"
)

type UserCommandType struct {
	Name     string
	Callback func(UserCommand)
}

type UserCommand struct {
	Name        string
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	Caller      *discordgo.User
	Target      *discordgo.User
}

func (u *UserCommandType) IntoDiscordCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name: u.Name,
		Type: discordgo.UserApplicationCommand,
	}
}

// Wrapper for simple interaction responses.
func (c *UserCommand) Reply(msg string, private bool) {
	r := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}

	if private {
		r.Data.Flags = 1 << 6
	}
	c.Session.InteractionRespond(c.Interaction.Interaction, &r)
	Log("Sent reply: " + msg)
}

// Wrapper for sending a message to a spesified channel
func (c *UserCommand) MessageChannel(msg string, channelId string) {
	c.Session.ChannelMessageSend(channelId, msg)
	Log("Sent message: " + msg)
}

// Sends a message in the same channel as the command was sent in.
func (c *UserCommand) Message(msg string) {
	c.MessageChannel(msg, c.Interaction.ChannelID)
}
