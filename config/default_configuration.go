package config

import (
	"encoding/json"
	"fmt"

	"github.com/evandrojr/expert-ai/def"
	"github.com/evandrojr/expert-ai/error"
	"github.com/evandrojr/expert-ai/filesystem"
)

type SettingsStruct struct {
	Browser                     string
	KillAndLoadBrowserAtStartUp bool
	PromptChatGpt3_5            bool
	PromptClaude3               bool
	PromptAllAIs                bool
	Prompt                      string
}

var Settings = SettingsStruct{
	Browser:                     "chromium",
	KillAndLoadBrowserAtStartUp: true,
	PromptChatGpt3_5:            true,
	PromptClaude3:               true,
	PromptAllAIs:                false,
	Prompt:                      "",
}

var HomeDir string
var ConfigDir string
var ConfigFile string

func Init() {
	HomeDir, err := filesystem.GetHomeDir()
	error.PanicOnError(err)
	ConfigDir := filesystem.JoinPaths(HomeDir, ".config", def.APP_NAME)
	ConfigFile := filesystem.JoinPaths(ConfigDir, def.CONFIG_FILENAME)
	filesystem.CreateDirectoryIfNotExists(ConfigDir)
	if !filesystem.FileExists(ConfigFile) {
		json, err := json.MarshalIndent(Settings, "", "    ")
		error.PanicOnError(err)
		filesystem.WriteFile(ConfigFile, string(json))
	} else {
		settingsData, err := filesystem.ReadFile(ConfigFile)
		error.PanicOnError(err)
		settingsDataBytes := []byte(settingsData)

		err = json.Unmarshal(settingsDataBytes, &Settings)
		error.PanicOnError(err)
		fmt.Println(Settings)

	}

}
