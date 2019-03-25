# Multiverse OS: Router Controller 
The most recent design for the network controller is one that takes in a 
physical passed through network.

After trying several options, Debian9 base system with two (2) nested Alpine
Linux based routers.

  (1) Galaxy Router 

  (2) Star Router

In the previous designs we had three (3) Alpine Linux routers, but we are 
making the encapsulating Debian9 system Universe which holdes the galaxy and
star router, which makes the metaphor work better. 

**Potential ways to save CPUs and prevent overcommitting CPUs**
Because the current system requires 3 CPUs, this is not viable for most
hardware without overcomitting which is incredibly determimental to security. 

So in the future we want to migrate to a sandbox container style segregation
of routers. This will be done using a modified version of:

    `github.com/google/gvisor`

For the `galaxy` and `star` routers. Ideally we can get this to be a single CPU
containing up to 20+ routers with no issue. Then we can dedicate the rest of
our CPUs to application and service VMs.

_______________________________________________________________________________
## Router Controller Setup
Below is the guide to configuring the network controller on Debian9.

#### Package Configuration

**TODO: Disable apt-get suggestions and recommendations** so we can delete 
things like firefox without chrome being installed automatically.

````
sudo apt-get update
sudo apt-get upgrade

sudo apt-get install virt-manager ovmf libvirt libvirt-daemon qemu qemu-img

sudo apt-get install vim

sudo apt-get remove nano firefox
  five-or-more four-in-a-row gnome-chess gnome-games gnome-klotski
  gnome-mahjongg gnome-mines gnome-nibbles gnome-robots gnome-sudoku
  gnome-taquin gnome-tetravex hitori hoichess iagno iputils-arping
  libgnome-games-support-1-2 libgnome-games-support-common libminiupnpc10
  libmission-control-plugins0 libnatpmp1 libndp0 libqqwing2v5 libteamdctl0
  lightsoff minissdpd mobile-broadband-provider-info polari quadrapassel
  swell-foop tali telepathy-idle telepathy-logger telepathy-mission-control-5
  transmission-common transmission-gtk

````
#### Storage Configuration
Use the user session libvirt configuration, add the following storage
directories:


  (1) base-images
  (2) os-images
  (3) portal-images
  (4) network-controller (the 'portal' folder, which contains this guide and
      the files and config details needed for setup and configuration)

#### Galaxy Configuration
Add a single plan9 share (Driver: Path, Mode: Mapped, Cache: None). 


Path will be to the `galaxy.portal.multiverse` named `galaxy-portal`. Then finalize the 
setup by adding the following line to the newly created VM's `/etc/fstab`:



````
galaxy-portal /mnt/galaxy-portal 9p trans=virtio,9p2000.L,posixacl,rw,cache=none,nofail 0 0
````

Reboot, and if the mount is showing up, run the two scripts that handle
setting up the router.

````
cd /mnt/galaxy-portal
./packages.sh
./provision.sh
````

Next install their open VPN configuration

#### Virtual Bridges On Each Network Controller
Using three bridges, the first one implementing DHCP, and other features which
enable it to function as a router is the `universe0` router. Then the other
two (2) will only establish the subnet. This model enables simple creation f
several galaxy networks isolated with their own transparent proxy routing. 

  (1) universe0
      **subnet** 10.0.0.0/24
      **dhcp** 10.0.0.2-254
      **mac** 00:00:10:00:00:01
      **ip** 10.0.0.1

  (2) galaxy0
      **subnet** 10.1.1.0/24
      **dhcp** disabled
      **mac** 00:00:10:01:01:fe
      **ip** 10.1.1.254

      *note* ideally we disable access to this bridge altogether from the host

  (3) star0
      **subnet** 10.2.2.0/24
      **dhcp** disabled
      **mac** 00:00:10:02:02:fe
      **ip** 10.2.2.25

      *note* ideally we disable access to this bridge altogether from the host
      and if we opt to use whonix it will be 10.152.152.0/24

#### VM-to-Controller Communication
Until better solutions can be implemented using our more reliable VM-to-VM
communication we will utilize two (2) virt serial based communication channels:

    (1) `/var/run/multiverse/portals/galaxy0.router.network0/status.out`

    (2) `/var/run/multiverse/portals/galaxy0.router.network0/virtserial-sock`


_______________________________________________________________________________
#### Universe Router Setup
After configuring the `network.controller.portal` for use as a hypervisor, it 
needs to be then setup as Universe router. This allows us to reduce the 
total number of routers by 1, encapuslating the other routers within the 
unvierse router, which serves the metaphor better while being more efficient. 

Since in almost all cases, we will not be networking directly with the universe
router, this setup should work fine.

Start out by installing the necessary software:

````
sudo apt-get install shorewall
sudo systemctl enable shorewall
````

Some other conveinences to make the usage of Unvierse/Network controller more
fluid.

````
cd ~/.local/share/libvirt
ln -s /media/user/multiverse-portal/ ~/.local/share/libvirt/images
````


_______________________________________________________________________________
#### Galaxy & Star Router Setup
Start by initializing the run folders that will provide us data on each VM
instance. These folders enable copying from the generic 
`network-controller-portal` and then if local modifications are needed, they
can be done without affecting the generic starting point:

