package babel

import (
	"fmt"

	"github.com/fishykins/trevor/pkg/core"
)

func Insult(cmd core.Command) {
	cmd.Reply("Hang on a sec, I'm working on it...", true)

	user := cmd.GetArg("user").IntoDiscordOption().UserValue(cmd.Session)
	adjectives, _ := Dict.RandAdjective(1)
	noun, _ := Dict.RandNoun(1)

	var an string
	adj := adjectives[0].Word

	if adj[0] == 'a' || adj[0] == 'e' || adj[0] == 'i' || adj[0] == 'o' || adj[0] == 'u' {
		an = "an"
	} else {
		an = "a"
	}

	msg := fmt.Sprintf("%s is %s %s %s.", core.UserTag(user.ID), an, adj, noun[0].Word)
	cmd.Message(msg)
	cmd.EditReply(Affirmative())
}
