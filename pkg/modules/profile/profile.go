package profile

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/models"
)

var Users []*models.User

// The profile module handles indivudal user data. It has a basic handle on discord and steam ids built in.
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
	Users = GetUsersFromMongo()
	return nil
}

func GetId(u *discordgo.User) uint64 {
	userId, err := strconv.ParseUint(u.ID, 10, 64)
	if err != nil {
		return 0
	}
	return userId
}

// This should always return a user object, else somthing has gone very wrong.
func GetUser(u *discordgo.User) (*models.User, error) {
	userId := GetId(u)
	for _, user := range Users {
		if user.DiscordID == userId {
			return user, nil
		}
	}

	localUser := &models.User{
		DiscordID: userId,
		Name:      u.Username,
		SteamID:   0,
	}

	Users = append(Users, localUser)
	return localUser, nil
}

func UpdateUser(u *models.User) error {
	ctx := context.Background()
	return UpdateMongoUser(&ctx, u)
}

func RemoveUser(cmd core.Command) {
	ctx := context.Background()
	user, err := GetUser(cmd.User)
	if err != nil {
		cmd.Reply("Failed to remove user from database- contact admin for support.", true)
		return
	}
	RemoveMongoUser(&ctx, user)
	cmd.Reply("Removed user from database.", true)
	// TODO: Remove user from local cache as well.
}
