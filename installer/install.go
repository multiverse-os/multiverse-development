package install

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/AlecAivazis/survey"
	"github.com/multiverse-os/color"
	"github.com/zcalusic/sysinfo"
)

type InstallStep int

const (
	PrepareSystem InstallStep = iota
	DownloadOSInstallMedia
	InstallConfigFiles
	SetupNetworking
	BuildRouterVMs
)

type Installer struct {
	Step InstallStep
	User User
	Paths Paths
}


type Paths struct {
	Home string
	Git string
	Etc string
	Var string
}


func (self Installer) BaseFiles(m MachineType) string {
	return fmt.Sprintf("%s/%s/base-files", self.Paths.Git, m.String())
}

type User struct {
	uid int
	gid int
}

var usr User

// # Multiverse OS Script Color Palette
// #==============================================================================
// header="\e[0;95m" Fuchsia
// accent="\e[37m" Silver
// subheader="\e[98m" ??
// strong="\e[96m" Aqua
// text="\e[94m" skyBlue
// success="\e[92m" Lime
// warning="\e[93m" Yellow
// fail="\e[91m" Red
// reset="\e[0m"
// #==============================================================================
func Header(text string) string  { return color.Fuchsia(text) }
func Accent(text string) string  { return color.Silver(text) }
func Strong(text string) string  { return color.Aqua(text) }
func Text(text string) string    { return color.SkyBlue(text) }
func Success(text string) string { return color.Lime(text) }
func Warning(text string) string { return color.Yellow(text) }
func Fail(text string) string    { return color.Red(text) }

// Actually maybe replace log with something that wraps lines with above color
// funcs

func Install(m MachineType) {

	switch m {
	case ControllerVM:
	default: // HostMachine
		installer := Installer{
			Step: PrepareSystem,
			Paths: Paths{
				Home: "/home/user",
				Git: "/var/multiverse/development",
				Etc: "/etc/multiverse",
				Var: "/var/multiverse",
			},
		}
	}

	// Check if superuser
	provisioner.user, err = user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if provisioner.user.Uid != "0" {
		log.Fatal(Fail("Must be superuser"))
	}

	//// User
	fmt.Println(Text("Setting up user account...."))
	AskRetry(SetupUser)

	//// Packages
	pm := NewPackageManager(Apt)
	fmt.Println(Text("Updating package lists......"))
	AskRetry(pm.Update)
	fmt.Println(Text("Upgrading packages......"))
	AskRetry(pm.Upgrade)
	fmt.Println(Text("Installing packages......"))
	// TODO best way to call this with AskRetry. Global with list of packages?
	if err := pm.InstallPackages([]string{
		"ovmf",
		"qemu",
		"qemu-system-common",
		"virt-manager",
		"pass",
		"git",
		"dirmngr",
		"vim",
	}); err != nil {
		log.Fatalf("can't install packages: %v\n", err)
	}
	fmt.Println(Text("Removing unnecessary packages......"))
	// TODO best way to call this with AskRetry. Global with list of packages?
	if err = pm.RemovePackages([]string{"nano", "minissdpd"}); err != nil {
		log.Fatal(Fail(fmt.Sprintf("can't remove packages: %v\n", err)))
	}
	AskRetry(pm.Autoremove)

	//// Default Paths
	// TODO handle CreateDir, etc errors

	fmt.Println(Text("Creating default filepath...."))
	AskRetry(CreateMultiversePaths)

	//// Configurations
	//// Install Config files
	fmt.Println(Text("Downloading default config files...."))
	AskRetry(DownloadConfigFiles)

	fmt.Println(Text("Copying default config files...."))
	AskRetry(CopyGeneralConfigFiles)

	//// Enable IOMMU in grub
	fmt.Println(Text("Copying processor specific config files and enabling IOMMU in grub...."))

	AskRetry(DoProcessorSpecificConfig)
	fmt.Println(Text("Adding modules to initramfs...."))
	AskRetry(DoInitramfsConfig)

}

type step func() error

func AskRetry(s step) error {
	if err := s(); err != nil {
		var q = &survey.Select{
			Message: fmt.Sprintf("Step failed due to error: %v\nRetry?", err),
			Options: []string{"retry", "skip", "exit"},
			Default: "retry",
		}
		var resp string
		survey.AskOne(q, &resp)
		if resp == "retry" {
			return AskRetry(s)
		} else if resp == "skip" {
			return nil
		} else if resp == "exit" {
			log.Println(Fail(err.Error()))
			os.Exit(1)
			return err
		}
	}
	return nil
}

func (self *Provisioner) CreateMultiversePaths() error {
	if err := CreateDir(self.Var, 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portal-gun", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portal-gun/os-image", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/share", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/sockets", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/sockets/serial", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/sockets/channel", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/sockets/console", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := CreateDir(self.Var+"/portals/sockets/parallel", 0700, usr.uid, usr.gid); err != nil {
	}
	// Because libvirt recreates its default image folder if it's not detected,
	// let's link it to our primary default
	if err := os.Remove(self.Home + "/.local/share/libvirt/images"); err != nil {
	}
	if err := CreateDir(self.Home+"/.local/share/libvirt", 0755, usr.uid, usr.gid); err != nil {
	}
	if err := os.Symlink(self.Var+"/portals/disks/", self.Home+"/.local/share/libvirt/images"); err != nil {
	}

	return nil
}

