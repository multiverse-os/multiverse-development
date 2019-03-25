##
##  Multiverse OS: Image Builder (Installer, Live CD, and VMs)
=========================================================================
### Visualize Crawl
[Network]
**Phylotree like graph**
https://github.com/datastorm-open/visNetwork
http://datastorm-open.github.io/visNetwork/
**visualize image tree versioning**
====================================
# Multiverse OS Unikernels
http://unikernel.org/projects/

<Resources: http://unikernel.org/resources/>

**OSv is the most complete**
==============================
# Multiverse OS Init Systems

https://github.com/mesanine/ginit

https://github.com/QuentinPerez/busygox
https://github.com/artyom/boxinit
https://github.com/u-root/u-root
https://github.com/chazomaticus/minit


https://github.com/cloudius-systems/virtramfs/blob/master/virtramfs.c
==========================
https://github.com/aelsabbahy/goss
##
## uinput
https://github.com/ynsta/uinput
https://github.com/eikenb/udev-notify
## simple cotnainer
https://github.com/converseai/simple_container
https://github.com/kasheemlew/xperiMoby
!!!!!!!!!!!!!!!!!!!
BEST SYS INFO https://github.com/zcalusic/sysinfo

!!!!!!!!!!!!!!
https://github.com/box-builder/box
USES FUCKING MRUBYT O CONFIG CONTAINERS (DOCKER BUT WHO GIVES A FUCK!)



