package rust

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/lang"
)

var Dict lang.Dictionary

var Babel core.Module = core.Module{
	Name:        "rust",
	Description: "Handler for in-game interactions in Rust.",
	Commands: []core.CommandType{
		{
			Name:        "sw",
			Description: "Toggle a switch.",
			Args: []core.CommandArgType{
				{
					Name:        "switch",
					Description: "The name of the switch.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "state",
					Description: "what to do with it.",
					Required:    true,
					Default:     2,
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "on",
							Value: 1,
						},
						{
							Name:  "off",
							Value: 0,
						},
						{
							Name:  "toggle",
							Value: 2,
						},
					},
				},
				{
					Name:        "delay",
					Description: "How many seconds to wait before running the switch command.",
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionInteger,
					Default:     0,
				},
			},
			Callback: nil,
		},
		{
			Name:        "addDevice",
			Description: "Add a device to the list of devices that can be controlled.",
			Args: []core.CommandArgType{
				{
					Name:        "id",
					Description: "The ID of the device.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "name",
					Description: "The name you wish to allocate to this device",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "wipe",
			Description: "Wipe the rust settings.",
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
