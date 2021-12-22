package profile

import (
	"strconv"

	"github.com/fishykins/trevor/pkg/core"
)

const SID_LENGTH = 17

func SetSteamId(cmd core.Command) {
	core.Log("Setting steam id for user " + cmd.User.Username + "...")
	user := core.GetUser(cmd.User)

	user.Name = cmd.User.Username
	user.Tokens["steam"] = "SET"

	steamIdStr := cmd.GetArg("id").IntoDiscordOption().StringValue()
	steamId, err := strconv.ParseUint(steamIdStr, 10, 64)
	if err == nil {
		user.SteamID = steamId
		cmd.Reply("SteamID set to "+steamIdStr, true)
		core.UpdateUser(user)
		// TODO: Check the steam id is valid
	} else {
		cmd.Reply("Invalid SteamID: "+err.Error(), true)
	}
}
