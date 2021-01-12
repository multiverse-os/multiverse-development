package install

import (
	"fmt"
	"path/filepath"
	
	machine "./machine"
)


func (self *Installer) CopyGeneralConfigFiles() (err error) {
	err = Copy(self.Paths.BaseFile(machine.Host, "/home/user/.gitconfig"), self.Paths.Home("/.gitconfig"))

	err = self.ChangeOwnerToUser(self.Paths.Home("/.gitconfig"))

	err = self.CreateDirectory("/etc/multiverse")

	err = self.InstallFile("/etc/motd")
	err = self.InstallFile("/etc/issue")
	err = self.InstallFile("/etc/security/limits.conf")
	err = self.InstallFile("/etc/sysctl.conf")
	err = self.InstallFile("/etc/sysctl.d/30-tracker.conf")
	err = self.InstallFile("/etc/sysctl.d/99-sysctl.conf")
	err = self.InstallFile("/etc/rc.local")

	return nil
}

func (self *Installer) InstallFile(path string) error {
	return Copy(self.Paths.BaseFile(machine.Host, path), path)
}

func (self *Installer) CloneGitRepository() error {
	// TODO: Better than erroring if the directory is already there is checking 
	//       the git error and cd + git pull instead.
	return Terminal(fmt.Sprintf("git clone https://github.com/multiverse-os/multiverse-development %s", self.Paths.GitPath))
}

func (self *Installer) ParseConfigFiles() (files []string, err error) {
	return filepath.Glob(self.Paths.BaseFiles(machine.Host))
}
