package process

import (
	"os/exec"
	"testing"
)

func TestKillProcess(t *testing.T) {
	sleep := exec.Command("sleep","5")
	sleep.Start()
	
	err := KillProcess("sleep.exe")
	if err != nil {
		t.Error("Killing process failed:", err)
	}
}
