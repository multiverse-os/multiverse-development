to# Host Lockdown

## Multiverse OS ENV, gnome settings, etc 
this should all be consoldiated, and put into a fast DB
then, the data can be oputput in legacy format, ie xession-errors
or .selected_editor. But it should be stored in CBOR, in binary,
then just have many ways of outputting it. T

Then the .selected_editor file, can be described as a query from 
the DB. 

**and the ENV too can just be described as a query from the**
general settings. This would be awesome, fast, light weight. 
and super advanced while keeping legacy shit around for oldheads

IN ADDITON, things that ARE scripts, like .bashrc, .logout are
confusing, its hard to tell what is a script, what is a setting. 

Its time to build a system that holds these concepts better.

__An inventory type system for each script style, like prompt, bashrc__
so its very fcleary what is availble to script, be able to quickly
and easily customize these scripts, switching out with other 
comm,unity items


## AUTO COMBINE KEYS like .ssh, .gnupg

This should be done via overlay, or the legacy overlay
should be stored, so it can be made visible, 

**like the files should be loose in the Multiverse OS FS**
**then we just use GraphDB to store the various maps**
**like one could be map of traditional fs hiearchy** 
**or can have a map where its organized in different ways**
**these maps ideally are defined by SCRIPTING LANGUAGE**
**so it doesnt ever have to be manually sorted, and it can**
be shared!



## Host HOME Configuration

`~/.selected_editor` should be moved into ~/.config/... 

And probably should just be merged with other similar empty config 
files so it can be PUT INTO A DATABASE! Just like wtih all the other
.config data too!

> # Generated by /usr/bin/select-editor
> SELECTED_EDITOR="/bin/nano"

This sucks, should be JSON or CBOR or something. Can always just
print the data in human readable. WHY STORE IT HUMAN READBLE
we are storing it because tis CONFIG , as in COMPOUTERS READ IT 

lets fix this shit, this shit both makes linux suck and hard to use


## gvfs-metadata seems insane way of handling this data, why not just use a real db
it would let programmers to hit it, and uis to use the data. etc



## GNome shell for USER on HOST

* `share/gnome-shell/application_state`

very intersting, never knew about this, should be incorproated into Multiverse dAEMON
it gives window info in xml. 

		<?xml version="1.0"?>
		<application-state>
		  <context id="">
		    <application id="org.gnome.FileRoller.desktop" open-window-count="0" 
                     core="0" last-seen="1535708846"/>
		    <application id="org.gnome.Nautilus.desktop" open-window-count="1" 
                     score="281" last-seen="1535709002"/>


## Host USER home files


  [!][GPG .gpg folder]
  * GPG in c folder structure is garbage. Lets just rewrite this whole thing, merge it into 
    Scarmble key system, and then we can have a consistent API, consistent files, configs, 
    consistent backup, etc


    **never used this itneresting featuer**
# List of allowed ssh keys.  Only keys present in this file are used
# in the SSH protocol.  The ssh-add tool may add new entries to this
# file to enable them; you may also add them manually.  Comment
# lines, like this one, as well as empty lines are ignored.  Lines do
# have a certain length limit but this is not serious limitation as
# the format of the entries is fixed and checked by gpg-agent. A
# non-comment line starts with optional white spaces, followed by the
# keygrip of the key given as 40 hex digits, optionally followed by a
# caching TTL in seconds, and another optional field for arbitrary
# flags.   Prepend the keygrip with an '!' mark to disable it.

    __Important Example__ Keys are stored in a stupid, non readable format that cant even
    be used as backup. wtf. thats so dumb. make the keys you are saving to the disk in
    the right foramt. or at least encrypt them if you want to hide them, dont do stupid
    obstification

    *another one* is the random seed, with this method, we make it impossible to regeneate
    keys and make everythignd dficcult while not being super secure. do this much beter.

    *ano-ther one* is the trusted db, its not even a normal db format? why? also why cant I 
    access the data via api? this is a great example of BAD software and where multiverse os
    osftware should be consistent, and learning new software should be dead simple if you
    know any software 

    __Important Example__ Revocation is SUPER important to security, and complex key systems
    and they are just simply WAY underused, beacuse they are not well incorproated into the
    system or UI, or DIR structure

		This is a revocation certificate for the OpenPGP key:

			pub   rsa4096 2018-08-30 [S]
			      1BA0349812C4C907B4BE0E51188B1354CC6219D3
			uid          multiverse <contact@multiverse-os.org>

		A revocation certificate is a kind of "kill switch" to publicly
		declare that a key shall not anymore be used.  It is not possible
		to retract such a revocation certificate once it has been published.

		Use it to revoke this key in case of a compromise or loss of
		the secret key.  However, if the secret key is still accessible,
		it is better to generate a new revocation certificate and give
		a reason for the revocation.  For details see the description of
		of the gpg command "--generate-revocation" in the GnuPG manual.




