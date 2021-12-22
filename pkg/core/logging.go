package core

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	color_none   string = "\033[0m"
	color_red    string = "\033[31m"
	color_green  string = "\033[32m"
	color_yellow string = "\033[33m"
	color_blue   string = "\033[34m"
	color_white  string = "\033[37m"
	color_orange string = "\033[38;5;208m"
)

func Log(a ...interface{}) {
	colorMessage("Log", color_green, a...)
}

func Logf(format string, a ...interface{}) {
	colorMessagef("Log", color_green, format, a...)
}

func Logp(format string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Printf(format, a...)
		}
	} else {
		fmt.Printf(format, a...)
	}
}

func Error(a ...interface{}) {
	colorMessage("Error", color_red, a...)
}

func Warn(a ...interface{}) {
	colorMessage("Warning", color_yellow, a...)
}

func Relay(author *discordgo.User, channel string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Println(append([]interface{}{trevor.name + color_blue + " " + author.Username + " (" + channel + "):" + color_none}, a...)...)
		}
	} else {
		fmt.Println(a...)
	}
}

func Info(a ...interface{}) {
	colorMessage("Info", color_orange, a...)
}

func colorMessage(title string, color string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Println(append([]interface{}{trevor.name + color + " " + title + ":" + color_none}, a...)...)
		}
	} else {
		fmt.Println(a...)
	}
}

func colorMessagef(title string, color string, format string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Printf("%s"+": "+format, append([]interface{}{color_white + trevor.name + color + " " + title + color_none}, a...)...)
		}
	} else {
		fmt.Printf(format, a...)
	}
}
