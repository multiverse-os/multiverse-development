# Multiverse OS: Network (Router) Controller 
We are using Alpine Linux for a variety of reasons for the Router VMs but a 
major reason is the security implications of a monoculture cluster for 
Multiverse OS. By mixxing two different operating systems one based on Debian 
and the other on Hardened Gentoo, it makes it much more difficult and unlikely 
for a single exploit to work across the entire Multiverse OS cluster.

**On the HOST**
Multiverse OS installation starts out by configuring the HOST machine. This is
done by taking a vanilla debian linux install and turning it into a slim
hypervisor, and eventually will evolve patching the kernel and so on. 

For now, while we still work the the structure of Multiverse OS, this mostly
involves properly setting up all the virtual machines and nested virtual
machines.

Modify `/etc/qemu/bridge.conf` to allow access to host bridges:

````
allow net0br0
allow net0br1
allow net0br2
allow net1br0
allow net1br1
allow net1br2
````

Then update the permissions of the new bridges:

````
#!/bin/bash
chown -R root:libvirt /usr/lib/qemu/
chmod 4750 /usr/lib/qemu/qemu-bridge-helper
````

After this is done, a base HD and a portal QCOW2 HD need to
be mounted with the 3 subnets to each network(interface)controller.

**Create the `RUN` folder in our `/var/multiverse` structure**
To hold our sockets, status files, pie files, logs, and eventually everything related to the VM running. 

````
mkdir -p /var/multiverse/run/portals/network0.portal
mkdir -p /var/multiverse/run/portals/network1.portal
mkdir -p /var/multiverse/run/portals/user0.debian9.portal
mkdir -p /var/multiverse/run/portals/user0.ubuntu17.portal
# Since we want to move away from libvirt and eventually
# abandon it completely, lets use kvm whenever possible
# for our groups.
chown -R user:kvm /var/multiverse/run/ 
````

**Basic VM-to-HOST Communication**
So we have successfully got VERY basic communication to the HOST machine to
simplify the boot process on the host machine, alllowing us
to report to the HOST machine on successful boots, internet established, 
and whatever else we need. 

````
echo "{
  'timestamp': 'Time.now',
  'action': 'boot',
  'status': 'complete'
}" > ttyS0
````
_______________________________________________________________________________
## Portals
For now portals will need to stay p9 shares, and not qcow2 isos. The change to
using images will not make sense until the base images are actually generic
and not specific.

_______________________________________________________________________________
## Router Controller Setup
We have two types of Multiverse OS `network.controller`.

[Provision & Configure][Debian9 Based Router Controller]
**METHOD 1** 
*Galaxy and Star virtual machines, using the top level as universe*

Create `universe0` bridge, make it dhcp (`10.0.0.2`-`10.0.0.254`).

We will create two virtual machines, one (1) will be the router of the galaxy
network, and two (2) will be the router for the star network. 

While the top level encapsulating VM will handle routing for the universe0 
network.




**METHOD 2**
*Three virtual machines*

Start by installing the basic packages.

````
sudo apt-get install virt-manager
sudo apt-get install vim
sudo apt-get install libvirt
sudo apt-get install libvirt-daemon
sudo apt-get install qemu qemu-img ovmf
````

Remove some of the default packages

````
sudo apt-get remove nano
````

Then add the `libvirt-daemon` to startup via `systemctl`

````
systemctl enable libvirtd
````

## Adding user groups 
Next we need to provide the correct permissions to the user by appending the
VM related groups.

````
sudo usermod -a -G libvirt user
sudo usermod -a -G qemu user
sudo usermod -a -G kvm user
sudo usermod -a -G libvirt-qemu user
````

## Configure other user permissions
[Configuring userspace memory limits]
Next the memory limit needs to be removed so user session VMs can access the whole machine's memory. Two lines need to be added to the `/etc/security/limits.conf file:

````
@kvm             soft    memlock         -1
@kvm             hard    memlock         -1
````

[Add `/etc/qemu/bridge.conf` file]
Now the newly created bridges need to be exposed to userspace by listing them in `/etc/qemu/bridge.conf`.


