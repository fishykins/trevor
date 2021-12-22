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
			Name:        "set_token",
			Description: "Add a token to your profile",
			Args: []core.CommandArgType{
				{
					Name:        "key",
					Description: "key of the token",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "value",
					Description: "value of the token",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
			Callback: SetPlayerToken,
		},
		{
			Name:        "get_token",
			Description: "Gets a token value",
			Args: []core.CommandArgType{
				{
					Name:        "key",
					Description: "key of the token",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
			Callback: GetPlayerToken,
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
	Ready:   ready,
	Stop:    nil,
	Snooper: snoop,
}

func ready(a *core.Application) error {
	Users = core.App().Databass().Users()
	return nil
}

func snoop(s *discordgo.Session, m *discordgo.MessageCreate) {
	user, err := core.TryGetUser(m.Author)
	if err != nil || user != nil {
		HandleMessage(user, m.Content)
		return
	}
	if user == nil {
		user = core.GetUser(m.Author)
		HandleMessage(user, m.Content)
		core.UpdateUser(user)
	}

}

func RemoveUser(cmd core.Command) {
	ctx := context.Background()
	user := core.GetUser(cmd.User)
	core.App().Databass().RemoveUser(&ctx, user)
	cmd.Reply("Removed user from database.", true)
	// TODO: Remove user from local cache as well.
}

func SetPlayerToken(cmd core.Command) {
	key := cmd.GetArg("key").IntoDiscordOption().StringValue()
	value := cmd.GetArg("value").IntoDiscordOption().StringValue()
	user := core.GetUser(cmd.User)
	if value != "" {
		user.Tokens[key] = value
		cmd.Reply("Successfully set token.", true)
	} else {
		delete(user.Tokens, key)
		cmd.Reply("Successfully removed token.", true)
	}
	core.UpdateUser(user)
}

func GetPlayerToken(cmd core.Command) {
	key := cmd.GetArg("key").IntoDiscordOption().StringValue()
	user := core.GetUser(cmd.User)
	value := user.Tokens[key]
	if value != "" {
		cmd.Reply("Token value: "+value, true)
	} else {
		cmd.Reply("Token not found.", true)
	}
}
