# (Old) Multiverse OS: Router Setup & Networking Guide

## Host Machine Networking Information
In general the goal is to disable all networking inside the bare-metal **Host Machine**.

Disable all the networking scripts to avoid lengthy waits for devices or scripts to run unncessarily

```
systemctl disable NetworkManager
systemctl disable networking
systemctl stop NetworkManager
systemctl stop networking
```

The following tunnels are notes on potentially useful upgrades but are not actively used in the current implementation of Multiverse OS.

**Multicast tunnel**
A multicast group is setup to represent a virtual network. Any VMs whose network devices are in the same multicast group can talk to each other even across hosts. This mode is also available to unprivileged users. There is no default DNS or DHCP support and no outgoing network access. To provide outgoing network access, one of the VMs should have a 2nd NIC which is connected to one of the first 4 network types and do the appropriate routing. The multicast protocol is compatible with that used by user mode linux guests too. The source address used must be from the multicast address block.

```
...
<devices>
  <interface type='mcast'>
    <mac address='52:54:00:6d:90:01'/>
    <source address='230.0.0.1' port='5558'/>
  </interface>
</devices>
...
```

**TCP tunnel**
A TCP client/server architecture provides a virtual network. One VM provides the server end of the network, all other VMS are configured as clients. All network traffic is routed between the VMs via the server. This mode is also available to unprivileged users. There is no default DNS or DHCP support and no outgoing network access. To provide outgoing network access, one of the VMs should have a 2nd NIC which is connected to one of the first 4 network types and do the appropriate routing.

```
...
<devices>
  <interface type='server'>
    <mac address='52:54:00:22:c9:42'/>
    <source address='192.168.0.1' port='5558'/>
  </interface>
  ...
  <interface type='client'>
    <mac address='52:54:00:8b:c9:51'/>
    <source address='192.168.0.1' port='5558'/>
  </interface>
</devices>
...
```

**UDP unicast tunnel**
A UDP unicast architecture provides a virtual network which enables connections between QEMU instances using QEMU's UDP infrastructure. The xml "source" address is the endpoint address to which the UDP socket packets will be sent from the host running QEMU. The xml "local" address is the address of the interface from which the UDP socket packets will originate from the QEMU host. Since 1.2.20

```
...
<devices>
  <interface type='udp'>
    <mac address='52:54:00:22:c9:42'/>
    <source address='127.0.0.1' port='11115'>
      <local address='127.0.0.1' port='11116'/>
    </source>
  </interface>
</devices>
...
```


**NOTE** This is very handy to know about, `/sys/class/net/` provides all the information in a filesystem style. This is much easier than chaining sed grep and so on. This is worth mentioning in guides. Multiverse should take advantage of this.
______
## Universe Router
The universe router is analogous to QubesOS `net-sys` router, it is the router that has the network cards passed using PCI passthrough.

**General Configuration**

To begin we will provide hostnames for the router to make accessing it in the local networks easier.

```
cat "127.0.0.1         router.universe0 universe0 uni0 localhost.localdomain localhost" > /etc/hosts
```

Next the `motd` and `issue` files will be updated to indicate which Multiverse OS router is being accessed which is important since there are three or more routers in the Multiverse OS system.

```
# [ROUTER] Update the motd
  cat "Multiverse OS Router  [ router0.universe.mv ]" > /etc/motd

# [ROUTER] Update the issue
  cat "Multiverse OS Router  [ router0.universe.mv ]" > /etc/issue
  cat "[ Built using Alpine Linux 3.4 ]" > /etc/issue
```



**DEVELOPMENT** Another very important feature that needs to be introduced soon as possible is the ability for the **Router VM** to support combining multiple internet connections and serving these connections into the deeper Multiverse OS networks as a single connection. For example, several wifi connections.  

Packages required by Multiverse OS **Router VM**:

```
  apk update
  apk add ca-certificates
  apk add openssl
  apk add vim
  apk add shorewall
```

Some default packages need to be deleted, eventually more should be added to this list to lockdown the **Router VM** and reduce the attack surface as much as possible:

```
# Remove SSH, this should only be managed other ways, preferably a custom Multiverse OS protocol
# to avoid generic malware from attacking
apk del ssh
```

**Configuring Shorewall Firewall**

