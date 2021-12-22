package core

import (
	"fmt"
)

var CoreModule Module = Module{
	Name:        "core",
	Description: "Core commands for Trevor.",
	Commands: []CommandType{
		{
			Name:        "ping",
			Description: "Ping the bot to make sure it's alive.",
			Callback: func(c Command) {
				c.Reply("Pong!", true)

			},
		},
	},
	UserCommands: []UserCommandType{
		{
			Name:     "beam",
			Callback: beam,
		},
	},
	Snooper: nil,
}

func beam(c UserCommand) {
	c.Reply("Beaming...", true)
	msg := fmt.Sprintf("%s you just got beamed you fuckin' looser", UserTag(c.Target.ID))
	c.Message(msg)
	// err := c.Session.GuildMemberRoleRemove(c.Interaction.GuildID, c.Target.ID, "905124284022792294")
	// if err != nil {
	// 	Warn(err)
	// }
}
