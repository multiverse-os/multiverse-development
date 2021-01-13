package install

import (
	"fmt"

	machine "./machine"
	terminal "./terminal"
	cpu "./machine/hardware/cpu"

	"github.com/zcalusic/sysinfo"
)


///////////////////////////////////////////////////////////////////////////////
type Packages struct {
	Install []string
	Remove []string
}




///////////////////////////////////////////////////////////////////////////////
func To(m machine.Type) (installer Installer) {
	switch m {
	case machine.Controller:
	case machine.Host:
		installer = Installer{
			Packages: Packages{
				Install: []string{ "ovmf", "virt-manager", "pass", "git", "golang", "ruby", "neovim"},

				Remove: []string{"nano"},
			},
			Step: PrepareSystem,
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
func (self *Installer) Start() (err error) {
	if !IsRoot() { panic(fmt.Errorf("must be root")) }


	//// User
	terminal.Output("Setting up user...")
	AskRetry(self.SetupUser)

	//// Packages

	terminal.Output("Removing unnecessary packages......")
	// TODO best way to call this with AskRetry. Global with list of packages?
	AskRetry(pm.Autoremove)

	terminal.Output("Creating default filepath....")
	AskRetry(self.CreateMultiversePaths)

	//// Configurations
	//// Install Config files
	terminal.Output("Downloading default config files....")
	AskRetry(self.CloneGitRepository)

	terminal.Output("Copying default config files....")
	AskRetry(self.CopyGeneralConfigFiles)

	//// Enable IOMMU in grub
	terminal.Output("Copying processor specific config files and enabling IOMMU in grub....")
	AskRetry(self.DoProcessorSpecificConfig)

	terminal.Output("Adding modules to initramfs....")
	AskRetry(self.DoInitramfsConfig)

	return err
}


func (self *Installer) DoProcessorSpecificConfig() error {
	var sInfo sysinfo.SysInfo
	sInfo.GetSysInfo()
	if sInfo.CPU.Vendor == "AuthenticAMD" {
		Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-amd"), "/etc/default/grub")
		Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-amd"), "/etc/modules")
	} else if sInfo.CPU.Vendor == "GenuineIntel" {
		Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-intel"), "/etc/default/grub")
		Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-intel"), "/etc/modules")
	}
	terminal.Run("update-grub")
	return nil
}

func (self *Installer) DoInitramfsConfig() error {
	Copy(self.Paths.BaseFile(machine.Host, "/etc/initramfs-tools/modules"), "/etc/initramfs-tools/modules")
	terminal.Run("update-initramfs -u")
	return nil
}