[css only fileicons](https://github.com/picturepan2/fileicon.css)


https://github.com/box-builder/overmount

https://github.com/linuxkit/linuxkit

https://github.com/hackwave/goiardi - CHEF IN GO

https://github.com/AdaptiveScale/lxdui
https://github.com/hackwave/lxdimage even more simple image builder with lxd

https://github.com/hackwave/golang-kmod - modify modules loaded

https://github.com/hackwave/railcar/ RUST CONTAINER
https://github.com/cyphar/initrs INIT AND SIGNAL IN RUST

https://github.com/hackwave/goiardi CHEF IN GO

https://github.com/taskcluster/taskcluster-worker/blob/f980ce1e71a94a1b159b0c0e0a3a2dff2f3e4f72/engines/qemu/vm/vm.go
https://github.com/Nilesh20/Dependencies/tree/766ecc1586690cdc54c43e54ef8f0a93f0003bc8/gitrepo/github.com/vmware/govmomi/govc/vm
https://github.com/yangaowei/VMonitor/blob/9f5429d0262848dfd91ffe40fc9459afd777e282/kvmtop.go

https://github.com/linuxdeepin/dde-daemon/blob/50c198f69cf8bbcf1a928d676acb0bf4c952b724/network/manager_device.go


https://github.com/laincloud/networkd/blob/6ef4c8ee01b186fdbd1cbd85f00adc5b7edbbc4a/virtualip.go

https://github.com/GuomeiQtl/cgroupfs/blob/a5033c07d29c74210348985d1b89b711adb70676/fs/hello.go

https://github.com/mdlayher/talks/blob/6a208cf817e4880c1dcd9c5e35460791a2af70dc/ethernet-and-go/3/main.go


pening a raw socket
To open a socket, you have to know three things  the socket family, socket type and protocol. For a raw socket, the socket family is AF_PACKET, the socket type is SOCK_RAW and for the protocol, see the if_ether.h header file. To receive all packets, the macro is ETH_P_ALL and to receive IP packets, the macro is ETH_P_IP for the protocol field.
int sock_r;
sock_r=socket(AF_PACKET,SOCK_RAW,htons(ETH_P_ALL));
if(sock_r<0)
{
printf(error in socket\n);
return -1;
}


Tools:
genisoimage
debootstrap

https://github.com/cloudfoundry/syslog-release

https://github.com/godbus/dbus
https://github.com/TheCreeper/go-notify
https://github.com/AmandaCameron/go.networkmanager
https://github.com/zalando/go-keyring

https://github.com/j-keck/arping
https://github.com/fabric8io/kansible

https://github.com/manifoldco/promptui UI
https://github.com/NVZGUL/NetInterfaceApi

https://github.com/samalba/buse-go
https://github.com/teemow/loopback
https://github.com/aelsabbahy/goss

https://github.com/vechain/solidb

https://github.com/mars9/ramfs - 9p but all in memory!


https://github.com/multiformats/go-multiaddr

https://github.com/vasi/qcow2
https://github.com/briandowns/aion - rest server cron
## MV Spec
https://github.com/containers/virtcontainers/blob/master/hook.go - multiverse should support hooks like OCI

### OHT Networking
Using a combination of either ((1) Virtual Device, (2) Tunnel Device, (3) Unix Socket, (3) Device (FIFO, Plain FD)) & a built in transparent proxy. Route all traffic to tor going through the device. 

Then build a piece a CLI to enable/disable routing all traffic through this device, or routing specific ports, and detect all onion services to add to available endpoints.
https://github.com/containernetworking/plugins/tree/master/plugins/main/macvlan
### Multiverse Networking
Using a combination of either (1) Virtual Device (2) Unix Socket (3) Device (FIFO, Plain FD). Likely we will do this on the HOST machine and PCI pass this over to the VM. 

The next option will be determining how best to connect these devices, because the goal is to avoid using ANY HOST KERNELSPACE!

So the first idea would be:
Then using *Userland* pipe software we connect the two devices together. 

But this is not ideal, and gives too much information to the host, we are trying to prevent the host from being able to get anything valuable.

*the best way will liekly be to connect mount devices built in userland on the HOST to the computer it will connect to AND to the router. The devices should ecnrypt traffic to prevent leaking. then the userland OR kernel space pipe can then function as a router to route all the data to these devices. One benefit is we could use multi or tee to dump packets and do analysis. 

Additionally we could write the software as Softare defined routing software and support OHT based jumping / wormholes between any rotuer (and supporting device-to-device encryption)

[Resourecs]
https://github.com/containernetworking/plugins/tree/master/plugins/main/macvlan


// IPv4ForwardingEnabled returns true if the kernel is configured to forward IPv4 packets.
func IPv4ForwardingEnabled() (bool, error) {
	d, err := ioutil.ReadFile("/proc/sys/net/ipv4/ip_forward")
	if err != nil {
		return false, err
	}
	if len(d) != 2 {
		return false, fmt.Errorf("expected single byte read, got %d", len(d))
	}
	return d[0] == '1', nil
}

// IPv4EnableForwarding enables or disables forwarding of IPv4 packets.
func IPv4EnableForwarding(state bool) error {
	outData := "0"
	if state {
		outData = "1"
	}
	return ioutil.WriteFile("/proc/sys/net/ipv4/ip_forward", []byte(outData), 0644)
}


https://github.com/twitchyliquid64/bob-the-builder
############ https://github.com/lxc/distrobuilder <IMPORTANT>
https://github.com/linuxkit/linuxkit

jail containers https://github.com/singularityware/singularity


<transparent proxy>https://github.com/elazarl/goproxy

https://github.com/coreos/fleet distributed init system

qemu-utils
systemd-container (systemd-nspawn)
btrfs-progs / btrfs-tools

namespaces
https://github.com/coreos/go-namespaces
tcp proxy for ns https://github.com/coreos/nsproxy

<BIOS>https://github.com/coreos/seabios

https://github.com/coreos/bcrypt-tool
<IMPORTANT>https://github.com/coreos/go-iptables

https://github.com/coreos/coreos-cloudinit
<IMPORTANT> https://raw.githubusercontent.com/coreos/baselayout/master/modprobe.d/aliases.conf
## Examples
	# Create ISOLINUX configuration file.
	echo 'default kernel.bz  initrd=rootfs.gz root=/dev/ram0' > ./isolinux.cfg

	# Now we generate the ISO image file.
	genisoimage -J -r -o ../debtrap_linux_live.iso -b isolinux.bin -c boot.cat -no-emul-boot -boot-load-size 4 -boot-info-table ./

	# This allows the ISO image to be bootable if it is burned on USB flash drive.
	isohybrid ../deptrap_linux_live.iso 2>/dev/null || true



## General
https://github.com/traviscross/debootstrap-img

https://help.ubuntu.com/community/LiveCDCustomizationFromScratch


[uspin](https://github.com/solus-project/USpin)
A LiveOS image is an ISO9660 image containing a live operating system. This is the dracut LiveOS image type, currently used by Solus, Fedora, available in Gentoo and potentially others.
		Add parser for the Solus image specification format
		Port the Stack implementation from old image creator
		Add config format for the main image configuration
		Add utilities for image format & creation (dd/fallocate, etc)
		Implement full eopkg support in generic pkg.Manager interface
		Add basic ISO9660 support once again
		Add complete Legacy Boot bootloader support for isolinux

**Scripts**
https://github.com/josug-book1-materials/ubuntu-virtinst

https://github.com/grml/grml-debootstrap/blob/master/packer/debian64_provision.sh

**deboostrap+shell script**
https://github.com/ericdwhite/mirrored-debootstrap-install
This repository contains scripts to automatically create a KVM image based on the parameters specified in config.sh.


https://github.com/larshvile/image-debootstrap/tree/master/install
even has modprobe stuff


https://github.com/rmariotti/debian_auto_installer

**preeseed**
https://github.com/hitsuji/steamos/blob/master/preseed

**Ansible**
https://github.com/nilsmeyer/ansible-debootstrap/blob/master/defaults/main.yml


https://github.com/kkndyu/wanna-ansible

## Installer

## Live
