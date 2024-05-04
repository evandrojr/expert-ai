package setup

import (
	"github.com/evandrojr/expert-ai/def"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/tool"
)

func CreateConfigDirectory() {
	homeDir, err := filesystem.GetHomeDir()
	if err != nil {
		tool.LogFatal(err.Error())
	}
	configDir := filesystem.JoinPaths(homeDir, def.APP_NAME)
	filesystem.CreateDirectoryIfNotExists(configDir)
}
