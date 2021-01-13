package install

type Context struct {
	PackageManager packagemanager.PackageManager

	Step InstallStep
	ser *User
	Paths Paths
	Packages Packages
	System System
	Paths []string
}
