package ui

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/evandrojr/expert-ai/error"
	"github.com/evandrojr/expert-ai/filesystem"
	"github.com/evandrojr/expert-ai/tool"
)

var data = []string{"a", "string", "list"}

func Build() {
	ui := app.New()
	window := ui.NewWindow("Expert AI")
	prompt, err := filesystem.ReadFile("prompt.txt")
	error.PanicOnError(err)
	// window.SetContent(container.NewHSplit())

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

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
		err := filesystem.WriteFile("prompt.txt", input.Text)
		if err != nil {
			panic(err)
		}
		// doublePrompt("prompt.txt", "answer_poe.txt", "answer_chatgpt.txt", "combined_prompt.txt")
		// doublePrompt("prompts/combined_prompt.txt", "2ndAnswer_poe.txt", "2ndAnswer_chatgpt.txt", "2ndCombined_prompt.txt")

	})
	content := container.NewVBox(list, input, prepareButton, submitButton)
	window.SetContent(content)
	window.ShowAndRun()

}
