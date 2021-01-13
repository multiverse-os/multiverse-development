package install

import (
	"fmt"

	machine "./machine"
	terminal "./terminal"
	step "./step"
)

///////////////////////////////////////////////////////////////////////////////
type Packages struct {
	Install []string
	Remove []string
}

///////////////////////////////////////////////////////////////////////////////
func To(m machine.Type) (installer Context) {
	switch m {
	case machine.Controller:
	case machine.Host:
		installer = Context{
			Machine: machine.Init(),
			Packages: Packages{},
			//Step: PrepareSystem,
			Paths: Paths{
				HomePath: "/home/user",
				GitPath: "/var/multiverse/development",
			},
		}
	default:
		panic(fmt.Errorf("invalid machine type"))
	}
	return installer
}

///////////////////////////////////////////////////////////////////////////////
func (self *Context) Start() (err error) {
	if !IsRoot() { panic(fmt.Errorf("must be root")) }


	//// User
	terminal.Output("Setting up user...")
	step.AskRetry(self.SetupUser)

	//// Packages

	terminal.Output("Removing unnecessary packages......")
	// TODO best way to call this with AskRetry. Global with list of packages?
	//step.AskRetry(pm.Autoremove)

	terminal.Output("Creating default filepath....")
	step.AskRetry(self.CreateMultiversePaths)

	//// Configurations
	//// Install Config files
	terminal.Output("Downloading default config files....")
	step.AskRetry(self.CloneGitRepository)

	terminal.Output("Copying default config files....")
	step.AskRetry(self.CopyGeneralConfigFiles)

	//// Enable IOMMU in grub
	terminal.Output("Copying processor specific config files and enabling IOMMU in grub....")
	step.AskRetry(self.DoProcessorSpecificConfig)

	terminal.Output("Adding modules to initramfs....")
	step.AskRetry(self.DoInitramfsConfig)

	return err
}


func (self *Context) DoProcessorSpecificConfig() error {
	switch self.Machine.CPU.Vendor { 
	case "AuthenticAMD":
		// TODO: These don't work
		Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-amd"), "/etc/default/grub")
		Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-amd"), "/etc/modules")
	case "GenuineIntel":
		Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-intel"), "/etc/default/grub")
		Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-intel"), "/etc/modules")
	}
	terminal.Run("update-grub")
	return nil
}

func (self *Context) DoInitramfsConfig() error {
	Copy(self.Paths.BaseFile(machine.Host, "/etc/initramfs-tools/modules"), "/etc/initramfs-tools/modules")
	terminal.Run("update-initramfs -u")
	return nil
}

