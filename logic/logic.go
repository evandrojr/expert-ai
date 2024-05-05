package logic

import (
	artificialintelligence "github.com/evandrojr/expert-ai/artificial_intelligence"
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/error"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/os"
)

var AnswerChan chan string

func Init() {
	AnswerChan = make(chan string)
	StartBrowserIfNeed()
}

func sendPrompt(ai artificialintelligence.ArtificialIntelligence, prompt string) string {
	var _ai = ai.Setup()
	answer, err := _ai.SubmitPrompt(prompt)
	error.PanicOnError(err)
	return answer
}

func RunClaudeIfRequired(settings config.SettingsStruct) {
	if settings.PromptClaude3 {
		var claude3 artificialintelligence.Claude3
		answerClaude := sendPrompt(claude3, settings.Prompt)
		err := filesystem.WriteFile(filesystem.JoinPaths(config.AnswersDir, "claude3.txt"), answerClaude)
		error.PanicOnError(err)
		// ui.TextWindow(answerClaude, "Answer Claude 3")

	}
}

func RunChatGptIfRequired(settings config.SettingsStruct) {
	if settings.PromptChatGpt3_5 {
		var chatgpt artificialintelligence.Chatgpt
		answerChatgpt := sendPrompt(chatgpt, settings.Prompt)
		err := filesystem.WriteFile(filesystem.JoinPaths(config.AnswersDir, "ChatGPT3.5.txt"), answerChatgpt)
		error.PanicOnError(err)
		// answerChatgpt := "jfkjdjflksdj fds fkçkf adsf ~kljflãsd flsm dsalf al~kfasdlk jdsljf sdjfsdf sa~f dsflkjs ldafjsdl kçfsdjf sdjfkldsj fkjsdlfjl"
		AnswerChan <- answerChatgpt
	} else {
		AnswerChan <- "Sem resposta para o ChatGPT"
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
