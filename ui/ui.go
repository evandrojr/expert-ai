package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/ierror"
	"github.com/evandrojr/expert-ai/logic"
	"github.com/evandrojr/expert-ai/os"
)

var ui fyne.App

func TextWindow(text string, title string) {
	newWindow := ui.NewWindow(title)
	newWindow.Resize(fyne.NewSize(400, 250))
	label := widget.NewMultiLineEntry()
	label.Resize(fyne.NewSize(350, 240))
	label.SetText(text)
	label.Wrapping = fyne.TextWrapBreak
	newWindow.SetContent(
		label,
	)
	newWindow.Show()
}

func getPromptTextarea() *widget.Entry {
	promptTextarea := widget.NewMultiLineEntry()
	promptTextarea.SetPlaceHolder("Type a prompt:")
	promptTextarea.SetText(config.Settings.Prompt)
	promptTextarea.Resize(fyne.NewSize(500, 400))
	promptTextarea.Wrapping = fyne.TextWrapBreak
	return promptTextarea
}

func getSettingsContainer() *container.Split {
	settingsTextarea := widget.NewMultiLineEntry()
	settingsTextarea.SetPlaceHolder("Expert AI settings:")
	settings, err := config.GetSettingsString()
	ierror.PanicOnError(err)
	settingsTextarea.SetText(settings)
	settingsTextarea.Resize(fyne.NewSize(500, 400))
	settingsTextarea.Wrapping = fyne.TextWrapBreak

	saveButton := widget.NewButton("Save settings", func() {
		config.WriteSettingsFile(settingsTextarea.Text)
	})

	restoreDefaulstButton := widget.NewButton("Restore defaults", func() {
		config.SaveDefautSettings()
		config.Load()
		settings, err = config.GetSettingsString()
		ierror.PanicOnError(err)
		settingsTextarea.SetText(settings)
	})

	hBoxButtons := container.NewHBox(saveButton, restoreDefaulstButton)
	centeredButtons := container.NewCenter(hBoxButtons)
	split := container.NewVSplit(settingsTextarea, centeredButtons)
	split.SetOffset(.95)
	return split
}

func getPromptContainer() *container.Split {
	promptTextarea := getPromptTextarea()

	submitButton := widget.NewButton("Submit prompt", func() {
		SubmitPrompt(promptTextarea.Text)
	})

	prepareButton := widget.NewButton("Prepare", func() {
		os.PrepareBrowser()
	})
	hBoxButtons := container.NewHBox(submitButton, prepareButton)
	centeredButtons := container.NewCenter(hBoxButtons)
	split := container.NewVSplit(promptTextarea, centeredButtons)
	split.SetOffset(.95)

	return split
}

func Build() {
	ui = app.New()
	mainWindow := ui.NewWindow("Expert AI")
	mainWindow.Resize(fyne.NewSize(800, 600))

	promptContainer := getPromptContainer()
	settingsContainer := getSettingsContainer()

	tabs := container.NewAppTabs(
		container.NewTabItem("Prompt", promptContainer),
		container.NewTabItem("AI models", widget.NewLabel("AI models")),
		container.NewTabItem("Settings", settingsContainer),
	)

	tabs.SetTabLocation(container.TabLocationTrailing)
	mainWindow.SetContent(tabs)
	mainWindow.ShowAndRun()
}

func SubmitPrompt(promptText string) {
	log.Println("Content was:", promptText)
	config.Settings.Prompt = promptText
	config.SaveDefautSettings()
	go logic.Prompt(config.Settings)
	answer := <-logic.AnswerChan
	TextWindow(answer.Answer, answer.Title)
}
