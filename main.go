package main

import (
	"os"

	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/ilog"
	"github.com/evandrojr/expert-ai/logic"
	"github.com/evandrojr/expert-ai/ui"
)

func main() {
	config.Init()
	ilog.InitializeLogger(filesystem.JoinPaths(config.ConfigDir, "log.txt"))
	// log.Fatal(errors.New("Que merda"))
	logic.Init()
	ui.Build()
	os.Exit(0)
}
