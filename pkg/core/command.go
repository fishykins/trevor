package core

import "github.com/bwmarrin/discordgo"

type Command struct {
	Type        string
	Args        []CommandArg
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	User        *discordgo.User
}

func (c *Command) GetArg(name string) *CommandArg {
	for _, arg := range c.Args {
		if arg.Name == name {
			return &arg
		}
	}
	return nil
}

// Wrapper for sending a simple interaction response.
func (c *Command) Reply(msg string, private bool) {
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

// Builds and returns a suitable reply without sending it.
func (c *Command) BuildReply(msg string) discordgo.InteractionResponse {
	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
}

// Takes and sends a pre-built message.
func (c *Command) SendReply(r discordgo.InteractionResponse) {
	c.Session.InteractionRespond(c.Interaction.Interaction, &r)
	Log("Sent reply: " + r.Data.Content)
}

func (c *Command) SetPrivate(r discordgo.InteractionResponse) {
	r.Data.Flags = 1 << 6
}

// Edits the message that was previously sent for this command.
func (c *Command) EditReply(msg string) {
	c.Session.InteractionResponseEdit(c.Session.State.User.ID, c.Interaction.Interaction, &discordgo.WebhookEdit{
		Content: msg,
	})
}

// Wrapper for sending a message to a spesified channel
func (c *Command) MessageChannel(msg string, channelId string) {
	c.Session.ChannelMessageSend(channelId, msg)
	Log("Sent message: " + msg)
}

// Sends a message in the same channel as the command was sent in.
func (c *Command) Message(msg string) {
	c.MessageChannel(msg, c.Interaction.ChannelID)
	Log("Sent message: " + msg)
}
