package process

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
)

func KillProcess(name string) {
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println("Unable to retrieve list of active processes.")
		return
	}
	for _, process := range processes {
		if process.Executable() == name {
			pid := process.Pid()
			proc, err := os.FindProcess(pid)
			if err == nil {
				if os.Getpid() != pid {
					proc.Kill()
				}
			} else {
				fmt.Println("no process named '%v' found", name)
			}
		}
	}
}