## Delete files



In `~/.cache/*` excess files

  * Only keep following folders: ['dconf', 'libvirt'] and the files: ['user-dirs.dirs', 'user-dirs.locale']


In `~/.cache/*` excess files

  * Only keep following folders: ['gnome-software', 'libvirt', 'virt-manager']
  * In `gnome-software`, there is a `gnome.conf` file that contains json of
    all extension data. If we were to customize gnome from the ground up, 
    we would get rid of all curreent extensions since they are a massive security
    hazard

  * **Very interesting virt-manager folder contains interesting files**
    specifically the virt-manager log. 

> [Fri, 31 Aug 2018 05:24:08 virt-manager 1964] DEBUG (connection:755) domain lifecycle event: domain=galaxy.router.multiverse event=4 reason=0
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:1052) domain=universe.voyager-sagan.multiverse status=Shutoff added
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:1052) domain=universe.router.multiverse status=Running added
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:790) Using domain events
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:827) Using network events
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:1052) interface=lo status=Active added
[Fri, 31 Aug 2018 05:22:25 virt-manager 1964] DEBUG (connection:1052) pool=os-images status=Active added
__[Fri, 31 Aug 2018 04:25:55 virt-manager 1985] DEBUG (virt-manager:183) virtManager import: <module 'virtManager' from '/usr/share/virt-manager/virtManager/__init__.pyc'>__

We should be getting rid of virt-manager ASAP. Then further and get rid of libvirt, its not right for Multiverse OS. WE have much better and bigger plans

    This is very important for building a multiverse-daemon on BOTH the HOST and CONTROLLER.
    since these are the primary VM machines. If we want live updates without short polling
    like an asshole, we need to find out where the log is generated from or just as a short
    term solution, we essentially tail the log and parse it. Make the file a memory file, NOT 
    actual file, because IO is our most limited resource in modern computing, so lets stop 
    dicking around and wasting it on things that could be mmeory.


IN `~/.local/*` excess files

  only keep: ['libvirt']


## Delete unncessary software

apt-get remove gnome-nibbles
apt-get remove quadrapassel
apt-get remove iagno 
apt-get remove inkscape
apt-get remove firefox-esr
apt-get remove gnome-tetravex
apt-get remove hitori
apt-get remove gnome-klotski 
apt-get remove libreoffice*
apt-get remove lightsoff
apt-get remove gnome-mahjongg
apt-get remove gnome-maps
apt-get remove gnome-mines
apt-get remove four-in-a-row
apt-get remove swell-foop
apt-get remove gnome-robots
apt-get remove gnome-sudoku
apt-get remove tali
apt-get remove gnome-taquin
apt-get remove gnome-weather
apt-get remove gnome-user-share
apt-get remove gnome-sound-recorder
apt-get remove gnome-bluetooth
apt-get remove gnome-online-miners
apt-get remove gnome-contacts
apt-get remove gnome-getting-started-docs
apt-get remove gnome-music
apt-get remove gnome-screenshot
apt-get remove gnome-tweak-tool
apt-get remove gnome-chess
apt-get remove gnome-dictionary
apt-get remove gnome-video-effects
apt-get remove gnome-user-guide
apt-get remove gnome-calendar 
apt-get remove gnome-logs
apt-get remove rhythmbox
apt-get remove transmission-*
apt-get remove reportbug
apt-get remove gnome-online-accounts
apt-get remove rhythmbox-data
apt-get remove gnome-maps
apt-get remove gnome-accessibility-themes
apt-get remove doc-debian
apt-get remove debian-faq 
apt-get remove dleyna-server
apt-get remove aspell
apt-get remove apache2-bin
apt-get remove baobab
apt-get remove brasero
apt-get remove brasero*
apt-get remove bogofilter-bdb
apt-get remove bluez*
apt-get remove cups-pk-helper
apt-get remove colord-data 
apt-get remove cheese-common 
apt-get remove cdrdao
apt-get remove caribou
apt-get remove crda
apt-get remove dvdauthor
apt-get remove evince*
apt-get remove evolution-*
apt-get remove enchant
apt-get remove espeak-ng-data
apt-get remove eog
apt-get remove folks-common
apt-get remove freepats 
apt-get remove hunspell-en-us
apt-get remove hyphen-en-us
apt-get remove hoichess
apt-get remove java-common
apt-get remove five-or-more
apt-get remove polari
apt-get remove totem
apt-get remove gnome-music
apt-get remove virt-viewer
apt-get remove simple-scan
apt-get remove gnome-calendar
apt-get remove gnome-clocks
apt-get remove gnome-software
apt-get remove synaptic
apt-get remove vinagre
apt-get remove evolution
apt-get remove gnome-calculator

