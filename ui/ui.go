package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/logic"
	"github.com/evandrojr/expert-ai/os"
)

var ui fyne.App

func TextWindow(text string, title string) {
	newWindow := ui.NewWindow(title)
	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapBreak
	newWindow.SetContent(container.NewVBox(
		label,
	))
	newWindow.Show()
}

func Build() {
	ui = app.New()
	window := ui.NewWindow("Expert AI")
	promptTextarea := widget.NewMultiLineEntry()
	promptTextarea.SetPlaceHolder("Type a prompt:")
	promptTextarea.SetText(config.Settings.Prompt)
	promptTextarea.Resize(fyne.NewSize(500, 400))

	tabs := container.NewAppTabs(
		container.NewTabItem("Prompt", promptTextarea),
		container.NewTabItem("AI models", widget.NewLabel("AI models")),
		container.NewTabItem("Settings", widget.NewLabel("Settings")),
	)

	submitButton := widget.NewButton("Submit prompt", func() {
		SubmitPrompt(promptTextarea.Text)
	})

	prepareButton := widget.NewButton("Prepare", func() {
		os.PrepareBrowser()
	})
	bottonSplit := container.NewHBox(submitButton, prepareButton)
	main := container.NewVSplit(tabs, bottonSplit)
	tabs.SetTabLocation(container.TabLocationTrailing)
	window.SetContent(main)
	window.SetFullScreen(true)
	window.ShowAndRun()
}

func SubmitPrompt(promptText string) {
	log.Println("Content was:", promptText)
	config.Settings.Prompt = promptText
	config.Save()
	go logic.Prompt(config.Settings)
	TextWindow(<-logic.AnswerChan, "Resposta")
}
