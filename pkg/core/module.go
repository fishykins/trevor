package core

import "github.com/bwmarrin/discordgo"

type Module struct {
	Name        string
	Description string
	Commands    []CommandType
	// Function to call on application init. Avoid Print calls.
	Init func() error
	// Function to call on application start (post init). Avoid Print calls.
	Ready func(a *Application) error
	// Function to call on application exit. Avoid Print calls.
	Stop func(a *Application) error
	// Function to call on all messages. Useful for AI assisted learning, moderating and other such things.
	Snooper func(s *discordgo.Session, m *discordgo.MessageCreate)
}
