package install

import (
	packagemanager "./machine/packagemanager"
)


type Context struct {
	PackageManager packagemanager.PackageManager

	Machine *Machine

	Step *InstallStep

	User *User
	Group *Group

	Paths Paths

	Packages Packages


	RequiredNewPaths []string
}
