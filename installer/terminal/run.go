package terminal

import (
	"os/exec"
)

func Run(command string) error {	
	cmd := exec.Command("/bin/bash", "-c", command)
	_, err := cmd.Output()
	return err
}

