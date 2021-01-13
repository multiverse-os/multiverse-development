package install

import (
	machine "./machine"
	packagemanager "./machine/packagemanager"
)


type Context struct {
	PackageManager packagemanager.PackageManager

	Machine *machine.Machine

	Step *InstallStep

	User *User
	Group *Group

	Paths Paths

	Packages Packages


	RequiredNewPaths []string
}
