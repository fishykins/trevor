package main

import (
	"os"
	"os/signal"

	"github.com/fishykins/trevor/internal/core"
	"github.com/fishykins/trevor/internal/modules/chatter"
)

func main() {
	core.Start([]core.Module{chatter.Chatter})
	defer core.Stop()

	end := make(chan os.Signal)
	signal.Notify(end, os.Interrupt)
	<-end
}
