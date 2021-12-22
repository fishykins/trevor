package babel

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/lang"
)

var Dict lang.Dictionary

var Babel core.Module = core.Module{
	Name:        "babel",
	Description: "Babel is a simple language tool.",
	Commands: []core.CommandType{
		{
			Name:        "define",
			Description: "Define a word.",
			Args: []core.CommandArgType{
				{
					Name:        "word",
					Description: "The word to define.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "private",
					Description: "If true, responds only to you.",
					Required:    false,
					Default:     false,
					Type:        discordgo.ApplicationCommandOptionBoolean,
				},
			},
			Callback: Define,
		},
		{
			Name:        "insult",
			Description: "Insult a user.",
			Args: []core.CommandArgType{
				{
					Name:        "user",
					Description: "The user to dish it on.",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionUser,
				},
			},
			Callback: Insult,
		},
	},
	UserCommands: []core.UserCommandType{},
	Ready:        ready,
}

func ready(a *core.Application) error {
	Dict = lang.NewDictionary(a.Databass().EnglishDictionary())
	return nil
}
