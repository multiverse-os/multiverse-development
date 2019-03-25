# install.notes.md

## os-images should not be in /var/multiverse

instead it should reside in either `/usr/lib/multiverse` or `~/.local/share/multiverse`, since it is not specific tothe machine it is on but rather shared between **Host Machine** and all **Controller VMs**. 

## Clarify whether commands should be run as root or unprivileged user
For example, running the `virsh pool-define ...` as root will not add the pools to the QEMU/KVM User session. I assumed those commands should be run as root because so far all of the other commands have required root to function.

!!! TEST !!!
can I `virsh net-define virbr0` as unpriv user since I have the bridge helper permissions and stuff set up?

## VT-d check script
It said yes, detected on my machine, even though it was disabled in BIOS

## Initramfs blacklist/enable modules
Add discussion of enabling/disabling in initramfs (pros/cons vs disabling in grub or /etc/modprobe.d/... blacklist)
Enabling in initramfs:

```
# /etc/initramfs-tools/modules
module_name
```

And update initramfs with `update-initramfs -u`


## Universe router Post Installation Configuration section
What files installed by provision.sh might need to be edited for the user's specific situation?
/etc/shorewall/hosts
/etc/shorewall/rules

Anything else? Preferred DHCP servers?

## Galaxy Router setup
Leaves out rc-update commands and setup instructions for dhcpd and shorewall. I'm assuming I don't want it the same as Universe?

## Star Router setup
### Whonix version
Note that this version of shrinking the qcow2 will still allow it to balloon to 100 GB (for example from runaway logs), you'd have to actually go into the VM and resize the filesystem and that kind of stuff to really make it stay small.

Take out info on using xml provided by Whonix team, doesn't work out of the box.

What are recommended values (memory, for example) when manually adding the qcow2? 256 sounds low, since it has a graphical ui. What is the min for it to run KDE? (make a note of the KDE auto-disable with low RAM allocation)
the XML file provided by Whonix says 524288 KiB

Assuming you assigned the machine 524 MB ram, it will boot into KDE.
Boot the VM, go through setup and enable Tor and the stable repo. 

Setup:
Agree to terms ("I understand" twice)
Select "I am ready to enable Tor"
Hit "Next" four times
Select "Whonix Stable Repository"
Hit "Next" three times
Hit "Finish"

`whonixcheck` will run automatically. It should succeed, if it doesn't there may be problems with the virtual bridge set up or the **Galaxy Router**. Make sure openvpn is running and routing traffic on the Galaxy Router.

If `whonixcheck` gives you the message "apt-get reports that packages can be updated", follow the instructions to update.

Delete the logs, delete the bash history, and shutdown. Clone and snapshot.

### Alpine version
Again missing rc-update commands (and more)

## Controller VM setup
### Recommended Debian install options
#### Disk formatting method?
Throughout the guide it says LVM can be a security problem, should we manually create an ext4 partition, encrypt it, and install to that then?

#### Desktop and other package options?
Gnome desktop (Multiverse team will release a slimmed down version of Gnome in the future)
Standard system utilities
SSH server

#### Postinstall
- Change `stretch` to `buster` and delete CD-ROM entries in `/etc/apt/sources.list`. `apt-get update && apt-get dist-upgrade && apt autoremove`
- Add lines for modules `9p`, `9pnet`, `9pnet_virtio` to `/etc/initramfs-tools/modules` and `update-initramfs -u`.
- Add 9p mount lines to `/etc/fstab`, example:

```
multiverse  /media/user/multiverse  9p  trans=virtio,9p2000.L,rw,posixacl,cache=none  0 0
```

- Install other required/useful packages:
  - Basic tools: pass virt-manager vim gnupg tor git
  - Video drivers for video card to be passed through

