package os

import "testing"

func Test_isProcessRunning(t *testing.T) {
	got := IsProcessRunning("chromexx")
	println(got)
}
