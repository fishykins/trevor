package main

import (
	"os"
	"os/signal"

	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/modules/babel"
)

func main() {
	core.InitApplication("exampleApp", []core.Module{babel.Babel}, true)
	core.StartApplication()
	defer core.StopApplication()

	end := make(chan os.Signal)
	signal.Notify(end, os.Interrupt)
	<-end
}