In order to configure Shorewall Firewall `rc-update` needs to be used but in addition, `/etc/shorewall/shorewall.conf` needs to be modified:

```
rc-update add shorewall default
```

Then modify `/etc/shorewall/shorewall.conf` on the line with `STARTUP_ENABLED=No` must be changed to `Yes`.

The configuration file `/etc/shorewall/masq` does not exist by default, so it must be manually created:

```
touch /etc/shorewall/masq
```

**TODO** Configure shorewall, likley best handled with a deploy shared folder that is taken out when complete.

```
cat "lan               eth0:192.168.1.0/24                   -" >> /etc/shorewall/hosts
```


**Configuring DHCPd**
DHCPd comes with an example configuration file, and copying this and modifying it is a good starting point:

```
cp /etc/dhcp/dhcpd.conf.example /etc/dhcp/dhcpd.conf
```

```
cat "auto lo" > /etc/network/interfaces
cat "iface lo inet loopback" >> /etc/network/interfaces
cat "" >> /etc/network/interfaces
cat "auto eth0" >> /etc/network/interfaces
cat "iface eth0 inet static" >> /etc/network/interfaces
cat "      address 10.1.1.1" >> /etc/network/interfaces
cat "      netmask 255.255.255.0" >> /etc/network/interfaces
cat "      broadcast 10.1.1.255" >> /etc/network/interfaces
cat "      network 10.1.1.0" >> /etc/network/interfaces
cat "      gateway 10.1.1.1" >> /etc/network/interfaces
cat "      up route del -net default gw 10.1.1.1 netmask 0.0.0.0" >> /etc/network/interfaces
cat "" >> /etc/network/interfaces
cat "auto eth1" >> /etc/network/interfaces
cat "iface eth1 inet dhcp" >> /etc/network/interfaces
```

```
# [ROUTER] Create the symbolic to relevant config files to make maintenance easier, eventually these should be done from the Multiverse OS persistent share. This way config files can just be quickly deployed and easily edited.
  ln -s /etc/dhcp/dhcpd.conf ~/dhcp-dhcpd.conf
  ln -s /etc/shorewall/hosts ~/shorewall-hosts
  ln -s /etc/shorewall/interfaces ~/shorewall-interfaces
  ln -s /etc/shorewall/masq ~/shorewall-masq
  ln -s /etc/shorewall/policy ~/shorewall-policy
  ln -s /etc/shorewall/rules ~/shorewall-rules
  ln -s /etc/shorewall/zones ~/shorewall-zones
  ln -s /etc/shorewall/shorewall.conf ~/shorewall-shorewall.conf
  ln -s /etc/hosts ~/hosts
  ln -s /etc/motd ~/motd
  ln -s /etc/issue ~/issue
  ln -s /etc/network/interfaces ~/network-interfaces
  ln -s /etc/sysctl.d/00-alpine.conf
```

**DEVELOPMENT** MAC addresses should be monitored, automatically generated and regularly changed.

Below is the DHCPd configuration

```
# [ROUTER] Configure DHCP server. Define the DNS server #{DNS_SERVER}, for router.galaxy - it should be a good trusted European privacy non-profit DNS.

# DNS_SERVER=["85.214.20.141","194.150.168.168","213.73.91.35"]

#This ddns-update-style line is required for the routing to work, do not overlook it.
dd-update-style interim;

subnet 192.168.1.0 netmask 255.255.255.0 {
}

subnet 10.1.1.0 netmask 255.255.255.0 {
  option routers 10.1.1.1;
  option subnet-mask 255.255.255.0;
  option broadcast-address 10.1.1.255;
  option domain-name-servers #{DNS_SERVER};
  option domain-name-servers #{DNS_SERVER};
  range 10.1.1.2 10.1.1.254;
}

subnet 10.2.2.0 netmask 255.255.255.0 {
}

host firewall.universe0.mv {
  option host-name "firewall";
  hardware ethernet {MAC_ADDRESS};
  fixed-address 10.2.2.2;
}
```




______



# [HOST] For router.universe the network virsh net-edit files need to be updated to switch the ip address to *.*.*.254 for every network. Eventually it may be best to just unplug entirely. But for now we just rely on not routing that direction, then shorewall disabling all incoming and only allow ssh out. Eventually all talking needs to be moved to VirtIO custom networking.


