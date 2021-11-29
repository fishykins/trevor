package main

import (
	"fmt"

	"github.com/fishykins/trevor/tools"
)

func main() {
	fmt.Println("Wiping discord bot commands...")
	tools.DiscordClean()
	fmt.Println("Done!")
}
