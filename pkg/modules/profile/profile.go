package profile

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/models"
)

var Users []*models.User

// The profile module handles indivudal user data. It has a basic handle on discord and steam ids built in.
// User data can still be used if this module is not loaded, Profile simply facilitates user caching and a few helper commands.
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
			Callback: SetSteamId,
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
			Callback: RemoveUser,
		},
	},
	Ready: ready,
	Stop:  nil,
}

func ready(a *core.Application) error {
	Users = core.App().Databass().Users()
	return nil
}

// This should always return a user object, else somthing has gone very wrong.
func GetUser(u *discordgo.User) *models.User {
	userId := core.GetUserId(u)
	for _, user := range Users {
		if user.DiscordID == userId {
			return user
		}
	}

	localUser := models.NewUser(userId, 0, u.Username)
	Users = append(Users, localUser)
	return localUser
}

// This should always return a user object, else somthing has gone very wrong.
func TryGetUser(u *discordgo.User) (*models.User, error) {
	userId := core.GetUserId(u)
	for _, user := range Users {
		if user.DiscordID == userId {
			return user, nil
		}
	}
	return nil, nil
}

func UpdateUser(u *models.User) error {
	ctx := context.Background()
	return core.App().Databass().UpdateUser(&ctx, u)
}

func RemoveUser(cmd core.Command) {
	ctx := context.Background()
	user := GetUser(cmd.User)
	core.App().Databass().RemoveUser(&ctx, user)
	cmd.Reply("Removed user from database.", true)
	// TODO: Remove user from local cache as well.
}
