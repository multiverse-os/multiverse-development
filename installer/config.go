package install

import (
	"os"
	"fmt"
	"path/filepath"
	
	machine "./machine"
)


func (self *Installer) CopyGeneralConfigFiles() (err error) {
	Copy(self.Paths.BaseFile(machine.Host, self.Paths.Home(".gitconfig"), self.Paths.Home(".gitconfig"))

	self.SetUserAsOwner(self.Paths.Home("/.gitconfig"))

	CreateDirectory("/etc/multiverse")

	self.InstallFile("/etc/motd")
	self.InstallFile("/etc/issue")
	self.InstallFile("/etc/security/limits.conf")
	self.InstallFile("/etc/sysctl.conf")
	self.InstallFile("/etc/sysctl.d/30-tracker.conf")
	self.InstallFile("/etc/sysctl.d/99-sysctl.conf")
	self.InstallFile("/etc/rc.local")

	return nil
}

func (self *Installer) InstallFile(path string) error {
	return Copy(self.Paths.BaseFile(machine.Host, path), path)
}

func FileOrDirectoryExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("os.Stat ... !os.IsNotExist(err)")
		return true
	}
	return false
}

func (self *Installer) CloneGitRepository() error {
	// TODO: Better than erroring if the directory is already there is checking 
	//       the git error and cd + git pull instead.
	if FileOrDirectoryExists(self.Paths.GitPath) {
		fmt.Println("FileOrDirectoryExists('path')...")
		fmt.Println("self.Paths.GitPath", self.Paths.GitPath)
fmt.Printf("cd %s && git pull", self.Paths.GitPath)
return nil
	}else{
		return Terminal(fmt.Sprintf("git clone https://github.com/multiverse-os/multiverse-development %s", self.Paths.GitPath))
	}
}

func (self *Installer) ParseConfigFiles() (files []string, err error) {
	return filepath.Glob(self.Paths.BaseFiles(machine.Host))
}
