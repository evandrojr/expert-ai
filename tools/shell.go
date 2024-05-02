package tools

import (
	"os/exec"
)

func PrepareChrome() error {
	// Executar o comando "killall chrome"
	err := exec.Command("killall", "chrome").Run()
	if err != nil {
		return err
	}

	// Executar o comando "google-chrome --remote-debugging-port=9222"
	err = exec.Command("google-chrome", "--remote-debugging-port=9222").Start()
	if err != nil {
		return err
	}

	return nil
}
