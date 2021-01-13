package install

import (
	"fmt"
	"os/user"
	"strconv"

	terminal "./terminal"
)

type Group struct {
	Name string
	GID int
}


///////////////////////////////////////////////////////////////////////////////
func (self *Context) CreateGroup(name string) error {
	return terminal.Run(fmt.Sprintf("groupadd --system %s", name))
}


///////////////////////////////////////////////////////////////////////////////
func (self *Context) SetupMultiverseGroup() error {
	self.CreateGroup("multiverse")

	g, err := user.LookupGroup("multiverse")
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		return err
	}

	self.Group = &Group{
		Name: "multiverse",
		GID: gid,
	}
	return nil
}