```
# [ROUTER] Rebuild the /etc/sysctl.d/00-alpine.conf file for routing and disable ipv6 because right now that is a security hazard until there is better Tor support
cat "net.ipv4.ip_forward = 1" > /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.tcp_syncookies = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.default.rp_filter = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.send_redirects = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.accept_source_route = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.conf.all.rp_filter = 1" >> /etc/sysctl.d/00-alpine.conf
cat "net.ipv4.ping_group_range=999 59999" >> /etc/sysctl.d/00-alpine.conf
cat "kernel.panic = 120" >> /etc/sysctl.d/00-alpine.conf
```

```
rc-update add dhcpd default 
```

```
cat "RESOLV_CONF=\"NO\"" >> /etc/udhcpc.conf
```

**TODO** Use customized hardware addresses to obfuscate the use of QEMU, currently all the 54:52... hardware addresses are dead giveaway to anyone snooping that its a QEMU/kvm setup.


# [ROUTER] Configure shorewall to route traffic
**TODO** Configure shorewall, likley best handled with a deploy shared folder that is taken out when complete.

```
cat "wan               ipv4" >> /etc/shorewall/zones
cat "lan:wan           ipv4" >> /etc/shorewall/zones
cat "uni               ipv4" >> /etc/shorewall/zones
```


```
cat "uni               eth0:10.1.1.0/24                      -" >> /etc/shorewall/hosts
cat "uni               eth0:10.1.1.0/24                      -" >> /etc/shorewall/hosts
cat "lan               eth0:192.168.1.0/24                   -" >> /etc/shorewall/hosts
```

**TODO** eth1 should be whatever is the pci-passthrough. In the case of one of the machines we have 4 network devices. So determine which ones are working then use those as the wan.

```
cat "wan eth1          routefilter,tcpflags,logmartians,nosmurfs,sourceroute=0" >> /etc/shorewall/interfaces
cat "uni eth0          dhcp,routefilter,tcpflags,logmartians,nosmurfs" >> /etc/shorewall/interfaces
```

```
cat "eth1 0.0.0.0/0" >> /etc/shorewall/masq
```

```
cat "fw  all ACCEPT" >> /etc/shorewall/policy
cat "uni lan ACCEPT" >> /etc/shorewall/policy
cat "uni wan ACCEPT" >> /etc/shorewall/policy
cat "all all DROP"   >> /etc/shorewall/policy
```

```
cat "##" >> /etc/shorewall/rules
cat "##" >> /etc/shorewall/rules
cat "## Voyager Port Forwarding For LAN Services" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Transmission RPC/WebUI" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100    tcp 9091" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# SSH" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 22" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Bittorrent" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 51413" >> /etc/shorewall/rules
cat "" >> /etc/shorewall/rules
cat "# Samba LAN" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 139" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 445" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 137" >> /etc/shorewall/rules
cat "# DNAT lan    uni:10.1.1.100   tcp 138" >> /etc/shorewall/rules
```



**NOTE** Oddly /etc/shorewall/masq was not included so this needs to be initialized with the prefix comments to match the rest of the files.

```
cat "wan eth0          routefilter,tcpflags,logmartians,nosmurfs"      >> /etc/shorewall/interfaces
cat "gal eth1          dhcp,routefilter,tcpflags,logmartians,nosmurfs" >> /etc/shorewall/interfaces
cat "vpn tun0          nosmurfs,tcpflags"                              >> /etc/shorewall/interfaces
```

```
cat "tun0 0.0.0.0/0" >> /etc/shorewall/masq
cat "eth0 0.0.0.0/0" >> /etc/shorewall/masq
```

```
cat "gal vpn ACCEPT" >> /etc/shorewall/policy
cat "gal lan DROP"   >> /etc/shorewall/policy
cat "gal uni DROP"   >> /etc/shorewall/policy
cat "air gal DROP"   >> /etc/shorewall/policy
cat "all all DROP"   >> /etc/shorewall/policy
```

```
cat "wan             ipv4" >> /etc/shorewall/zones
cat "lan:wan         ipv4" >> /etc/shorewall/zones
cat "uni:wan         ipv4" >> /etc/shorewall/zones
cat "gal             ipv4" >> /etc/shorewall/zones
cat "vpn             ipv4" >> /etc/shorewall/zones
cat "air:vpn         ipv4" >> /etc/shorewall/zones
```

