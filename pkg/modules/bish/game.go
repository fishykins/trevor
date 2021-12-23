package bish

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fishykins/trevor/pkg/core"
)

const cost_hunt int = 3
const cost_swim int = 0
const cost_sleep int = 1

type Game struct {
	channel *discordgo.Channel
	fish    []*Fish
	turn    int
	Timeout int
	Running bool
}

func NewGame(channel *discordgo.Channel) *Game {
	return &Game{
		channel: channel,
		fish:    make([]*Fish, 0),
		turn:    0,
		Timeout: 2,
		Running: false,
	}
}

func (g *Game) AddFish(f *Fish) {
	g.fish = append(g.fish, f)
}

func (g *Game) GetFish(user string) *Fish {
	for _, f := range g.fish {
		if f.Owner.ID == user {
			return f
		}
	}
	return nil
}

func (g *Game) CheckReady(s *discordgo.Session) {
	ready := 0
	if len(g.fish) <= 1 {
		return
	}
	for _, f := range g.fish {
		if f.Move != Move_None {
			ready++
		}
	}
	if ready == len(g.fish) {
		g.Turn(s)
	}
}

func (g *Game) Turn(s *discordgo.Session) bool {
	hunters := make([]*Fish, 0)
	sleepers := make([]*Fish, 0)
	swimmers := make([]*Fish, 0)

	alive := 0
	output := fmt.Sprintf("*Turn %d...*\n", g.turn)

	rand.Seed(time.Now().UnixNano())

	for _, f := range g.fish {
		if f.Energy < 0 {
			// Dead fish
			continue
		}
		alive++
		switch f.Move {
		case Move_None:
			f.Move = RandomMove()
		case Move_Hunt: // Hunt
			hunters = append(hunters, f)
		case Move_Swim: // Swim
			swimmers = append(swimmers, f)
		case Move_Sleep:
			sleepers = append(sleepers, f)
		}
	}

	// Handle hunters and sleepers
	if len(hunters) > len(sleepers) {
		// Hunters loose
		for _, f := range hunters {
			f.Energy -= cost_hunt
		}
		for _, f := range sleepers {
			f.Energy += cost_sleep
		}
	} else {
		// Hunters win
		for _, hunter := range hunters {
			i := rand.Intn(len(sleepers))
			sleeper := sleepers[i]
			sleepers = append(sleepers[:i], sleepers[i+1:]...)
			hunter.Energy += sleeper.Energy
			output += fmt.Sprintf("%s has eaten %s!\n", hunter.Owner.Username, sleeper.Owner.Username)
			sleeper.Energy = -1
		}
	}

	for _, f := range swimmers {
		f.Energy -= cost_swim
	}

	for _, f := range g.fish {
		if f.Energy >= 0 {
			if alive == 1 {
				output = fmt.Sprintf("%s is the winner!", f.Owner.Username)
				g.end(f)
				g.Running = false
				break
			} else {
				output += fmt.Sprintf("%s has %d energy left.\n", f.Owner.Username, f.Energy)
			}
		} else {
			output += fmt.Sprintf("%s is dead.\n", f.Owner.Username)
		}
		f.Move = Move_None
	}
	s.ChannelMessageSend(g.channel.ID, output)
	g.turn++
	return alive > 1
}

func (g *Game) end(f *Fish) {
	user := core.GetUser(f.Owner)
	user.Stats["bishWins"]++
	core.UpdateUser(user)
}
