package config

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/evandrojr/expert-ai/def"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/ierror"
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
	err := errors.New("")
	HomeDir, err = filesystem.GetHomeDir()
	ierror.PanicOnError(err)
	ConfigDir = filesystem.JoinPaths(HomeDir, ".config", def.APP_NAME)
	ConfigFile = filesystem.JoinPaths(ConfigDir, def.CONFIG_FILENAME)
	AnswersDir = filesystem.JoinPaths(ConfigDir, def.ANSWERS_DIR)
	err = filesystem.CreateDirectoryIfNotExists(ConfigDir)
	ierror.PanicOnError(err)
	err = filesystem.CreateDirectoryIfNotExists(AnswersDir)
	ierror.PanicOnError(err)
	if filesystem.FileExists(ConfigFile) {
		Load()
	} else {
		Save()
	}

}

func Save() {
	jsonBytes, err := json.MarshalIndent(&Settings, "", "    ")
	ierror.PanicOnError(err)
	err = filesystem.WriteFile(ConfigFile, string(jsonBytes))
	ierror.PanicOnError(err)
}

func Load() {
	settingsData, err := GetSettingsString()
	ierror.PanicOnError(err)
	settingsDataBytes := []byte(settingsData)
	err = json.Unmarshal(settingsDataBytes, &Settings)
	ierror.PanicOnError(err)
	fmt.Println(Settings)
}

func GetSettingsString() (string, error) {
	return filesystem.ReadFile(ConfigFile)
}

func BrowserExecutableFileName() string {
	if Settings.Browser == "chrome" {
		return "google-chrome"
	}
	return Settings.Browser
}
