package bish

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MoveType int

const startingEnergy int = 2

const (
	Move_None MoveType = iota
	Move_Hunt
	Move_Swim
	Move_Sleep
)

type Fish struct {
	Owner  *discordgo.User
	Age    int
	Energy int
	Move   MoveType
}

func NewFish(owner *discordgo.User) *Fish {
	return &Fish{
		Owner:  owner,
		Age:    0,
		Energy: startingEnergy,
		Move:   Move_None,
	}
}

func RandomMove() MoveType {
	rand.Seed(time.Now().UnixNano())
	return MoveType(rand.Intn(3))
}
