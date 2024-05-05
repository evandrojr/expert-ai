package ui

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/evandrojr/expert-ai/internal/config"
	"github.com/evandrojr/expert-ai/internal/filesystem"
	"github.com/evandrojr/expert-ai/internal/ierror"
	"github.com/evandrojr/expert-ai/internal/logic"
	"github.com/evandrojr/expert-ai/internal/os"
)

var ui fyne.App

func ShowErrorWindow(error string) {
	newWindow := ui.NewWindow("error")
	newWindow.Resize(fyne.NewSize(400, 250))
	// errorIcon := theme.NewPrimaryThemedResource(theme.ErrorIcon())
	label := widget.NewLabel(error)
	label.Resize(fyne.NewSize(350, 240))
	label.Wrapping = fyne.TextWrapBreak
	closeButton := widget.NewButton("Close", func() {
		newWindow.Close()
	})
	vbox := container.NewVBox(label, closeButton)
	newWindow.SetContent(
		vbox,
	)
	newWindow.Show()
}

func TextWindow(text string, title string) {
	newWindow := ui.NewWindow(title)
	newWindow.Resize(fyne.NewSize(400, 250))
	entry := widget.NewMultiLineEntry()
	entry.Resize(fyne.NewSize(350, 240))
	entry.SetText(text)
	entry.Wrapping = fyne.TextWrapBreak
	newWindow.SetContent(
		entry,
	)
	newWindow.Show()
}

func ShowProgressWindow(mensage string, finished chan bool) fyne.Window {
	newWindow := ui.NewWindow("Wait")
	newWindow.CenterOnScreen()

	progress := widget.NewProgressBar()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 100)
			progress.SetValue(i)
		}
		finished <- true
	}()
	label := widget.NewLabel(mensage)
	newWindow.SetContent(container.NewVBox(label, progress))
	newWindow.Show()
	return newWindow
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
		config.Load()
		finishedChan := make(chan bool)
		progressWindow := ShowProgressWindow("Saving settings", finishedChan)
		finished := <-finishedChan
		if finished {
			progressWindow.Close()
		}
	})

	restoreDefaulstButton := widget.NewButton("Restore defaults", func() {
		config.SaveDefautSettings()
		config.Load()
		settings, err = config.GetSettingsString()
		ierror.PanicOnError(err)
		finishedChan := make(chan bool)
		progressWindow := ShowProgressWindow("Restoring defaults", finishedChan)
		finished := <-finishedChan
		if finished {
			progressWindow.Close()
		}
		settingsTextarea.SetText(settings)
	})

	hBoxButtons := container.NewHBox(saveButton, restoreDefaulstButton)
	centeredButtons := container.NewCenter(hBoxButtons)
	split := container.NewVSplit(settingsTextarea, centeredButtons)
	split.SetOffset(.95)
	return split
}

func getLogContainer() fyne.Widget {
	logsTextarea := widget.NewMultiLineEntry()
	logsTextarea.SetPlaceHolder("Application logs should appear here:")
	logs, err := filesystem.ReadFile(filesystem.JoinPaths(config.ConfigDir, "log.txt"))
	if err != nil {
		logs = ""
	}
	// logs = " Ola bebe"
	logsTextarea.SetText(logs)
	logsTextarea.Resize(fyne.NewSize(500, 400))
	// logsTextarea.Wrapping = fyne.TextWrapBreak
	// logsTextarea.Disable()

	// saveButton := widget.NewButton("Save settings", func() {
	// 	config.WriteSettingsFile(settingsTextarea.Text)
	// 	finishedChan := make(chan bool)
	// 	progressWindow := ShowProgressWindow("Saving settings", finishedChan)
	// 	finished := <-finishedChan
	// 	if finished {
	// 		progressWindow.Close()
	// 	}
	// })

	// restoreDefaulstButton := widget.NewButton("Restore defaults", func() {
	// 	config.SaveDefautSettings()
	// 	config.Load()
	// 	settings, err = config.GetSettingsString()
	// 	ierror.PanicOnError(err)
	// 	finishedChan := make(chan bool)
	// 	progressWindow := ShowProgressWindow("Restoring defaults", finishedChan)
	// 	finished := <-finishedChan
	// 	if finished {
	// 		progressWindow.Close()
	// 	}
	// 	settingsTextarea.SetText(settings)
	// })

	// hBoxButtons := container.NewHBox(saveButton, restoreDefaulstButton)
	// centeredButtons := container.NewCenter(hBoxButtons)

	// split := container.NewVSplit(settingsTextarea, centeredButtons)
	// split.SetOffset(.95)
	return logsTextarea
}

func getPromptContainer() *container.Split {
	promptTextarea := getPromptTextarea()

	submitButton := widget.NewButton("Submit prompt", func() {
		SubmitPrompt(promptTextarea.Text)
	})

	prepareButton := widget.NewButton("Launch browser", func() {
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
	mainWindow.CenterOnScreen()

	tabs := container.NewAppTabs(
		container.NewTabItem("Prompt", getPromptContainer()),
		container.NewTabItem("AI models", widget.NewLabel("AI models")),
		container.NewTabItem("Settings", getSettingsContainer()),
		container.NewTabItem("Logs", getLogContainer()),
	)

	tabs.SetTabLocation(container.TabLocationTrailing)
	mainWindow.SetContent(tabs)
	mainWindow.ShowAndRun()
	// mainWindow.RequestFocus()
}

func SubmitPrompt(promptText string) {
	log.Println("Content was:", promptText)
	config.Settings.Prompt = promptText
	config.SaveSettings()
	go logic.Prompt(config.Settings)
	answer, ok := <-logic.AnswerChan
	if ok {
		if answer.Error != nil {
			ShowErrorWindow(answer.Error.Error())
		} else {
			TextWindow(answer.Answer, answer.Title)
		}
	}
	answer, ok = <-logic.AnswerChan
	if ok {
		if answer.Error != nil {
			ShowErrorWindow(answer.Error.Error())
		} else {
			TextWindow(answer.Answer, answer.Title)
		}
	}
}
