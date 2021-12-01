package core

import "github.com/bwmarrin/discordgo"

// Wrapper type for discordgo.ApplicationCommandOption
type CommandType struct {
	Name        string
	Description string
	Args        []CommandArgType
	Callback    func(Command)
}

func (c *CommandType) IntoDiscordCommand() discordgo.ApplicationCommandOption {
	opt := discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Options:     []*discordgo.ApplicationCommandOption{},
		Type:        discordgo.ApplicationCommandOptionSubCommand,
	}

	for _, arg := range c.Args {
		cmd := &discordgo.ApplicationCommandOption{
			Name:        arg.Name,
			Description: arg.Description,
			Required:    arg.Required,
			Options:     []*discordgo.ApplicationCommandOption{},
			Type:        arg.Type,
		}
		opt.Options = append(opt.Options, cmd)
	}
	return opt
}
