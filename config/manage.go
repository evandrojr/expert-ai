package config

import (
	"encoding/json"
	"errors"
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
	BasedOnTheQuestion          string
	SynthesizeAnswer            string
	AnswerLanguage              string
}

var Settings = SettingsStruct{
	Browser:                     "chromium",
	KillAndLoadBrowserAtStartUp: true,
	PromptChatGpt3_5:            true,
	PromptClaude3:               true,
	PromptAllAIs:                false,
	Prompt:                      "",
	BasedOnTheQuestion:          "Based on the question:",
	SynthesizeAnswer:            "Analyze the two responses and synthesize a single response that captures the main points of both responses. Your response should be concise, clear, and cover the main points of responses to create a correct, and complete response.",
	AnswerLanguage:              "English",
}

var HomeDir string
var ConfigDir string
var ConfigFile string
var AnswersDir string

func Init() {
	err := errors.New("This is a custom error")
	HomeDir, err = filesystem.GetHomeDir()
	error.PanicOnError(err)
	ConfigDir = filesystem.JoinPaths(HomeDir, ".config", def.APP_NAME)
	ConfigFile = filesystem.JoinPaths(ConfigDir, def.CONFIG_FILENAME)
	AnswersDir = filesystem.JoinPaths(ConfigDir, def.ANSWERS_DIR)
	err = filesystem.CreateDirectoryIfNotExists(ConfigDir)
	error.PanicOnError(err)
	err = filesystem.CreateDirectoryIfNotExists(AnswersDir)
	error.PanicOnError(err)
	if filesystem.FileExists(ConfigFile) {
		Load()
	} else {
		Save()
	}

}

func Save() {
	jsonBytes, err := json.MarshalIndent(&Settings, "", "    ")
	error.PanicOnError(err)
	err = filesystem.WriteFile(ConfigFile, string(jsonBytes))
	error.PanicOnError(err)
}

func Load() {
	settingsData, err := filesystem.ReadFile(ConfigFile)
	error.PanicOnError(err)
	settingsDataBytes := []byte(settingsData)
	err = json.Unmarshal(settingsDataBytes, &Settings)
	error.PanicOnError(err)
	fmt.Println(Settings)
}
