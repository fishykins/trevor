package core

import (
	"github.com/bwmarrin/discordgo"
)

type Button struct {
	CustomID string
	Label    string
	Style    discordgo.ButtonStyle
	Callback func(Command)
	Emoji    string
}

func AddButtonsToInteraction(i *discordgo.InteractionResponse, b []string) {
	actionRow := discordgo.ActionsRow{
		Components: make([]discordgo.MessageComponent, 0),
	}

	for buttonIndex, buttonId := range b {
		if buttonIndex >= 5 {
			Warn("Too many buttons! Only 5 buttons are allowed per row.")
			break
		}
		for key, button := range App().buttonCommands {
			if key == buttonId {
				newButton := discordgo.Button{
					Label:    button.Label,
					Style:    button.Style,
					CustomID: button.CustomID,
				}

				if button.Emoji != "" {
					newButton.Emoji = discordgo.ComponentEmoji{
						Name:     "",
						ID:       button.Emoji,
						Animated: false,
					}
				}
				actionRow.Components = append(actionRow.Components, newButton)
			}
		}
	}
	i.Data.Components = append(i.Data.Components, actionRow)
}
