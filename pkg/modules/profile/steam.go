package profile

import (
	"context"
	"fmt"

	"github.com/fishykins/trevor/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
)

const SID_LENGTH = 17

func SetSteamId(command core.Command) {
	ctx := context.Background()
	usersCollection := core.App().Databass().Users()

	discordID := command.Interaction.User.ID
	//steamID := command.GetArg("id").IntoDiscordOption().IntValue()

	cursor, err := usersCollection.Find(ctx, bson.M{"_id": discordID})
	if err != nil {
		core.Error(err)
		command.Reply("I can't find you in my databass, which is my bad. Tell Fishy to fix this...", true)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			core.Error(err)
			msg := fmt.Sprintf("Oh shit, this is an error yo: %s\n", err.Error())
			command.Reply(msg, true)
		}
		fmt.Println(user)
		command.Reply("I got you bro", true)
	}

}
