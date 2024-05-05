package os

import "testing"

func Test_isProcessRunning(t *testing.T) {
	got, _ := IsProcessRunning("chromexx")
	println(got)
}
