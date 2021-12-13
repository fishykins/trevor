package profile

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
)

var Profile core.Module = core.Module{
	Name:        "profile",
	Description: "Profile related commands",
	Commands: []core.CommandType{
		{
			Name:        "steam",
			Description: "Set your steam ID",
			Args: []core.CommandArgType{
				{
					Name:        "id",
					Description: "Steam user ID",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
			Callback: nil,
		},
		{
			Name:        "wipe",
			Description: "Wipes all assosiated data with your user profile.",
			Args: []core.CommandArgType{
				{
					Name:        "confirm",
					Description: "Confirm that you want to wipe the rust settings.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionBoolean,
				},
			},
			Callback: nil,
		},
	},
	Ready: nil,
}