```
# [ROUTER] Add a MultiverseOS persistence share. This can later be used to supply config files to simplify the maintence and setup process.
cat "Multiverse /mnt/mv         9p   trans=virtio,9p2000.L,rw,posixacl,cache=none  0 0" >> /etc/fstab
```

**DEVELOPMENT** Remove the /dev/cdrom, /dev/usb, and the last drive. 

**DEVELOPMENT** Should disable USB and other kernel modules to limit the attack surface.

**DEVELOPMENT** Setup ephemeral status

**DEVELOPMENT** The template doesn't require a password, but after setup is done, before the ephemeral status is set, then you just need to set the password and store it in the pass-store

Before shutdown and templating, clear the history, remove the logs, clean the machine for general purpose use. 

```
history -c 
```

Shutdown the machines, take snapshots of each machine to enable ephemeral state, clone each machine for use as a template. Always keep a copy beyond the snapshots, since snapshots are stored inside the qcow2, if the file becomes corrupt, they are also corrupted and do little to help.


[?1000h[?1049h[?1h=[1;20r[?12;25h[?12l[?25h[27m[23m[m[H[2J[?25l[20;1H"segregated-from-readme.notes.md" 164L, 7005C[1;1HSetup passthrough by enabling ioummu, enabling vfio-pci (vfio is for 4.0+ kerneles and pci-stub is for previous versions).

Then echo the lspci -n 00:00.00.0 number into unbind. Then pass 0000 0000 number into the vfio-set bind, which creates a /dev/vfio/##.

Before (or after) the networking is setup, qemu must be configured to allow access to the bridges to unprividedged users on the host. This  [6;1His important because it prevents breakouts from immediately acheiving root access on the host. That you can edit /etc/qemu/bridge.conf to aa[7;1Hdd allow virbr0.

[35m`[msudo virsh net-add multiverse[35m`[m

Then paste in below, switch out the macs. The DNS is currently routing to google, it would be better to route DNS requets to Tor as it provv[12;1Hides a cheap, easy acccess distributed DNS setup. This could be done at the top router level.

I decided on a cosmic naming scheme for my setup. Multiverse OS has 3 network levels:

[38;5;130m1.[m Universe-0 - Universe prime is the computer which you operate, the one you are actively at. Others will start with Universe-1 and so on..[17;1H This is the PCI passthrough level, local area network is accessible. Good for torrent machines (depending on region and stealthiness of yoo[18;1Hur torrent site of choice. A simple VM with torrent client and samba shares for the media center goes here.[20;122H1,1[11CTop[1;1H[?12l[?25h[?25l[20;1HType  :quit<Enter>  to exit Vi[20;32H[K[20;122H1,1[11CTop[1;1H[?12l[?25h[?25l[20;112H[A[1;1H[20;112H  [1;1H[?12l[?25h[?25l[20;122H[K[20;122H1,1[11CTop[1;1H[?12l[?25h[?25l[20;112H~@k[1;1H[20;112H   [1;1H[?12l[?25h[?25l[20;112H^[[1;1H[20;112H  [1;1H[20;112H^[[1;1H[20;112H  [1;1H[?12l[?25h[?25l[20;112H:[1;1H[20;1H[K[20;1H:[?12l[?25hq[?25l[?12l[?25h[?25l[?1000l[20;1H[K[20;1H[?1l>[?12l[?25h[?1049lSetup passthrough by enabling ioummu, enabling vfio-pci (vfio is for 4.0+ kerneles and pci-stub is for previous versions).

Then echo the lspci -n 00:00.00.0 number into unbind. Then pass 0000 0000 number into the vfio-set bind, which creates a /dev/vfio/##. 

Before (or after) the networking is setup, qemu must be configured to allow access to the bridges to unprividedged users on the host. This is important because it prevents breakouts from immediately acheiving root access on the host. That you can edit /etc/qemu/bridge.conf to add allow virbr0.

`sudo virsh net-add multiverse`

Then paste in below, switch out the macs. The DNS is currently routing to google, it would be better to route DNS requets to Tor as it provides a cheap, easy acccess distributed DNS setup. This could be done at the top router level.

I decided on a cosmic naming scheme for my setup. Multiverse OS has 3 network levels:

1. Universe-0 - Universe prime is the computer which you operate, the one you are actively at. Others will start with Universe-1 and so on. This is the PCI passthrough level, local area network is accessible. Good for torrent machines (depending on region and stealthiness of your torrent site of choice. A simple VM with torrent client and samba shares for the media center goes here.

2. Galaxy-0 - Galaxy prime is the network layer which is protected, all traffic is isolation proxyified with an openvpn connection. This can be your normal internet traffic for searching or when something is very tedious to do with Tor. This traffic can not access the lan.

3. Sol-0 The safest place, close to Earth, all traffic passes through Tor. This traffic can not access galaxy or lan networks. It is isolated, but several workstations can be behind the same whonix box and serve stuff to galaxy or lan via onion services. 

*Galaxy and Sol could be combined, and a switch could switch between VPN or Tor but never allow direct access to Universe. Multiple second level Firewall/Proxy VMs could be deployed below Universe. Multiple Universe routers could be deployed if the server has multiple physical NIC cards to assign. The initial server has three NIC cards to distribute, but these are all assigned currently to universe0*

Modify the rc.local to automatically do this. I have seen scritps that bind everything, it would be nice to find one that would bind essentially all ethernet,wireless,etc devices.

`/etc/rc.local`

````
echo "0000:03:00.0" > /sys/bus/pci/devices/0000\:03\:00.0/driver/unbind
echo "1969 e091" > /sys/bus/pci/drivers/vfio-pci/new_id

exit 0
````

When I do `lspci -n` I get

03:00.0 0200: 1969:e091 on the same line that doing just lspci showed the NIC card. Just match the 03:00.0 part. 

One for the PCI passthrough, one for the proxy/firewall VM, and one for whonix. This will let you have fine grain control over how your VM accesses the internet.

`/etc/qemu/bridge.conf`


````
allow virbr0
allow virbr1
allow virbr2
````


 That lets the unprividedged user create devices connected to it. For alpine linux, which I used for my routers. I think it was a better choice than say ddWRT or a larger OS. 

Setup the alpine box, sys-net needs a static route set.



##### net-firewall, the firewall and proxy server

This is where your openvpn client will be setup. Ideally you are connecting to a server you rent and put a VPN on yourself or you are using a reputable VPN seller that takes Bitcoin. The whonix box will route its traffic through this to hide Tor access from your ISP.

Alpine linux does not just pick up the routes, so you need to manually define them in the interfaces file. This is assuming your LAN is 192.168.1.1, which is pretty standard. 

`/etc/network/interfaces`

````
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet dhcp
	#up ip route add net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
	dns-nameservers 8.8.8.8 8.8.4.4
	hostname sys-firewall
	up route add -net 192.168.1.0 netmask 255.255.255.0 gw 10.1.1.1
	up route add -net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
````


###++++

Remaining steps

# openvpn on firewall/proxy alpine linux VM

# setup whonix box, route through firewall/proxy

# firewall the host, only outbound SSH connections allowed. on sys-firewall and sys-net disable access to it.

# 

###++++






# automatically setup the DNS at universe-0's router, then all names should go there or the router below it. This way connecting like ssh user@world-2 will work 

-- automatically remove meta data from images in a specific folder? optional all 

================================================================
writing script to deploy server and ehre are notes i made while figuring out which tool to use to manage accounts, get pci info for pci pass throughs and other junk
=======================================================================

##### Notes on provisioning Alpine Linux for use as a router
Setup Alpine, install vim, install shorewall, copy shorewall configs, setup dhcp, enable shorewall

Alpine Linux does not just pick up the routes (depends on the settings of the virbrX), so you need to manually define them in the interfaces file. This is assuming your LAN is 192.168.1.1, which is pretty standard. 

`/etc/network/interfaces`

```
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet dhcp
	#up ip route add net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
	dns-nameservers 8.8.8.8 8.8.4.4
	hostname sys-firewall
	up route add -net 192.168.1.0 netmask 255.255.255.0 gw 10.1.1.1
	up route add -net 0.0.0.0 netmask 0.0.0.0 gw 10.1.1.1
```

#### Utility VMs
I put some of this into the Utility VM section it was sitting under, thought there might be details left that you'd want to keep.
Utility VMs should provide an interface for the controller VM. That way for example the controller VM for the torrent VM receives enhancements like terminal commands to add magnet files by url or hash. 

Each registered Utility VM should expose a interface. Optional but ideally each running utility provides access to its data without requiring /dev/input (mouse and keyboard) type input.

It could be a HTTP page that contains scripts and possible interface.

**Lan Maker VM**

gpio (to create a utility VM for connecting to raspberry pis)
https://github.com/platinasystems/gpio  -- minimalist lib

**Captive Portal VM**

Simple VM to login into captive portal, attaches to router.universe (LAN)

Kicks up a desposable firefox (also has a chromium browser incase the portal is poorly written)

**IP-over-DNS server** Can be running so you can obtain internet access over captive portals

**Backup and Archive VM**

A VM which easily configures, using YAML files (and a UI to modify YAML files, define which folders get uploaded where) to backup via:
* rsync to remote servers (ssh)
* spideroak
* github
* ?

**Example** *.yaml configuration*

```
servers:
  server:
    # Optional Name
    name: "LA"
    host: 125.125.125.1
    type: rsync
    key: ~/.ssh/chance.pub
    directores: ["~/keys", "~/pictures"]
```

# Multiverse Host Kernel Networking bypass



## template-router.wan.universe0
##
<domain type='kvm'>
  <name>template-router.universe0.mv</name>
  <uuid>48d4f548-eef1-4818-ba58-6388f29168dd</uuid>
  <memory unit='KiB'>524288</memory>
  <currentMemory unit='KiB'>524288</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <os>
    <type arch='x86_64' machine='pc-i440fx-2.6'>hvm</type>
  </os>
  <features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
  </features>
  <cpu mode='custom' match='exact'>
    <model fallback='allow'>Haswell</model>
  </cpu>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/bin/kvm</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source file='/home/user/.local/share/libvirt/images/universe.router.mv-clone.qcow2'/>
      <target dev='hda' bus='ide'/>
      <boot order='2'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x2'/>
    </controller>
    <controller type='virtio-serial' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='ide' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <interface type='bridge'>
      <mac address='52:54:00:92:ba:73'/>
      <source bridge='virbr0'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target port='0'/>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <channel type='spicevmc'>
      <target type='virtio' name='com.redhat.spice.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='spice' autoport='yes'>
      <listen type='address'/>
      <image compression='off'/>
    </graphics>
    <video>
      <model type='qxl' ram='65536' vram='65536' vgamem='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <hostdev mode='subsystem' type='pci' managed='yes'>
      <source>
        <address domain='0x0000' bus='0x03' slot='0x00' function='0x0'/>
      </source>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </hostdev>
    <hostdev mode='subsystem' type='pci' managed='yes'>
      <source>
        <address domain='0x0000' bus='0x04' slot='0x00' function='0x0'/>
      </source>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x08' function='0x0'/>
    </hostdev>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
</domain>

## template-router.firewall.universe0
##
<domain type='kvm'>
  <name>template-firewall.universe0.mv</name>
  <uuid>56e15fbf-2221-4134-ade2-216bc621531c</uuid>
  <memory unit='KiB'>524288</memory>
  <currentMemory unit='KiB'>524288</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <os>
    <type arch='x86_64' machine='pc-i440fx-2.6'>hvm</type>
  </os>
  <features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
  </features>
  <cpu mode='custom' match='exact'>
    <model fallback='allow'>Haswell</model>
  </cpu>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/bin/kvm</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source file='/home/user/.local/share/libvirt/images/universe.router.mv-clone-clone-clone.qcow2'/>
      <target dev='hda' bus='ide'/>
      <boot order='2'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x2'/>
    </controller>
    <controller type='virtio-serial' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='ide' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <filesystem type='mount' accessmode='mapped'>
      <source dir='/home/user/multiverse-os/persistent/firewall.universe0'/>
      <target dir='MultiverseOS'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x08' function='0x0'/>
    </filesystem>
    <interface type='bridge'>
      <mac address='52:54:00:33:7d:c3'/>
      <source bridge='virbr0'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <interface type='bridge'>
      <mac address='52:54:00:fa:09:e2'/>
      <source bridge='virbr1'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target port='0'/>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <channel type='spicevmc'>
      <target type='virtio' name='com.redhat.spice.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='spice' autoport='yes'>
      <listen type='address'/>
      <image compression='off'/>
    </graphics>
    <video>
      <model type='qxl' ram='65536' vram='65536' vgamem='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
</domain>

##
##
