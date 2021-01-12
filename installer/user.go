package install

type User struct {
	Name string

	Path string

	UID int
	GID int

	Groups []string
}

///////////////////////////////////////////////////////////////////////////////
func IsRoot() bool {
	u, err = user.Current()
	if err != nil {
		return false
	}
	return (u.Uid != "0")

}


///////////////////////////////////////////////////////////////////////////////
func (self *Installer) CreateGroup(name string) error {
	return Terminal(fmt.Sprintf("groupadd --system %s", name))
}


///////////////////////////////////////////////////////////////////////////////
func (self *Installer) SetUserAsOwner(path string) error { return os.Chown(path, self.User.uid, self.User.gid) }


///////////////////////////////////////////////////////////////////////////////
func (self User) AddToGroup(group string) error {
	return Terminal(fmt.Sprtinf("usermod -a -G %s %s", group, self.Name))
}

///////////////////////////////////////////////////////////////////////////////
func (self *Installer) SetupUser() (err error) {
	// TODO: Need to create the user 'user', but it at this step we 
	//       dont have a real installer,  so this will be created 
	//       during the debian installation.

	u, err = user.Lookup("user")
	if err != nil {
		return err
	}

	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)

	self.User = User{
		Name: "user",
		UID: uid,
		GID: gid,
	}

	err = self.RemoveDirectory(self.Paths.Home("/Desktop"))
	err = self.RemoveDirectory(self.Paths.Home("/Downloads"))
	err = self.RemoveDirectory(self.Paths.Home("/Documents"))
	err = self.RemoveDirectory(self.Paths.Home("/Music"))
	err = self.RemoveDirectory(self.Paths.Home("/Videos"))
	err = self.RemoveDirectory(self.Paths.Home("/Pictures"))


	err = self.User.AddToGroup("kvm")
	err = self.user.AddToGroup("libvirt")

	return nil
}
