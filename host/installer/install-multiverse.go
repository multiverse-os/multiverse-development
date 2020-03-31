package main

import (
	"log"
	"os"
	"os/user"
)

const (
	USER_HOME      = "/home/user"
	GIT_SRC_PATH   = USER_HOME + "/multiverse/"
	MV_CONFIG_PATH = "/etc/multiverse/"
)

func main() {
	// Check if superuser
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("Must be superuser")
	}

	// Get user "user"
	uzer, err := user.Lookup("user")
	if err != nil {
		log.Fatal("requires a user named \"user\": %v\n", err)
	}

	if uzer.HomeDir != USER_HOME {
		log.Printf("User home directory mismatch, setting it to '%v'\n", USER_HOME)
		uzer.HomeDir = USER_HOME
	}

	//// Packages
	pm := NewPackageManager(Apt)
	pm.Update()
	pm.Upgrade()
	if err := pm.InstallPackages(
		"ovmf",
		"qemu",
		"qemu-system-common",
		"virt-manager",
		"pass",
		"git",
		"dirmngr",
		"vim",
	); err != nil {
		log.Fatalf("can't install packages: %v\n", err)
	}
	if err = pm.RemovePackages("nano", "minissdpd"); err != nil {
		log.Fatalf("can't remove packages: %v\n", err)
	}

	//// Default Paths
	// TODO handle os.Mkdir, etc errors
	os.Mkdir("/var/multiverse/portal-gun/", 0700)
	os.Mkdir("/var/multiverse/portal-gun/os-image", 0700)
	os.Mkdir("/var/multiverse/portals/share", 0700)
	os.Mkdir("/var/multiverse/portals/disk", 0700)
	os.Mkdir("/var/multiverse/portals/sockets/serial", 0700)
	os.Mkdir("/var/multiverse/portals/sockets/channel", 0700)
	os.Mkdir("/var/multiverse/portals/sockets/console", 0700)
	os.Mkdir("/var/multiverse/portals/sockets/parallel", 0700)
	// TODO is os.Chown recursive or do I have to filewalk it?
	os.Chown("var/multiverse/", uzer.Uid, uzer.Gid)

	os.Mkdir("/etc/multiverse", 0700)
	// Because libvirt recreates its default image folder if it's not detected,
	// let's link it to our primary default
	os.Remove(uzer.Homedir + "/.local/share/libvirt/images")
	os.Symlink("/var/multiverse/portals/disks/", uzer.Homedir+"/.local/share/libvirt/images")

	//// User
	os.Remove(uzer.Homedir + "Desktop")
	os.Remove(uzer.Homedir + "Downloads")
	os.Remove(uzer.Homedir + "Documents")
	os.Remove(uzer.Homedir + "Music")
	os.Remove(uzer.Homedir + "Videos")
	os.Remove(uzer.Homedir + "Pictures")

	////// VM Setup (Usermode)
	//// NOTE: Would be better to move this to root:kvm and avoid needing libvirt group altogether

	Terminal("usermod -a -G kvm user")
	Terminal("usermod -a -G libvirt user")

	os.Chdir(uzer.Homedir)
	Terminal("git clone https://github.com/multiverse-os/multiverse-development multiverse")
	// TODO wtf is this rm sh clone sh?
	//cd uzer.Homedir/multiverse/ && rm -rf sh && git clone https://github.com/multiverse-os/sh
	// TODO is os.Chown recursive or do I have to filewalk it?
	os.Chown("multiverse", uzer.Uid, uzer.Gid)

	//// Configurations
	//// Install Config files
	os.Chdir(uzer.Homedir + "/multiverse/host/base-files")
	Copy("./home/user/.gitconfig", uzer.HomeDir+"/.gitconfig")
	os.Chown(uzer.HomeDir+".gitconfig", uzer.Uid, uzer.Gid)

	// NOTE: Track all changes needed for setting up Multiverse, this will simplify the process and all these can be kept in /etc/multiverse and symbolically linked. Then the rest of the /et/multiverse folder can be custom Multiverse OS config files which will most likely be ruby or YAML based.
	Copy("./etc/motd", "/etc/motd")
	Copy("./etc/modules", "/etc/modules")
	Copy("./etc/issue", "/etc/issue")
	Copy("./etc/security/limits.conf", "/etc/security/limits.conf")
	Copy("./etc/sysctl.conf", "/etc/sysctl.conf")
	Copy("./etc/sysctl.d/30-tracker.conf", "/etc/sysctl.d/30-tracker.conf")
	Copy("./etc/sysctl.d/99-sysctl.conf", "/etc/sysctl.d/99-sysctl.conf") // TODO everything is commented out , is this file necessary?
	// TODO some of the rc.local stuff is vfio passthrough that should be done in
	// another step
	Copy("./etc/rc.local", "/etc/rc.local")
	// TODO is bridge.conf obsolete yet?
	//Copy("./etc/qemu/bridge.conf", "/etc/qemu/bridge.conf")

	//// Enable IOMMU in grub
	// TODO
	// Use https://github.com/zcalusic/sysinfo to get CPU info and install correct
	// grub config
	//// If using Intel procecssor, comment out grub-amd and uncomment grub-intel
	Copy("./base-files/etc/default/grub-amd", "/etc/default/grub")
	//cp $GIT_SRC_PATH/base-files/etc/default/grub-intel /etc/default
	Terminal("update-grub")
}

////// SH Framework
//// Copy over vfio-bind into binary execution path

////// Network Bridges
//// NOTE: To be replaced with sockets
//chown -R root:kvm /usr/lib/qemu/
//chmod 4750 /usr/lib/qemu/qemu-bridge-helper

//$GIT_SRC_PATH/host/scripts/add-bridge.sh $GIT_SRC_PATH/host/xml/networks/net0br0.xml
////net0br1.xml  net0br2.xml  net1br0.xml  net1br1.xml  net1br2.xml
