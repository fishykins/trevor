package main

import (
	"os"
	"os/signal"

	"github.com/fishykins/trevor/pkg/core"
	"github.com/fishykins/trevor/pkg/modules/babel"
	"github.com/fishykins/trevor/pkg/modules/profile"
)

func main() {
	core.InitApplication("Trevor", []core.Module{babel.Babel, profile.Profile}, true)
	core.StartApplication()
	defer core.StopApplication()

	end := make(chan os.Signal)
	signal.Notify(end, os.Interrupt)
	<-end
}
