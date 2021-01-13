package terminal

import (
	"os/exec"
)

func Run(command string) {	
	cmd := exec.Command("/bin/bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}
}