func (self *Provisioner) SetupUser() error {
	usr, err := user.Lookup("user")
	if err != nil {
		log.Println(Fail("User 'user' required"))
		return err
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return err
	}
	usr.uid = uid

	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		return err
	}
	usr.gid = gid

	// TODO: What the ?
	if usr.HomeDir != self.Paths.Home {
		log.Printf("User home directory mismatch, setting it to '%v'\n", self.Paths.Home)
		usr.HomeDir = self.Path.Home
	}
	if err := os.Remove(self.Paths.Home + "/Desktop"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}
	if err := os.Remove(self.Paths.Home + "/Downloads"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}
	if err := os.Remove(self.Paths.Home + "/Documents"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}
	if err := os.Remove(self.Paths.Home + "/Music"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}
	if err := os.Remove(self.Paths.Home + "/Videos"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}
	if err := os.Remove(self.Paths.Home + "/Pictures"); err != nil {
		log.Printf(Warning(fmt.Sprintf("Cannot remove directory, %v\n", err)))
	}

	////// VM Setup (Usermode)
	//// NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

	fmt.Println(Text("Adding user to kvm and libvirt groups..."))

	if err := Terminal("usermod -a -G kvm user"); err != nil {
		return err
	}
	if err := Terminal("groupadd --system libvirt"); err != nil {
		log.Println(Warning(err.Error()))
	}
	if err := Terminal("usermod -a -G libvirt user"); err != nil {
		return err
	}

	return nil
}

func (self Provisioner) DownloadConfigFiles() error {
	if err := Terminal("git clone https://github.com/multiverse-os/multiverse-development " + GIT_SRC_PATH); err != nil {
	// TODO
	// better than erroring if the directory is already there is checking the git
	// error and cd + git pull instead
	return err
}
// TODO wtf is this rm sh clone sh?
//cd USER_HOME/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh
// TODO is os.Chown recursive or do I have to filewalk it?
if err := filepath.Walk(self.Paths.Git, func(name string, info os.FileInfo, err error) error {
	if err := os.Chown(name, usr.uid, usr.gid); err != nil {
		return err
	}
	return nil
}); err != nil {
	return err
}
return nil
}

func (self Provisioner) CopyGeneralConfigFiles() error {
	// NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
	// TODO contemplate the implications of making these config files user
	// editable
	baseFilesPath := Provisioner.Paths.Git + "/host/base-files"
	if err := Copy(baseFilesPath+"/home/user/.gitconfig", self.Paths.Home+"/.gitconfig"); err != nil {
		return err
	}
	if err := os.Chown(self.Paths.Home+"/.gitconfig", usr.uid, usr.gid); err != nil {
		return err
	}

	if err := CreateDir("/etc/multiverse", 0700, usr.uid, usr.gid); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/motd", "/etc/motd"); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/issue", "/etc/issue"); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/security/limits.conf", "/etc/security/limits.conf"); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/sysctl.conf", "/etc/sysctl.conf"); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/sysctl.d/30-tracker.conf", "/etc/sysctl.d/30-tracker.conf"); err != nil {
		return err
	}
	if err := Copy(baseFilesPath+"/etc/sysctl.d/99-sysctl.conf", "/etc/sysctl.d/99-sysctl.conf"); err != nil { // TODO everything is commented out , is this file necessary?
		return err
	}
	// TODO some of the rc.local stuff is vfio passthrough that should be done in
	// another step
	if err := Copy(baseFilesPath+"/etc/rc.local", "/etc/rc.local"); err != nil {
		return err
	}
	// TODO is bridge.conf obsolete yet?
	//Copy(baseFilesPath + "/etc/qemu/bridge.conf", "/etc/qemu/bridge.conf")
	return nil
}

func (self *Provisioner) DoProcessorSpecificConfig() error {
	baseFilesPath := self.Git + "/host/base-files"
	var sInfo sysinfo.SysInfo
	sInfo.GetSysInfo()
	if sInfo.CPU.Vendor == "AuthenticAMD" {
		if err := Copy(baseFilesPath+"/etc/default/grub-amd", "/etc/default/grub"); err != nil {
			return err
		}
		if err := Copy(baseFilesPath+"/etc/modules-amd", "/etc/modules"); err != nil {
			return err
		}
	} else if sInfo.CPU.Vendor == "GenuineIntel" {
		if err := Copy(baseFilesPath+"/etc/default/grub-intel", "/etc/default/grub"); err != nil {
			return err
		}
		if err := Copy(baseFilesPath+"/etc/modules-intel", "/etc/modules"); err != nil {
			return err
		}
	}
	if err := Terminal("update-grub"); err != nil {
		return err
	}
	return nil
}

func (self *Provisioner) DoInitramfsConfig() error {
	baseFilesPath := self.Git + "/host/base-files"
	if err := Copy(baseFilesPath+"/etc/initramfs-tools/modules", "/etc/initramfs-tools/modules"); err != nil {
		return err
	}
	if err := Terminal("update-initramfs -u"); err != nil {
		return err
	}
	return nil
}

////// SH Framework
//// Copy over vfio-bind into binary execution path

////// Network Bridges
//// NOTE: To be replaced with sockets
//chown -R root:kvm /usr/lib/qemu/
//chmod 4750 /usr/lib/qemu/qemu-bridge-helper

//$GIT_SRC_PATH/host/scripts/add-bridge.sh $GIT_SRC_PATH/host/xml/networks/net0br0.xml
////net0br1.xml  net0br2.xml  net1br0.xml  net1br1.xml  net1br2.xml

// echo -e $strong"Downloading Linux distributions$accent os-images$reset needed for Multiverse OS installation..."$reset

// cd $GIT_SRC_PATH/images/os-images && ./alpine-dl-and-verify.sh
