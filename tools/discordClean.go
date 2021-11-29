package tools

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DiscordClean() {
	token := os.Getenv("TREVOR")
	c, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	c.Open()
	defer c.Close()

	commands, err := c.ApplicationCommands(c.State.User.ID, "")
	if err != nil {
		panic(err)
	}

	for _, command := range commands {
		// Remove the command from the list
		err = c.ApplicationCommandDelete(c.State.User.ID, "", command.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
