package core

type Module struct {
	Name        string
	Description string
	Commands    []DiscordCommand
	Start       func() error
	Stop        func() error
}
