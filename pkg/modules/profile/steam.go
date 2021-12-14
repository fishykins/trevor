package profile

import (
	"strconv"

	"github.com/fishykins/trevor/pkg/core"
)

const SID_LENGTH = 17

func SetSteamId(cmd core.Command) {
	core.Log("Setting steam id for user " + cmd.User.Username + "...")
	user, err := GetUser(cmd.User)
	if err != nil {
		cmd.Reply("Failed to update steam id: "+err.Error(), true)
		return
	}

	user.Name = cmd.User.Username

	steamIdStr := cmd.GetArg("id").IntoDiscordOption().StringValue()
	steamId, err := strconv.ParseUint(steamIdStr, 10, 64)
	if err == nil {
		user.SteamID = steamId
		cmd.Reply("SteamID set to "+steamIdStr, true)
		UpdateUser(user)
		// TODO: Check the steam id is valid
	} else {
		cmd.Reply("Invalid SteamID: "+err.Error(), true)
	}
}
