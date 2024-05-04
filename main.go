package main

import (
	"fmt"
	"os"

	artificialintelligence "github.com/evandrojr/expert-ai/artificial_intelligence"
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/ui"
)

func sendPrompt(ai artificialintelligence.ArtificialIntelligence, prompt string) string {
	var _ai = ai.Setup()
	answer, err := _ai.SubmitPrompt(prompt)
	if err != nil {
		panic(err)
	}
	return answer
}

func main() {
	fmt.Println("Run:")
	fmt.Println("killall chromium")
	fmt.Println("chromium --remote-debugging-port=9222")
	config.Init()
	ui.Build()
	os.Exit(0)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return
	}
	filesystem.CreateDirectoryIfNotExists(dir + "/answers")

	// ui()
}

func doublePrompt(promptFile, answer1File, answer2File, combinedPrompt string) {

	prompt, err := filesystem.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}
	headerForAnswer1 := "Resposta 1:\n\n"
	headerForAnswer2 := "\n\nResposta 2:\n\n"

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	err = filesystem.WriteFile("answers/"+answer1File, answerClaude)
	if err != nil {
		panic(err)
	}

	var chatgpt artificialintelligence.Chatgpt
	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = filesystem.WriteFile("answers/"+answer2File, answerChatgpt)
	if err != nil {
		panic(err)
	}

	basedOntheQuestion := `Baseado na pergunta:
	
	`

	compareAnswers, err := filesystem.ReadFile("prompts/" + "compare_answers.txt")
	if err != nil {
		panic(err)
	}

	combinedPromptText := basedOntheQuestion + prompt + "\n\n" +
		compareAnswers +
		headerForAnswer1 + answerClaude +
		headerForAnswer2 + answerChatgpt

	err = filesystem.WriteFile("prompts/"+combinedPrompt, combinedPromptText)
	if err != nil {
		panic(err)
	}
	// filesystem.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}
