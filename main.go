package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	// "fyne.io/fyne/widget"
	artificialintelligence "github.com/evandrojr/expert-ai/artificial_intelligence"
	"github.com/evandrojr/expert-ai/tools"
	// "fyne.io/fyne/v2/widget"
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
	fmt.Println("killall chrome")
	fmt.Println("google-chrome --remote-debugging-port=9222")
	ui()
}

func ui() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Type a prompt:")
	prompt, err := tools.ReadFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Type a prompt:")
	input.SetText(prompt)

	prepareButton := widget.NewButton("Prepare", func() {
		// go  func(){
		tools.PrepareChrome()
		// }

	})

	submitButton := widget.NewButton("Submit prompt", func() {
		log.Println("Content was:", input.Text)
		err := tools.WriteFile("prompt.txt", input.Text)
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

	prompt, err := tools.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}
	headerForAnswer1 := "Resposta 1:\n\n"
	headerForAnswer2 := "\n\nResposta 2:\n\n"

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	err = tools.WriteFile("answers/"+answer1File, answerClaude)
	if err != nil {
		panic(err)
	}

	var chatgpt artificialintelligence.Chatgpt
	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = tools.WriteFile("answers/"+answer2File, answerChatgpt)
	if err != nil {
		panic(err)
	}

	basedOntheQuestion := `Baseado na pergunta:
	
	`

	compareAnswers, err := tools.ReadFile("prompts/" + "compare_answers.txt")
	if err != nil {
		panic(err)
	}

	combinedPromptText := basedOntheQuestion + prompt + "\n\n" +
		compareAnswers +
		headerForAnswer1 + answerClaude +
		headerForAnswer2 + answerChatgpt

	err = tools.WriteFile("prompts/"+combinedPrompt, combinedPromptText)
	if err != nil {
		panic(err)
	}
	// tools.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}
