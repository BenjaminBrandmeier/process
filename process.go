package process

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
)

func KillProcess(name string) error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of processes:", err)
	}
	for _, process := range processes {
		if process.Executable() == name {
			pid := process.Pid()
			proc, _ := os.FindProcess(pid)
			if os.Getpid() != pid {
				err := proc.Kill()
				if err != nil {
					return fmt.Errorf("Unable to kill process:", err)
				}
				return nil
			}
		}
	}
	return fmt.Errorf("Unable to kill process")
}
