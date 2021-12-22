package bish

import "github.com/bwmarrin/discordgo"

type Game struct {
	Channel *discordgo.Channel
	Fish    []Fish
	Turn    int
	Timeout int
}
