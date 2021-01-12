package install

type Paths struct {
	HomePath string
	GitPath string
	EtcPath string
	VarPath string
}

///////////////////////////////////////////////////////////////////////////////
func (self *Installer) CreateDirectory(path string) error { return CreateDir(path, 0700, self.User.uid, self.User.gid) }

///////////////////////////////////////////////////////////////////////////////
func (self *Installer) CreateMultiversePaths() (err error) {
	err = self.CreateDirectory(self.Paths.VarPath)
	err = self.CreateDirectory(self.Paths.Var("/portal-gun/os-image"))
	err = self.CreateDirectory(self.Paths.Var("/portals/share"))
	err = self.CreateDirectory(self.Paths.Var("/portals"))
	err = self.CreateDirectory(self.Paths.Var("/portals/sockets"))
	err = self.CreateDirectory(self.Paths.Var("/portals/sockets/serial"))
	err = self.CreateDirectory(self.Paths.Var("/portals/sockets/channel"))
	err = self.CreateDirectory(self.Paths.Var("/portals/sockets/console"))
	err = self.CreateDirectory(self.Paths.Var("/portals/sockets/parallel"))

	err = os.Remove(self.Home("/.local/share/libvirt/images")) 

	err = CreateDir(self.Home("/.local/share/libvirt"), 0755, self.User.UID, self.User.GID)

	err = os.Symlink(self.Var("/portals/disks/"), self.Home("/.local/share/libvirt/images"))
	return err
}

///////////////////////////////////////////////////////////////////////////////
func (self Paths) BaseFile(m machine.Type, path string) string {
	return fmt.Sprintf("%s/%s", self.BaseFiles(m), path)
}

func (self Paths) BaseFiles(m MachineType) string {
	return fmt.Sprintf("%s/%s/base-files", self.GitPath, m.String())
}

///////////////////////////////////////////////////////////////////////////////
func (self Paths) Home(path string) string {
	return fmt.Sprintf("%s/%s", self.HomePath, path)
}

func (self Paths) Git(path string) string {
	return fmt.Sprintf("%s/%s", self.GitPath, path)
}

func (self Paths) Etc(path string) string {
	return fmt.Sprintf("%s/%s", self.EtcPath, path)
}

func (self Paths) Var(path string) string {
	return fmt.Sprintf("%s/%s", self.VarPath, path)
}

///////////////////////////////////////////////////////////////////////////////
