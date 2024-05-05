package os

import (
	"github.com/evandrojr/expert-ai/config"
	"os/exec"
)

func IsProcessRunning(processName string) bool {
	_, err := exec.Command("pgrep", processName).Output()
	if err != nil {
		return false
	}
	return true
}

func PrepareBrowser() error {
	KillBrowser()
	return RunBrowser()
}

func KillBrowser() {
	exec.Command("killall", config.Settings.Browser).Run()
}

func RunBrowser() error {
	err := exec.Command(config.BrowserExecutableFileName(), "--remote-debugging-port=9222").Start()
	if err != nil {
		return err
	}
	return nil
}