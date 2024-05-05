package main

import (
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/logic"
	"github.com/evandrojr/expert-ai/ui"
	"os"
)

func main() {
	config.Init()
	logic.Init()
	ui.Build()
	os.Exit(0)
}
