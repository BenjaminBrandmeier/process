package process

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
)

func KillProcess(name string) bool {
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, process := range processes {
		if process.Executable() == name {
			pid := process.Pid()
			proc, _ := os.FindProcess(pid)
			if os.Getpid() != pid {
				err := proc.Kill()
				if err != nil {
					fmt.Println(err)
					return false
				}
				return true
			}
		}
	}
	return false
}
