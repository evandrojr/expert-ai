package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	artificialintelligence "github.com/evandrojr/expert-ai/artificial_intelligence"
	"github.com/evandrojr/expert-ai/tool"
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
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return
	}
	tool.CreateDirectoryIfNotExists(dir + "/answers")

	ui()
}

func ui() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Type a prompt:")
	prompt, err := tool.ReadFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Type a prompt:")
	input.SetText(prompt)

	prepareButton := widget.NewButton("Prepare", func() {
		// go  func(){
		tool.PrepareBrowser()
		// }

	})

	submitButton := widget.NewButton("Submit prompt", func() {
		log.Println("Content was:", input.Text)
		err := tool.WriteFile("prompt.txt", input.Text)
		if err != nil {
			panic(err)
		}
		doublePrompt("prompt.txt", "answer_poe.txt", "answer_chatgpt.txt", "combined_prompt.txt")
		doublePrompt("prompts/combined_prompt.txt", "2ndAnswer_poe.txt", "2ndAnswer_chatgpt.txt", "2ndCombined_prompt.txt")

	})
	content := container.NewVBox(input, prepareButton, submitButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func doublePrompt(promptFile, answer1File, answer2File, combinedPrompt string) {

	prompt, err := tool.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}
	headerForAnswer1 := "Resposta 1:\n\n"
	headerForAnswer2 := "\n\nResposta 2:\n\n"

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	err = tool.WriteFile("answers/"+answer1File, answerClaude)
	if err != nil {
		panic(err)
	}

	var chatgpt artificialintelligence.Chatgpt
	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = tool.WriteFile("answers/"+answer2File, answerChatgpt)
	if err != nil {
		panic(err)
	}

	basedOntheQuestion := `Baseado na pergunta:
	
	`

	compareAnswers, err := tool.ReadFile("prompts/" + "compare_answers.txt")
	if err != nil {
		panic(err)
	}

	combinedPromptText := basedOntheQuestion + prompt + "\n\n" +
		compareAnswers +
		headerForAnswer1 + answerClaude +
		headerForAnswer2 + answerChatgpt

	err = tool.WriteFile("prompts/"+combinedPrompt, combinedPromptText)
	if err != nil {
		panic(err)
	}
	// tool.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}
