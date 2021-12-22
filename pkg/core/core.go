package core

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/models"
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
	Ready:   ready,
	Snooper: nil,
}

func ready(a *Application) error {
	a.users = a.Databass().Users()
	return nil
}

func beam(c UserCommand) {
	c.Reply("Beaming...", true)
	msg := fmt.Sprintf("%s you just got beamed you fuckin' looser", UserTag(c.Target.ID))
	c.Message(msg)
}

// This should always return a user object, else somthing has gone very wrong.
func GetUser(u *discordgo.User) *models.User {
	userId := GetUserId(u)
	for _, user := range App().users {
		if user.DiscordID == userId {
			return user
		}
	}

	localUser := models.NewUser(userId, 0, u.Username)
	App().users = append(App().users, localUser)
	return localUser
}

// Tries to find a pre-exisitng user, or returns none
func TryGetUser(u *discordgo.User) (*models.User, error) {
	userId := GetUserId(u)
	for _, user := range App().users {
		if user.DiscordID == userId {
			return user, nil
		}
	}
	return nil, nil
}

// Pushes user to database with any newly set values.
func UpdateUser(u *models.User) error {
	ctx := context.Background()
	return App().Databass().UpdateUser(&ctx, u)
}
