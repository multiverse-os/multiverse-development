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
	GIT_SRC_PATH   = USER_HOME + "/multiverse/"
	MV_CONFIG_PATH = "/etc/multiverse/"
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

	// Get user "user"
	uzer, err := user.Lookup("user")
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf("requires a user named \"user\": %v\n", err)))
	}

	if uzer.HomeDir != USER_HOME {
		log.Printf("User home directory mismatch, setting it to '%v'\n", USER_HOME)
		uzer.HomeDir = USER_HOME
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

	uid, err := strconv.Atoi(uzer.Uid)
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	gid, err := strconv.Atoi(uzer.Gid)
	if err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	if err = CreateDir("/var/multiverse", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portal-gun", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portal-gun/os-image", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/share", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/sockets", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/sockets/serial", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/sockets/channel", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/sockets/console", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir("/var/multiverse/portals/sockets/parallel", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	// TODO contemplate the implications of making these config files user
	// editable
	if err = CreateDir("/etc/multiverse", 0700, uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	// Because libvirt recreates its default image folder if it's not detected,
	// let's link it to our primary default
	if err = os.Remove(uzer.HomeDir + "/.local/share/libvirt/images"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = CreateDir(uzer.HomeDir+"/.local/share/libvirt", 0755, uid, gid); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Symlink("/var/multiverse/portals/disks/", uzer.HomeDir+"/.local/share/libvirt/images"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	//// User
	if err = os.Remove(uzer.HomeDir + "/Desktop"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Remove(uzer.HomeDir + "/Downloads"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Remove(uzer.HomeDir + "/Documents"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Remove(uzer.HomeDir + "/Music"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Remove(uzer.HomeDir + "/Videos"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Remove(uzer.HomeDir + "/Pictures"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}

	////// VM Setup (Usermode)
	//// NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

	fmt.Println(Text("Adding user to kvm and libvirt groups..."))

	if err = Terminal("usermod -a -G kvm user"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Terminal("groupadd --system libvirt"); err != nil {
		log.Println(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Terminal("usermod -a -G libvirt user"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	//// Configurations
	//// Install Config files
	fmt.Println(Text("Copying default config files...."))

	if err = os.Chdir(uzer.HomeDir); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Terminal("git clone https://github.com/multiverse-os/multiverse-development multiverse"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	// TODO wtf is this rm sh clone sh?
	//cd uzer.HomeDir/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh
	// TODO is os.Chown recursive or do I have to filewalk it?
	if err = filepath.Walk("multiverse", func(name string, info os.FileInfo, err error) error {
		if err = os.Chown(name, uid, gid); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	if err = os.Chdir(uzer.HomeDir + "/multiverse/host/base-files"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./home/user/.gitconfig", uzer.HomeDir+"/.gitconfig"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = os.Chown(uzer.HomeDir+"/.gitconfig", uid, gid); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	// NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
	if err = Copy("./etc/motd", "/etc/motd"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./etc/issue", "/etc/issue"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./etc/security/limits.conf", "/etc/security/limits.conf"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./etc/sysctl.conf", "/etc/sysctl.conf"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./etc/sysctl.d/30-tracker.conf", "/etc/sysctl.d/30-tracker.conf"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Copy("./etc/sysctl.d/99-sysctl.conf", "/etc/sysctl.d/99-sysctl.conf"); err != nil { // TODO everything is commented out , is this file necessary?
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	// TODO some of the rc.local stuff is vfio passthrough that should be done in
	// another step
	if err = Copy("./etc/rc.local", "/etc/rc.local"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	// TODO is bridge.conf obsolete yet?
	//Copy("./etc/qemu/bridge.conf", "/etc/qemu/bridge.conf")

	//// Enable IOMMU in grub
	fmt.Println(Text("Copying processor specific config files and enabling IOMMU in grub...."))
	if err = os.Chdir(uzer.HomeDir + "/multiverse/host/base-files"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	var sInfo sysinfo.SysInfo
	sInfo.GetSysInfo()
	if sInfo.CPU.Vendor == "AuthenticAMD" {
		if err = Copy("./etc/default/grub-amd", "/etc/default/grub"); err != nil {
			log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
		}
		if err = Copy("./etc/modules-amd", "/etc/modules"); err != nil {
			log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
		}
	} else if sInfo.CPU.Vendor == "GenuineIntel" {
		if err = Copy("./etc/default/grub-intel", "/etc/default/grub"); err != nil {
			log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
		}
		if err = Copy("./etc/modules-intel", "/etc/modules"); err != nil {
			log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
		}
	}
	if err = Terminal("update-grub"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

	fmt.Println(Text("Adding modules to initramfs...."))
	if err = Copy("./etc/initramfs-tools/modules", "/etc/initramfs-tools/modules"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}
	if err = Terminal("update-initramfs -u"); err != nil {
		log.Fatal(Fail(fmt.Sprintf(": %v\n", err)))
	}

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
