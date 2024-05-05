package logic

import (
	"errors"

	artificialintelligence "github.com/evandrojr/expert-ai/internal/artificial_intelligence"
	"github.com/evandrojr/expert-ai/internal/config"
	"github.com/evandrojr/expert-ai/internal/filesystem"
	"github.com/evandrojr/expert-ai/internal/ierror"
	"github.com/evandrojr/expert-ai/internal/os"
)

var AnswerChan chan AnswerStruct

type AnswerStruct struct {
	Answer string
	Title  string
	Error  error
}

func Init() {
	AnswerChan = make(chan AnswerStruct)
	StartBrowserIfNeed()
}

func sendPrompt(ai artificialintelligence.ArtificialIntelligence, prompt string) (string, error) {
	if !os.IsProcessRunning(config.Settings.Browser) {
		return "", errors.New(`Launch browser with --remote-debugging-port=9222 or press "Launch browser" button`)
	}

	var _ai = ai.Setup()
	answer, err := _ai.SubmitPrompt(prompt)
	if err != nil {
		return "", err
	}
	// ierror.PanicOnError(err)
	return answer, nil
}

func RunClaudeIfRequired(settings config.SettingsStruct) {
	if settings.PromptClaude3 {
		var claude3 artificialintelligence.Claude3
		answerClaude, err := sendPrompt(claude3, settings.Prompt)
		if err != nil {
			AnswerChan <- AnswerStruct{Error: err}
			return
		}
		err = filesystem.WriteFile(filesystem.JoinPaths(config.AnswersDir, "claude3.txt"), answerClaude)
		if err != nil {
			AnswerChan <- AnswerStruct{Error: err}

			return
		}
		ierror.PanicOnError(err)
	} else {
		AnswerChan <- AnswerStruct{Error: errors.New("norun")}
	}
}

func RunChatGptIfRequired(settings config.SettingsStruct) {
	if settings.PromptChatGpt3_5 {
		var chatgpt artificialintelligence.Chatgpt
		answerChatgpt, err := sendPrompt(chatgpt, settings.Prompt)
		if err != nil {
			AnswerChan <- AnswerStruct{Error: err}
			return
		}
		err = filesystem.WriteFile(filesystem.JoinPaths(config.AnswersDir, "ChatGPT3.5.txt"), answerChatgpt)
		if err != nil {
			AnswerChan <- AnswerStruct{Error: err}
			return
		}
		AnswerChan <- AnswerStruct{Answer: answerChatgpt, Title: "Answer ChatGPT 3.5"}
	} else {
		AnswerChan <- AnswerStruct{Error: errors.New("norun")}
	}
}

func Prompt(settings config.SettingsStruct) {

	//headerForAnswer1 := "Resposta 1:\n\n"
	//headerForAnswer2 := "\n\nResposta 2:\n\n"
	go RunClaudeIfRequired(settings)
	go RunChatGptIfRequired(settings)

	//basedOntheQuestion := `Baseado na pergunta:
	//
	//`
	//
	//compareAnswers, err := filesystem.ReadFile("prompts/" + "compare_answers.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//combinedPromptText := basedOntheQuestion + prompt + "\n\n" +
	//	compareAnswers +
	//	headerForAnswer1 + answerClaude +
	//	headerForAnswer2 + answerChatgpt
	//
	//err = filesystem.WriteFile("prompts/"+combinedPrompt, combinedPromptText)
	//if err != nil {
	//	panic(err)
	//}
	//// filesystem.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}

func StartBrowserIfNeed() {
	isBrowserRunning := os.IsProcessRunning(config.Settings.Browser)
	if !isBrowserRunning {
		os.RunBrowser()
	}
}
