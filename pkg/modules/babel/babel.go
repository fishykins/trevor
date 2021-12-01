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
			},
			Callback: Define,
		},
	},
	Ready: ready,
}

func ready(a *core.Application) error {
	Dict = lang.NewDictionary(a.Databass().EnglishDictionary())
	return nil
}
