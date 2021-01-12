package main

import (
	installer "../../../installer"
)

func main() {
	installer := install.To(machine.Host)

	installer.Start()
}
