package tools

import (
	"fmt"
	"os/exec"
)

func PrepareBrowser() error {
	// Executar o comando "killall chromium"
	err := exec.Command("killall", "chromium").Run()
	if err != nil {
		fmt.Println(err)
		// return err
	}

	// Executar o comando "google-chromium --remote-debugging-port=9222"
	err = exec.Command("chromium", "--remote-debugging-port=9222").Start()
	if err != nil {
		return err
	}

	return nil
}