- Set up Tor ssh:
  - Add a line (or uncomment) in `/etc/tor/torrc`:

  ```
  HiddenServiceDir /var/lib/tor/hidden_service
  HiddenServicePort 22 127.0.0.1:22
  ```
  
  - Restart Tor: `systemctl restart tor`
  - Note onion address for VM: `cat /var/lib/tor/hidden_service/hostname`
  - Make sure sshd is enabled: `systemctl enable ssh`
  - Lockdown SSH (disable root login, review accepted ciphers and versions and remove old/less secure ones, once other computers' keys have been added to `~/.ssh/authorized_keys` disable password login, etc)

### Add instructions for CPU pinning
### Other post-install steps
It seems like there should be more details here. Testing that GPU passthrough works, maybe a brief overview of what you'll be doing later in lockdown section to force the computer to boot into the Controller VM so the user never sees the Host Machine again

#### Testing GPU passthrough
- This method doesn't work because there's no `unbind` file for my GPU: `echo x.xxx.xx.x > /sys/bus/pci/devices/xxxx/driver/unbind && echo "xxxx xxxx" > /sys/.../vfio-bind && ulimit -l (look up unlimited ulimit) && sudo -u user -H virsh start controller.multiverse`
- `lspci -k`, found that `amdgpu` is the module, blacklisted in `/etc/modprobe.d/multiverse.conf`. Made a root session controller VM for testing because won't have permissions issues. Tested it boots normally, removed video and spice, added GPU PCI device, enabled start at boot. Added `virsh list` to `/etc/rc.local` to make sure the qemu session starts. Disabled login manager `systemctl disable lightdm`. Reboot, get host login prompt and dmesg errors

vfio-pci 0000:xx:xx.x: Invalid PCI ROM header signature: expecting .... got ....
vfio-pci 0000:xx:xx.x: BAR0: can't reserve [mem ....... 64bit pref]

- Note: can blacklist the vfio_pci driver in grub at boot to get graphics back for troubleshooting

# machines/universe.router.multiverse/provision.sh
# machines/galaxy.router.multiverse/provision.sh
## mkdirs
Errors when making links in /etc/dhcp and /etc/shorewall because those folders don't exist yet.
## sysctl.d
Repeats, once with a colored "creating symlink..." line, once without

# router issues
## galaxy.router.multiverse startup warnings:
Zone uni is defined as a sub-zone of wan, yet the two zones have no interface in common in /etc/shorewall/hosts (EOF)
** uni is an EMPTY ZONE ** /etc/shorewall/hosts (EOF)

## universe/galaxy dns
galaxy says the nameserver is 9.9.9.9 (/etc/resolv.conf) but idk where it's getting that and also the custom dns servers set in universe, should i be seeing them here or something? I can't remember if that's what I entered in alpine-setup, I thought I did 1.1.1.1 but I restarted setup a couple times. `dig google.com` confirms that I'm using 9.9.9.9 as my server

universe is using 192.168.0.1 as DNS, don't remember if I set that during alpine-setup or not or if it's grabbing it from the router dhcp when it gets an ip address. confirmed with `dig`.

`dig` is available in the package `bind-tools`

### Problem with dns on galaxy with openvpn
After reboot (and enabling openvpn with some vpn file as the config), I have dns problems, /etc/resolv.conf says 

```
# Generated by openvpn for interface tun0
nameserver 10.4.0.1
```

and won't resolve anything, not even the vpn servers. Manually setting the dns in resolv.conf (to 1.1.1.1 or whatever), I can restart openvpn and it works as expected (note it rewrites resolv.conf to 10.4.0.1).


# misc
## qemu storage pools
For some reason when I make new disks in virt-manager, it keeps putting them in the .local folder, even though I've deleted that storage pool multiple times. I've even deleted the .local/share/libvirt/images folder and it just remakes it. Look up how best to use `qemu-img` to make the qcow2 rather than virt-manager (as part of the process of getting away from virt-manager).
`qemu-img create -f qcow2 multiverse-os/images/controller.multiverse.qcow2 60G`

## templates with default qemu
Cloned machines have new (auto generated by qemu) mac addresses, is this desired behavior?

## VM startup order
libvirt will autostart VMs whose definition XML file is symlinked in `.config/libvirt/qemu/autostart`. Startup order is alphabetic, so for starting VMs in specific order, can prefix: ex. 10-universe.router.multiverse.xml, 20-galaxy.router.multiverse.xml... However, renaming the files will confuse virt-manager or virsh if you try to use them to change the autostart settings later. Solutions:

a) Don't count on libvirt to start machines, use a script (for ex. systemd or rc.local)
b) Count on libvirt, and write handling of autostart order prefixes into the multiverse virsh replacement tool
