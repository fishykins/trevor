package bish

import (
	"fmt"

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
			Callback:    play,
		},
		{
			Name:        "move",
			Description: "Set your fish's next move!",
			Callback:    move,
		},
	},
	Buttons: []core.Button{
		{
			CustomID: "bish_hunt",
			Label:    "Hunt",
			Style:    discordgo.DangerButton,
			Callback: setMove,
		},
		{
			CustomID: "bish_swim",
			Label:    "Swim",
			Style:    discordgo.PrimaryButton,
			Callback: setMove,
		},
		{
			CustomID: "bish_sleep",
			Label:    "Sleep",
			Style:    discordgo.SuccessButton,
			Callback: setMove,
		},
		{
			CustomID: "bish_play",
			Label:    "Join",
			Style:    discordgo.SuccessButton,
			Callback: join,
		},
		{
			CustomID: "bish_turn",
			Label:    "Next Round",
			Style:    discordgo.PrimaryButton,
			Callback: next,
		},
	},
	UserCommands: []core.UserCommandType{},
	Ready:        ready,
}

var games map[string]*Game

func ready(a *core.Application) error {
	games = make(map[string]*Game)
	return nil
}

func play(c core.Command) {
	i := c.BuildReply("Game Options")
	core.AddButtonsToInteraction(&i, []string{"bish_play", "bish_turn"})
	c.SendReply(i)
}

func join(c core.Command) {
	channel := c.Interaction.Interaction.ChannelID
	msg := "%s has %s the game!"
	fish := NewFish(c.User)

	if game, ok := games[channel]; ok {
		for _, f := range game.fish {
			if f.Owner.ID == c.User.ID {
				c.Reply("You are already in the game!", true)
				return
			}
		}
		msg = fmt.Sprintf(msg, core.UserTag(c.User.ID), "joined")
		game.AddFish(fish)
	} else {
		channel := discordgo.Channel{
			ID:      channel,
			GuildID: c.Interaction.GuildID,
		}
		game = NewGame(&channel)
		game.AddFish(fish)
		games[channel.ID] = game
		msg = fmt.Sprintf(msg, core.UserTag(c.User.ID), "started")
	}
	c.Reply(msg, false)

}

func next(c core.Command) {
	channel := c.Interaction.Interaction.ChannelID
	if game, ok := games[channel]; ok {
		if len(game.fish) >= 1 {
			if !game.Turn(c.Session) {
				c.Reply("The game has ended!", true)
				games[channel] = nil
			}
		} else {
			c.Reply("There are not enough fish to play!", true)
		}
	} else {
		c.Reply("There is no game currently running on this channel!", true)
	}
}

func move(c core.Command) {
	i := c.BuildReply("What would you like to do?")
	core.AddButtonsToInteraction(&i, []string{"bish_hunt", "bish_swim", "bish_sleep"})
	c.SetPrivate(i)
	c.SendReply(i)
}

func setMove(c core.Command) {
	moveType := c.GetArg("button").IntoDiscordOption().StringValue()
	c.Reply("You have chosen to "+moveType+".", true)
	channel := c.Interaction.Interaction.ChannelID
	if game, ok := games[channel]; ok {
		fish := game.GetFish(c.User.ID)
		if fish != nil {
			value := c.GetArg("button").IntoDiscordOption().StringValue()
			switch value {
			case "bish_hunt":
				fish.Move = Move_Hunt
			case "bish_swim":
				fish.Move = Move_Swim
			case "bish_sleep":
				fish.Move = Move_Sleep
			default:
				fish.Move = Move_None
			}
		}
		game.CheckReady(c.Session)
	}
}
