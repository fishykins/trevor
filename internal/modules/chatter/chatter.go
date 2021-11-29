package chatter

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/internal/core"
	"github.com/fishykins/trevor/internal/core/lang"
)

var Chatter core.Module = core.Module{
	Name:        "chatter",
	Description: "Chatter is a module that allows users to chat with each other.",
	Commands: []core.DiscordCommand{
		{
			Command: &discordgo.ApplicationCommandOption{
				Name:        "insult",
				Description: "Insult the given user.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "user",
						Description: "User to insult",
						Required:    true,
						Type:        discordgo.ApplicationCommandOptionUser,
					},
				},
			},
			Call: InsultUser,
		},
		{
			Command: &discordgo.ApplicationCommandOption{
				Name:        "compliment",
				Description: "Compliment the given user.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "user",
						Description: "User to compliment",
						Required:    true,
						Type:        discordgo.ApplicationCommandOptionUser,
					},
				},
			},
			Call: ComplimentUser,
		},
		{
			Command: &discordgo.ApplicationCommandOption{
				Name:        "define",
				Description: "Define a word.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "word",
						Description: "Word to define.",
						Required:    true,
						Type:        discordgo.ApplicationCommandOptionString,
					},
				},
			},
			Call: Define,
		},
	},
	Start: func() error { return nil },
	Stop:  func() error { return nil },
}

func InsultUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "hold tight, Im whipping up a zinger...",
		},
	})

	userId := core.GetArgument(i, "insult", "user").UserValue(nil).ID
	insult, _ := lang.Insult(core.UserTag(userId))
	s.ChannelMessageSend(i.ChannelID, insult)
}

func ComplimentUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "hold tight, I'm thinking about what to say...",
		},
	})

	userId := core.GetArgument(i, "compliment", "user").UserValue(nil).ID
	insult, _ := lang.Insult(core.UserTag(userId))
	s.ChannelMessageSend(i.ChannelID, insult)
}

func Define(s *discordgo.Session, i *discordgo.InteractionCreate) {
	str := core.GetArgument(i, "define", "word").StringValue()
	word, err := lang.LookupWord(str)

	if err != nil {
		fmt.Println(err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("I'm sorry, I don't know the word \"%s\" :(", str),
			},
		})
		return
	}
	msg := fmt.Sprintf("**%s** - *%s* : %s", word.Word, word.WordType, word.Description)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})

}