then autoremove. This will keep basics of gnome, which is useful until a mutliverse daemon can sit on the host, provide secure API over virtIO, and lockdown the server, using a mix between extreme disabling of features and setting up honeypots for any behavior that should be impossible to provide low level, hardware backdoor instrusion detection. Looking at you intel.

# Not sure if these will break everything
apt-get remove docbook-xml

apt-get remove default-jre 
apt-get remove default-jre-headless


## NEXT TRY THESE [not delted yet]
bind9-host


## wont delete 
apt-get remove ispell
apt-get remove ienglish-common
imagemagick 
apt-get remove mobile-broadband-provider-info
________________________________________________________________________________________________
## Review Proccesses Listening To Ports And Crush Them

`ss -a4n`

````
Netid  State      Recv-Q Send-Q Local Address:Port               Peer Address:Port              
udp    UNCONN     0      0         *:5353                  *:*                  
````

5353 by default is mdns (multicast dns), which is apart of avahi-daemon. It could also be a trojan.

`apt-get remove avahi-daemon`


````
udp    UNCONN     0      0         *:38346                 *:*                  
udp    UNCONN     0      0      10.1.1.254:53                    *:*                  
udp    UNCONN     0      0      10.0.0.254:53                    *:*                  
udp    UNCONN     0      0      10.152.152.254:53                    *:*                  
udp    UNCONN     0      0         *:1900                  *:*                  
tcp    LISTEN     0      5      10.1.1.254:53                    *:*                  
tcp    LISTEN     0      5      10.0.0.254:53                    *:*                  
tcp    LISTEN     0      5      10.152.152.254:53                    *:*                  
tcp    LISTEN     0      128       *:22                    *:*        
````

________________________________________________________________________________________________
## Disable ALL incoming connections, leave outgoing connections open for now
This can be done easily with `iptables` and eventually should be done with a userspace
network stack that just pulls the packets from the NIC directly and throws them away to
prevent any sort of special formed packet that can bypass kernel or soemthing.

````
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT ACCEPT

iptables -A INPUT -i lo -j ACCEPT
iptables -A OUPUT -o lo -j ACCEPT

iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
````

________________________________________________________________________________________________
## Disabling Unnecessary PCI devices


Disabling PCI devices for NIC by pretending them from ever being loaded in the kernel, prevents
any potential attack surface from the point they are loaded when the kernel starts (beginning of
initramfs) to disabling right before login via GDM. 





````
00:1f.3 Audio device [0403]: Intel Corporation Device [8086:a2f0]
	Subsystem: ASUSTeK Computer Inc. Device [1043:8724]
	Kernel modules: snd_hda_intel


00:1f.6 Ethernet controller [0200]: Intel Corporation Ethernet Connection (2) I219-V [8086:15b8]
	Subsystem: ASUSTeK Computer Inc. Ethernet Connection (2) I219-V [1043:8672]
	Kernel driver in use: vfio-pci
	Kernel modules: e1000e

02:00.0 Network controller [0280]: Wilocity Ltd. Wil6200 802.11ad Wireless Network Adapter [1ae9:0310] (rev 02)
	Subsystem: Wilocity Ltd. Wil6200 802.11ad Wireless Network Adapter [1ae9:0000]
	Kernel driver in use: vfio-pci
	Kernel modules: wil6210


04:00.0 Ethernet controller [0200]: Intel Corporation I211 Gigabit Network Connection [8086:1539] (rev 03)
	Subsystem: ASUSTeK Computer Inc. I211 Gigabit Network Connection [1043:85f0]
	Kernel driver in use: vfio-pci
	Kernel modules: igb


05:00.0 Network controller [0280]: Qualcomm Atheros QCA6174 802.11ac Wireless Network Adapter [168c:003e] (rev 32)
	Subsystem: ASUSTeK Computer Inc. QCA6174 802.11ac Wireless Network Adapter [1043:8751]
	Kernel driver in use: vfio-pci
	Kernel modules: ath10k_pci

19:00.1 Audio device [0403]: Advanced Micro Devices, Inc. [AMD/ATI] Device [1002:aaf8]
	Subsystem: Advanced Micro Devices, Inc. [AMD/ATI] Device [1002:aaf8]
	Kernel driver in use: vfio-pci
	Kernel modules: snd_hda_intel

````


### Remove Intel's hardware backdoor that always has some working exploit at all times
In fact its so broken disabling the mei_me does almost nothing to stop it from being
exploited.

Intel is terrible.


00:16.0 Communication controller [0780]: Intel Corporation Device [8086:a2ba]
	Subsystem: ASUSTeK Computer Inc. Device [1043:873c]
	Kernel driver in use: mei_me
	Kernel modules: mei_me


### Blacklisting kernel modules
Modify `/etc/modprobe.d/multiverse.conf` and add the following to the top of the file:


````
blacklist snd_hda_intel
blacklist amdgpu
blacklist mei_me
blacklist ath10k_pci
blacklist igb
blacklist wil6210
blacklist e1000e
````

