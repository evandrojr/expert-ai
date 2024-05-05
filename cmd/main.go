package main

import (
	"os"

	"github.com/evandrojr/expert-ai/internal/config"
	"github.com/evandrojr/expert-ai/internal/filesystem"
	"github.com/evandrojr/expert-ai/internal/ilog"
	"github.com/evandrojr/expert-ai/internal/logic"
	"github.com/evandrojr/expert-ai/internal/ui"
)

func main() {
	config.Init()
	ilog.InitializeLogger(filesystem.JoinPaths(config.ConfigDir, "log.txt"))
	// log.Fatal(errors.New("Que merda"))
	logic.Init()
	ui.Build()
	os.Exit(0)
}
