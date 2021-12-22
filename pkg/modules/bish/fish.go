package bish

import "github.com/fishykins/trevor/pkg/models"

type MoveType int

const startingEnergy int = 5

const (
	Move_None MoveType = iota
	Move_Hunt
	Move_Eat
	Move_Swim
	Move_Sleep
)

type Fish struct {
	Owner  *models.User
	Name   string
	Age    int
	Energy int
	Size   int
	Move   MoveType
}

func (f *Fish) New(owner *models.User, name string, size int) {
	f.Owner = owner
	f.Name = name
	f.Age = 0
	f.Energy = startingEnergy
	f.Size = size
	f.Move = Move_None
}
