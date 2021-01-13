package install

import (
	"os"
	"fmt"

	machine "./machine"
)

type Paths struct {
	HomePath string
	GitPath string
}

///////////////////////////////////////////////////////////////////////////////
func (self *Context) CreateDirectory(path string) error { 
	return CreateDir(path, 0700, self.User.Id, self.Group.Id)
}

///////////////////////////////////////////////////////////////////////////////
func (self *Context) CreateMultiversePaths() (err error) {
	self.CreateDirectory("/var/multiverse")
	self.CreateDirectory("/etc/multiverse")




	self.CreateDirectory(self.Paths.Var("/portal-gun/os-image"))
	self.CreateDirectory(self.Paths.Var("/portals/share"))
	self.CreateDirectory(self.Paths.Var("/portals"))
	self.CreateDirectory(self.Paths.Var("/portals/sockets"))
	self.CreateDirectory(self.Paths.Var("/portals/sockets/serial"))
	self.CreateDirectory(self.Paths.Var("/portals/sockets/channel"))
	self.CreateDirectory(self.Paths.Var("/portals/sockets/console"))
	self.CreateDirectory(self.Paths.Var("/portals/sockets/parallel"))

	os.Remove(self.Paths.Home("/.local/share/libvirt/images"))

	self.CreateDirecotry(self.Paths.Home("/.local/share/libvirt"))

	os.Symlink(self.Paths.Var("/portals/disks/"), self.Paths.Home("/.local/share/libvirt/images"))
	return err
}

///////////////////////////////////////////////////////////////////////////////
func (self Paths) BaseFile(m machine.Type, path string) string {
	return fmt.Sprintf("%s/%s", self.BaseFiles(m), path)
}

func (self Paths) BaseFiles(m machine.Type) string {
	return fmt.Sprintf("%s/%s/base-files", self.GitPath, m.String())
}

///////////////////////////////////////////////////////////////////////////////
func (self Paths) Home(path string) string {
	return fmt.Sprintf("%s/%s/", self.HomePath, path)
}

func (self Paths) Git(path string) string {
	return fmt.Sprintf("%s/%s/", self.GitPath, path)
}

func (self Paths) Etc(path string) string {
	return fmt.Sprintf("/etc/%s/", self.EtcPath, path)
}

func (self Paths) Var(path string) string {
	return fmt.Sprintf("/var/%s/", self.VarPath, path)
}

///////////////////////////////////////////////////////////////////////////////
func createDirectory(name string, perm os.FileMode, uid int, gid int) error {
	if err := os.MkdirAll(name, perm); err != nil {
		return err
	}

	if err := os.Chown(name, uid, gid); err != nil {
		return err
	}
	return nil
}
