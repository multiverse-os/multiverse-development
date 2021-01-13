package install

import (
	"os/user"
	"strconv"
)

func (self *Installer) SetupMultiverseGroup() error {
	self.AddGroup("multiverse")

	g, err := user.LookupGroup("multiverse")
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(g.guid)
	if err != nil {
		return err
	}

	self.Group = &Group{
		Name: "multiverse",
		ID: gid,
	}
	return nil
}


