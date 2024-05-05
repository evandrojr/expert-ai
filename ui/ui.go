package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/evandrojr/expert-ai/config"
	"github.com/evandrojr/expert-ai/logic"
	"github.com/evandrojr/expert-ai/os"
	"log"
)

func Build() {
	ui := app.New()
	window := ui.NewWindow("Expert AI")
	//window.Resize(fyne.NewSize(500, 400))

	promptTextarea := widget.NewMultiLineEntry()
	promptTextarea.SetPlaceHolder("Type a prompt:")
	promptTextarea.SetText(config.Settings.Prompt)
	promptTextarea.Resize(fyne.NewSize(500, 400))

	tabs := container.NewAppTabs(
		container.NewTabItem("Prompt", promptTextarea),
		container.NewTabItem("AI models", widget.NewLabel("AI models")),
		container.NewTabItem("Settings", widget.NewLabel("Settings")),
	)
	//tabs.MinSize()

	submitButton := widget.NewButton("Submit prompt", func() {
		SubmitPrompt(promptTextarea.Text)
		// doublePrompt("prompt.txt", "answer_poe.txt", "answer_chatgpt.txt", "combined_prompt.txt")
		// doublePrompt("prompts/combined_prompt.txt", "2ndAnswer_poe.txt", "2ndAnswer_chatgpt.txt", "2ndCombined_prompt.txt")

	})

	prepareButton := widget.NewButton("Prepare", func() {
		// go  func(){
		os.PrepareBrowser()
		// }

	})
	bottonSplit := container.NewHBox(submitButton, prepareButton)
	main := container.NewVSplit(tabs, bottonSplit)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationTrailing)

	window.SetContent(main)
	window.ShowAndRun()
	//window.FullScreen()
	////prompt, err := filesystem.ReadFile("prompt.txt")
	////error.PanicOnError(err)
	//// window.SetContent(container.NewHSplit())
	//
	//list := widget.NewList(
	//	func() int {
	//		return len(data)
	//	},
	//	func() fyne.CanvasObject {
	//		return widget.NewLabel("template")
	//	},
	//	func(i widget.ListItemID, o fyne.CanvasObject) {
	//		o.(*widget.Label).SetText(data[i])
	//	})
	//
	//split := container.NewHSplit(list, container.NewMax())
	//tabs.'' = 0.2

	//

	//

	//split := container.NewVBox(list, promptTextarea, prepareButton, submitButton)

}

func SubmitPrompt(promptText string) {
	log.Println("Content was:", promptText)
	config.Settings.Prompt = promptText
	config.Save()
	go logic.Prompt(config.Settings)
}
