package main

import (
	install "../.."
	machine "../../machine"
)

func main() {
	installer := install.To(machine.Host)

	installer.Start()
}