````
mkdir /etc/qemu/
echo -e "allow virbr0\nallow virbr1\nallow virbr2" > /etc/qemu/bridge.conf
# The -e flag allows for the interpretation of \n line breaks so that each allow
# {bridge_name} is on its own line.
````

## Configuring LIBVIRT

Add the three virtual bridges to the root session `qemu:///system` 




______________________________________________________________
[Provision & Configure][Alpine3.7 Based Router Controller]
Start by getting the right `setup alpine` configuration and the correct QEMU XML for the server. Then use the below to configure the server.


````
apk update
apk upgrade
apk add vim
# Edit the repositories 

apk add libvirt
apk add libvirt-daemon
apk add qemu
apk add qemu-system-x86_64
apk add qemu-img
apk add ovmf

rc-update add libvirtd
# Check /var/run/libvirt for the socket files
# Ensure that openssh is also starting

# Modify the /etc/fstab to add the p9 shares


## add `communiAty` repo to instlal usermod
## TODO: Dont need to if we just save the modifications made to /etc/group
# shadow has usermod
apk add shadow 

adduser user
usermod -a -G kvm user
usermod -a -G qemu user
usermod -a -G libvirt user
usermod -a -G netdev user

````


````[cmd][cat /etc/group]
...
kvm:x:34:kvm,user
...
libvirt:x:102:user
...
qemu:x:36:user
user:x:1000:
````


[Configuring userspace memory limits]
Next the memory limit needs to be removed so user session VMs can access the whole machine's memory. Two lines need to be added to the `/etc/security/limits.conf file:

````
@kvm             soft    memlock         -1
@kvm             hard    memlock         -1
````

[Add `/etc/qemu/bridge.conf` file]


Now the newly created bridges need to be exposed to userspace by listing them in `/etc/qemu/bridge.conf`.


````
mkdir /etc/qemu/
echo -e "allow virbr0\nallow virbr1\nallow virbr2" > /etc/qemu/bridge.conf
# The -e flag allows for the interpretation of \n line breaks so that each allow
# {bridge_name} is on its own line.
````

## Create the `run` folder for network controller

````
mkdir -p /var/multiverse/run/portals/universe.router
mkdir -p /var/multiverse/run/portals/galaxy.router
mkdir -p /var/multiverse/run/portals/star.router
````

**Back on the HOST**
In development, to simplify configuration of the 
Multiverse OS routers until we have better automation scripts
because we have a stable system to build on (building a system
to build a system currently). 

The below gsettings line modifies the ACTIVE configured connections
it does NOT modify the options in the '`File` > `Add Connection`' menu

````
gsettings set org.virt-manager.virt-manager.connections uris:
 
"[
  'qemu+ssh://root@10.0.0.10/system?socket=/var/run/libvirt/libvirt-sock',
  'qemu+ssh://root@10.0.0.10/session?socket=/var/run/libvirt/libvirt-sock',
  'qemu:///session',
  'qemu://system'
]"
````

So this will end up with virt-manager having 4 options, two local root/user
and root/user at 10.0.0.10. 

This is not ideal and is only being used to setup the VMs inside the new
Router Controller VM. Once we can get a development machine up, we should
be reimplementing this using scripts so that the entire process can be
done WITHOUT `virt-manager`. 





________________________________________________________________
## Notes

  * Should auto set `/etc/network/interfaces` based on the mac addresses
    should just be auomatic.


 * [!!] Experiment with removing the ipaddress section of the virbrX, and 
    disable the virbr-NIC and TAP. If this does not work, try writing packets
    to files in `P9` or `/dev/SHM` or ideally a `ring buffer with ...(something i cant remember)`. Then read the packets and feed them into the device. You
   can echo '{packet data} > /dev/nic0 and it will send the packet over the device
   so maybe we can just expose the device for each on the other. by mounting it
   remember the ideal situation is one where we avoid unecessary copies.

  * Implemement virtaul /dev/tmp on host to pass to VMs

  * COme up with good sources of entropy served over OHT

  * [!] All VMs should be real-time (at least emulated) so when we start switching to
using arm64 and releasing hardware we will be ready.

  * [!] Try to convert C code in QEMU to pure Go this way we can make different versions that are lighter weight and only include the features we want. And we 
can make security the priority in development. Can start with C->Go transpilers
and slowly work towards it.

  
  






