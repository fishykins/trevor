package bish

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
)

var Bish core.Module = core.Module{
	Name:        "bish",
	Description: "Bish is a game where you play a fish.",
	Commands: []core.CommandType{
		{
			Name:        "play",
			Description: "Play the game!",
			Args: []core.CommandArgType{
				{
					Name:        "channel",
					Description: "The channel to play in.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionChannel,
				},
			},
			Callback: nil,
		},
		{
			Name:        "move",
			Description: "Set your fish's next move!",
			Callback:    Move,
		},
	},
	Buttons: []core.Button{
		{
			CustomID: "hunt",
			Label:    "Hunt",
			Style:    discordgo.DangerButton,
			Callback: SetMove,
		},
		{
			CustomID: "eat",
			Label:    "Eat",
			Style:    discordgo.PrimaryButton,
			Callback: SetMove,
		},
		{
			CustomID: "swim",
			Label:    "Swim",
			Style:    discordgo.SecondaryButton,
			Callback: SetMove,
		},
		{
			CustomID: "sleep",
			Label:    "Sleep",
			Style:    discordgo.SuccessButton,
			Callback: SetMove,
		},
	},
	UserCommands: []core.UserCommandType{},
	Ready:        nil,
}

func Move(c core.Command) {
	i := c.BuildReply("What would you like to do?")
	core.AddButtonsToInteraction(&i, []string{"hunt", "eat", "swim", "sleep"})
	c.SetPrivate(i)
	c.SendReply(i)
}

func SetMove(c core.Command) {
	moveType := c.GetArg("button").IntoDiscordOption().StringValue()
	c.Reply("You have chosen to "+moveType+".", true)
}
