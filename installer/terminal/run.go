package terminal

func Run(cmd string) {	
	cmd := exec.Command("/bin/bash", "-c", cmd)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}
}
