package core

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func UserTag(id string) string {
	return fmt.Sprintf("<@%s>", id)
}

func BangTag(id string) string {
	return fmt.Sprintf("<@!%s>", id)
}

func GetUserId(u *discordgo.User) uint64 {
	userId, err := strconv.ParseUint(u.ID, 10, 64)
	if err != nil {
		return 0
	}
	return userId
}
