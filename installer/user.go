package install

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	terminal "./terminal"
)


type User struct {
	Name string

	Path string

	UID int
	GID int

	Groups []string
}

///////////////////////////////////////////////////////////////////////////////
func IsRoot() bool {
	u, err := user.Current()
	if err != nil {
		return false
	}
	return (u.Uid == "0")

}




///////////////////////////////////////////////////////////////////////////////
func (self *Context) SetUserAsOwner(path string) error { return os.Chown(path, self.User.UID, self.User.GID) }


///////////////////////////////////////////////////////////////////////////////
func (self User) AddToGroup(group string) error {
	return terminal.Run(fmt.Sprintf("usermod -a -G %s %s", group, self.Name))
}

///////////////////////////////////////////////////////////////////////////////
func (self *Context) SetupUser() error {
	// TODO: Need to create the user 'user', but it at this step we 
	//       dont have a real installer,  so this will be created 
	//       during the debian installation.

	u, err := user.Lookup("user")
	if err != nil {
		return err
	}

	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)

	self.User = &User{
		Name: "user",
		UID: uid,
		GID: gid,
	}

	os.Remove(self.Paths.Home("/Desktop"))
	os.Remove(self.Paths.Home("/Downloads"))
	os.Remove(self.Paths.Home("/Documents"))
	os.Remove(self.Paths.Home("/Music"))
	os.Remove(self.Paths.Home("/Videos"))
	os.Remove(self.Paths.Home("/Pictures"))


	self.CreateGroup("libvirt")
	self.CreateGroup("multiverse")

	self.User.AddToGroup("kvm")
	self.User.AddToGroup("libvirt")
	self.User.AddToGroup("multiverse")

	return nil
}
