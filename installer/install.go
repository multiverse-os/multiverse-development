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

type System struct {
	CPU cpu.Architecture
}

// TODO: The idea is overtime this can be used to work to function as an 
//       installer for various 
type Installer struct {
	Step InstallStep
	User *User
	Paths Paths
	Packages Packages
	System System
}


///////////////////////////////////////////////////////////////////////////////
func To(m machine.Type) (installer Installer) {
	switch m {
	case machine.Controller:
	case machine.Host:
		installer = Installer{
			Packages: Packages{
				Install: []string{ "ovmf", "virt-manager", "pass", "git", "golang", "ruby", "dirmngr", "vim", "neovim"},

				Remove: []string{"nano", "minissdpd"},
			},
			Step: PrepareSystem,
			Paths: Paths{
				HomePath: "/home/user",
				GitPath: "/var/multiverse/development",
				EtcPath: "/etc/multiverse",
				VarPath: "/var/multiverse",
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
	pm := NewPackageManager(Apt)
	terminal.Output("Updating package lists...")
	AskRetry(pm.Update)
	terminal.Output("Upgrading packages...")
	AskRetry(pm.Upgrade)
	terminal.Output("Installing packages...")

	if err = pm.RemovePackages(self.Packages.Remove); err != nil {
		panic(fmt.Errorf("can't remove packages: %v\n", err))
	}
	if err := pm.InstallPackages(self.Packages.Install); err != nil {
		panic(fmt.Errorf("can't install packages: %v\n", err))
	}

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

