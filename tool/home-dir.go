package tool

import (
	"os/user"
)

type homeDir struct{}

func (h *homeDir) Get() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}

func GetHomeDir() string {
	homeDir := &homeDir{}
	dir, err := homeDir.Get()
	if err != nil {
		LogFatal("Error getting homedir:" + err.Error())
	}

	return dir
}
