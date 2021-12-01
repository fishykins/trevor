package babel

import (
	"fmt"

	"github.com/fishykins/trevor/pkg/core"
)

func Define(cmd core.Command) {
	arg := cmd.GetArg("word").Value.(string)
	cmd.Reply(fmt.Sprintf("sit tight, I'm looking up the definition of **%s**...", arg), true)

	word, _ := Dict.LookupWord(arg)

	if word != nil {
		if word.Description != "" {
			cmd.Message(fmt.Sprintf("**%s** <*%s*>\n%s", word.Inner, word.Type, word.Description))
		} else {
			cmd.EditReply(fmt.Sprintf("While I can find \"%s\" in my dictionary, it doesn't seem to have a defenition attached yet. sorry!", arg))
		}
	} else {
		cmd.EditReply(fmt.Sprintf("No entry for the word \"%s\" found, sorry.", arg))
	}
}