**NOTE** This will not work until [`/var/`, `/home/`, `/tmp/`] are on separate 
disks.

````
sudo mkdir -p /var/run/multiverse/portals/galaxy0.router.network0
sudo mkdir -p /var/run/multiverse/portals/star0.router.network0
sudo chown -R user:user /var/run/multiverse
touch /var/run/multiverse/portals/galaxy0.router.network0/status.out
touch /var/run/multiverse/portals/star0.router.network0/status.out
````

Then we initialize the local `/var/` for controller specific data:

````
sudo mkdir -p /var/multiverse/images/os-images
sudo mkdir -p /var/multiverse/images/base-images
sudo mkdir -p /var/multiverse/scripts
sudo chown -R user:user /var/run/multiverse
````

Now copy the data from the `network-controller-portal`:

````
sudo cp /media/user/multiverse-portal/images/base-images/* /var/multiverse/images/base-images/
sudo cp /media/user/multiverse-portal/images/os-images/* /var/multiverse/images/os-images/
sudo cp /media/user/multiverse-portal/scripts/* /var/multiverse/scripts/
sudo chown -R user:user /var/run/multiverse
````


**Setup openVPN client** [future: add remote VPS inverse proxy, i2p, etc]
Download AirVPN configurations, copy one to /etc/openvpn/openvpn.conf

Add to defaults:

````
service openvpn start
rc-update openvpn defaults
````

Then this will transparently proxy the `user-interface.controller.portal`
and the `star0.router.portal`. 


>   [Development] A router-daemon needs preferably watch (not short poll!) for
>   events relating to network devices via `systemd` or `sysV` and restart the
>   service and automatically run through debugging checks.

>   For example, `/etc


#### Install & Configure Star0 Router

> [Development] **whonix does not support ovmf it appears, even after** 
> **reinstalling grub** this is another reason we need to abandon whonix and
> build our own tor based router.

Using the `galaxy0.router.portal` as a model, we can likely get most of the
configuration setup. 

**TODO: It would be nice to move away from shorewall in favor or either a**
**raw sockets options (with setcap +net_admin,etc) to route all traffic**
**BELOW even the iptables (which is what shorewall is an abstraction for)**




_______________________________________________________________________________
#### Connecting the routers to the higher level virbr0's
Using `session` in virt-manager (for now), we can experiment with the 
"Network Interaces" tab, with `brX`, `bondX`, and `vlanX` for connecting
the galaxy router and star router to their respective etherenet slots.


**NEED TO FIND A WAY TO ENSURE CONSISTENT ensX naming!** These are supposed to
be consistent, its the entire point of moving to ensX. It is incorporating BIOS
provided PCI express hotplug slot index numbber.

````
<interface type='bridge'>
 <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
</interface>
# The above XML will translate to a device with the name below:
2: ens3: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 00:00:10:00:00:01 brd ff:ff:ff:ff:ff:ff
````

The attribute `slot` in the `<address>` element is what is defining the `ensX`
value. So we we can consistently get our ensX naming to simplify and streamline
configuration of `network0.controller.portal` network devices. 

So for now we will use the following naming for the three (3) bridges and one
(1) PCI passthrough of the NIC card. The key here is we are moving all the 
network devices, to bus `0x02`, which should be unused in a VM with this light
of a VM. Then we can keep all the networking together with `0x01` standardized
as the PCI passthrough, then `0x02` to `0x04` is the virtual bridges. 

This model provides consistency, and keeps the networking together, resulting
in:

It is somewhat inconsistent, but the ordering stays correct:

	Network 0
        ens1: PCI Passthrough  (LAN/WAN Network) 
        ens2: Universe Network (Subnet: 10.0.0.0/24)
        ens3: Galaxy Network   (Subnet: 10.1.1.0/24)
        ens4: Star Network     (Subnet: 10.2.2.0/24)

	Network 1
        enp4s1: PCI Passthrough  (LAN/WAN Network) 
        enp4s2: Universe Network (Subnet: 10.0.0.0/24)
        enp4s3: Galaxy Network   (Subnet: 10.1.1.0/24)
        enp4s4: Star Network     (Subnet: 10.2.2.0/24)


So tried again using `0x05` as the bus, so all the devices are definitely
segregated. Still got `enp5sX` and `enp2sX`.

So I noticed, that network1 was consistently offset by 2, so I made it `0x03`
and network0 `0x05` and got them both to be `enp5sX` successfuly.

This is done using the following XML:




````
    <hostdev mode='subsystem' type='pci' managed='yes'>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x01' function='0x0'/>
    </hostdev>

    <interface type='bridge'>
      <mac address='00:01:10:00:00:01'/>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x02' function='0x0'/>
    </interface>
    <interface type='bridge'>
      <mac address='00:01:10:01:01:fe'/>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x03' function='0x0'/>
    </interface>
    <interface type='bridge'>
      <mac address='00:01:10:02:02:fe'/>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x04' function='0x0'/>
    </interface>

````




Wireless networking should be pushed off into its own bus, probably bus 6, but
that is not yet decided.

_______________________________________________________________________________
## Notes









