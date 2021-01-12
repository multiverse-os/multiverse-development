package install

import (
	"fmt"
	"os/user"

	machine "./machine"

	"github.com/zcalusic/sysinfo"
)


///////////////////////////////////////////////////////////////////////////////
type Packages struct {
	Install []string
	Remove []string
}

// TODO: The idea is overtime this can be used to work to function as an 
//       installer for various 
type Installer struct {
	Step InstallStep
	User *User
	Paths Paths
	Packages Packages
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
func (self *Installer) Install() {
	if !IsRoot() { panic(fmt.Errof("must be root")) }


	//// User
	AskRetry(SetupUser)

	//// Packages
	pm := NewPackageManager(Apt)
	Output("Updating package lists...")
	AskRetry(pm.Update)
	Output("Upgrading packages...")
	AskRetry(pm.Upgrade)
	Output("Installing packages...")

	if err = pm.RemovePackages(self.Packages.Remove); err != nil {
		panic(fmt.Errorf("can't remove packages: %v\n", err))
	}
	if err := pm.InstallPackages(self.Packages.Install); err != nil {
		panic(fmt.Errof("can't install packages: %v\n", err))
	}

	Output("Removing unnecessary packages......")
	// TODO best way to call this with AskRetry. Global with list of packages?
	AskRetry(pm.Autoremove)

	Output("Creating default filepath....")
	AskRetry(CreateMultiversePaths)

	//// Configurations
	//// Install Config files
	Output("Downloading default config files....")
	AskRetry(DownloadConfigFiles)

	Output("Copying default config files....")
	AskRetry(CopyGeneralConfigFiles)

	//// Enable IOMMU in grub
	Output("Copying processor specific config files and enabling IOMMU in grub....")
	AskRetry(DoProcessorSpecificConfig)

	Output("Adding modules to initramfs....")
	AskRetry(DoInitramfsConfig)
}


func (self *Installer) DoProcessorSpecificConfig() error {
	var sInfo sysinfo.SysInfo
	sInfo.GetSysInfo()
	if sInfo.CPU.Vendor == "AuthenticAMD" {
		err = Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-amd"), "/etc/default/grub")
		err = Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-amd"), "/etc/modules")
	} else if sInfo.CPU.Vendor == "GenuineIntel" {
		err = Copy(self.Paths.BaseFile(machine.Host, "/etc/default/grub-intel"), "/etc/default/grub")
		err = Copy(self.Paths.BaseFile(machine.Host, "/etc/modules-intel"), "/etc/modules")
	}
	err = Terminal("update-grub")
	return nil
}

func (self *Installer) DoInitramfsConfig() error {
	err = Copy(self.BaseFile(machine.Host, "/etc/initramfs-tools/modules"), "/etc/initramfs-tools/modules")
	err = Terminal("update-initramfs -u")
	return nil
}

