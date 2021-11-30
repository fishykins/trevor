package main

import (
	"os"
	"os/signal"

	"github.com/fishykins/trevor/pkg/core"
)

func main() {
	core.InitApplication("exampleApp", []core.Module{}, true)
	core.StartApplication()
	defer core.StopApplication()

	end := make(chan os.Signal)
	signal.Notify(end, os.Interrupt)
	<-end
}
