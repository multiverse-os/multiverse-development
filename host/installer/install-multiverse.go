package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/multiverse-os/color"
	"github.com/zcalusic/sysinfo"
)

const (
	USER_HOME      = "/home/user"
	GIT_SRC_PATH   = USER_HOME + "/multiverse"
	MV_CONFIG_PATH = "/etc/multiverse"
	MV_SYSTEM_PATH = "/var/multiverse"
)

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

func main() {
	// Check if superuser
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal(Fail("Must be superuser"))
	}

	//// User
	fmt.Println(Text("Setting up user account...."))
	// Get user "user"
	uzer, err := SetupUser()
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	uid, err := strconv.Atoi(uzer.Uid)
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	gid, err := strconv.Atoi(uzer.Gid)
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	//// Packages
	pm := NewPackageManager(Apt)
	fmt.Println(Text("Updating package lists......"))
	if err := pm.Update(); err != nil {
		log.Fatal(Fail(fmt.Sprintf("can't update package lists: %v\n", err)))
	}
	fmt.Println(Text("Upgrading packages......"))
	if err := pm.Upgrade(); err != nil {
		log.Fatal(Fail(fmt.Sprintf("can't upgrade packages: %v\n", err)))
	}
	fmt.Println(Text("Installing packages......"))
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
	if err := pm.Autoremove(); err != nil {
		log.Fatal(Fail(fmt.Sprintf("can't remove packages: %v\n", err)))
	}
	if err = pm.RemovePackages([]string{"nano", "minissdpd"}); err != nil {
		log.Fatal(Fail(fmt.Sprintf("can't remove packages: %v\n", err)))
	}

	//// Default Paths
	// TODO handle CreateDir, etc errors

	fmt.Println(Text("Creating default filepath...."))
	if err = CreateMultiversePaths(uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	//// Configurations
	//// Install Config files
	fmt.Println(Text("Downloading default config files...."))
	if err := DownloadConfigFiles(uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	fmt.Println(Text("Copying config files..."))
	if err := CopyUserConfigFiles(uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	if err := CopyGeneralConfigFiles(uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	//// Enable IOMMU in grub
	fmt.Println(Text("Copying processor specific config files and enabling IOMMU in grub...."))

	if err := DoProcessorSpecificConfig(); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	fmt.Println(Text("Adding modules to initramfs...."))
	if err := DoInitramfsConfig(); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

}

func CreateMultiversePaths(uid, gid int) error {
	if err := CreateDir(MV_SYSTEM_PATH, 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portal-gun", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portal-gun/os-image", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/share", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/sockets", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/sockets/serial", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/sockets/channel", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/sockets/console", 0700, uid, gid); err != nil {
		return err
	}
	if err := CreateDir(MV_SYSTEM_PATH+"/portals/sockets/parallel", 0700, uid, gid); err != nil {
	}
	// Because libvirt recreates its default image folder if it's not detected,
	// let's link it to our primary default
	if err := os.Remove(USER_HOME + "/.local/share/libvirt/images"); err != nil {
	}
	if err := CreateDir(USER_HOME+"/.local/share/libvirt", 0755, uid, gid); err != nil {
	}
	if err := os.Symlink(MV_SYSTEM_PATH+"/portals/disks/", USER_HOME+"/.local/share/libvirt/images"); err != nil {
	}

	return nil
}

func SetupUser() (*user.User, error) {
	uzer, err := user.Lookup("user")
	if err != nil {
		return uzer, err
	}

	if uzer.HomeDir != USER_HOME {
		log.Printf("User home directory mismatch, setting it to '%v'\n", USER_HOME)
		uzer.HomeDir = USER_HOME
	}
	if err := os.Remove(USER_HOME + "/Desktop"); err != nil {
		return uzer, err
	}
	if err := os.Remove(USER_HOME + "/Downloads"); err != nil {
		return uzer, err
	}
	if err := os.Remove(USER_HOME + "/Documents"); err != nil {
		return uzer, err
	}
	if err := os.Remove(USER_HOME + "/Music"); err != nil {
		return uzer, err
	}
	if err := os.Remove(USER_HOME + "/Videos"); err != nil {
		return uzer, err
	}
	if err := os.Remove(USER_HOME + "/Pictures"); err != nil {
		return uzer, err
	}

	////// VM Setup (Usermode)
	//// NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

	fmt.Println(Text("Adding user to kvm and libvirt groups..."))

	if err := Terminal("usermod -a -G kvm user"); err != nil {
		return uzer, err
	}
	if err := Terminal("groupadd --system libvirt"); err != nil {
		return uzer, err
	}
	if err := Terminal("usermod -a -G libvirt user"); err != nil {
		return uzer, err
	}

	return uzer, nil
}

func DownloadConfigFiles(uid, gid int) error {
	if err := Terminal("git clone https://github.com/multiverse-os/multiverse-development " + GIT_SRC_PATH); err != nil {
		return err
	}
	// TODO wtf is this rm sh clone sh?
	//cd USER_HOME/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh
	// TODO is os.Chown recursive or do I have to filewalk it?
	if err := filepath.Walk(GIT_SRC_PATH, func(name string, info os.FileInfo, err error) error {
		if err := os.Chown(name, uid, gid); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func CopyUserConfigFiles(uid, gid int) error {
	if err := Copy("./home/user/.gitconfig", USER_HOME+"/.gitconfig"); err != nil {
		return err
	}
	if err := os.Chown(USER_HOME+"/.gitconfig", uid, gid); err != nil {
		return err
	}

	return nil
}

func CopyGeneralConfigFiles(uid, gid int) error {
	// NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
	// TODO contemplate the implications of making these config files user
	// editable
	if err := CreateDir("/etc/multiverse", 0700, uid, gid); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/motd", "/etc/motd"); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/issue", "/etc/issue"); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/security/limits.conf", "/etc/security/limits.conf"); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/sysctl.conf", "/etc/sysctl.conf"); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/sysctl.d/30-tracker.conf", "/etc/sysctl.d/30-tracker.conf"); err != nil {
		return err
	}
	if err := Copy(MV_CONFIG_PATH+"/etc/sysctl.d/99-sysctl.conf", "/etc/sysctl.d/99-sysctl.conf"); err != nil { // TODO everything is commented out , is this file necessary?
		return err
	}
	// TODO some of the rc.local stuff is vfio passthrough that should be done in
	// another step
	if err := Copy(MV_CONFIG_PATH+"/etc/rc.local", "/etc/rc.local"); err != nil {
		return err
	}
	// TODO is bridge.conf obsolete yet?
	//Copy(MV_CONFIG_PATH + "/etc/qemu/bridge.conf", "/etc/qemu/bridge.conf")
	return nil
}

func DoProcessorSpecificConfig() error {
	var sInfo sysinfo.SysInfo
	sInfo.GetSysInfo()
	if sInfo.CPU.Vendor == "AuthenticAMD" {
		if err := Copy(MV_CONFIG_PATH+"/etc/default/grub-amd", "/etc/default/grub"); err != nil {
			return err
		}
		if err := Copy(MV_CONFIG_PATH+"/etc/modules-amd", "/etc/modules"); err != nil {
			return err
		}
	} else if sInfo.CPU.Vendor == "GenuineIntel" {
		if err := Copy(MV_CONFIG_PATH+"/etc/default/grub-intel", "/etc/default/grub"); err != nil {
			return err
		}
		if err := Copy(MV_CONFIG_PATH+"/etc/modules-intel", "/etc/modules"); err != nil {
			return err
		}
	}
	if err := Terminal("update-grub"); err != nil {
		return err
	}
	return nil
}

func DoInitramfsConfig() error {
	if err := Copy(MV_CONFIG_PATH+"/etc/initramfs-tools/modules", "/etc/initramfs-tools/modules"); err != nil {
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
