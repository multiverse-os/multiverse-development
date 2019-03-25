
## General Notes / Ideas [Generated while working on below, merge into different file]

  * DIRECT MEMORY ACCESS, the best way to move video around, fuck xephyr, fuck xpra,
    fuck vnc, fuck loking glass (that was my name!)

---


  [!][A better system for knowing, saving and using global environmental variables using]
     [a more complex scripting language, that can do closures for variables]




  [!][Stream lining application and service VMs using templating, scripts, one-shot VMs, etc]
     **Tree based VM templating**
   
   For example, Firefox VMs are going to be commong, as common as Tor VMs. And we don't
   want to reinstall our addons each time ,but we also want to ability to segregate
   bookmarks.

   WE need a way to define what parts should be apart of all VMs of a given app VM type
   ie Firefox app VM, should all have ublock origin. 

   WE do this using our TREE BASED VM templating system. So once a VM is setup, 
   one can branch, 2 or 3 or x bifuctations. all using the same base system. 

   ALSO BECAUSE WE ARE AWESOME, changes to the base system automatically update
   deriviatve systems, making updating simpeler!


  [!][Multiverse Daemon - the looking glass that allows you to see linux in new ways]


  [!][New Networking Structure Idea]
  * Like how the Controller VM nests both Application VM and Service VM. And
    multiple Controller can be ran to segregate service VMs. 

    It may be interesting to instead of having the Router VMs be loose on the HOST
    that they are put inside a simplified Controller VM. This would simplify complex
    networking paths, that can be spun up together or re wired. This also allows logical
    "softwawre defined networking" style switches and routers to be put inbetween a 
    networking path. Like for example, a simple script to sort packlets between two
    routers based on outgoing IP location without adding a third router, just a simple
    script that sorts packets.

    This also puts an encryption layer between the host and the Router VMs. Which would
    portect the **RAM**, the **DISKs**, and any other side channels. 

      * Another way to improve this new design would be to REMOVE netowrking from the
        controller, so it is never accidently used to access the itnernet directly
        for example, iut should never use firefox because it could have JS
        
        Firefox is loaded in an applicaiton VM. __Security could be improved MASSIVEDLY**
        if networking of packets from the application VM went DIRECTLY to the ROUTER VM
        and not using the kernel packet networking of any of the linux systems inbetween
        instead, it should just be direclty passed using userpsace networking stack.
      **This concept could be tested by disabling internet on the Controller VM, by disabling the**
      ability for these devicvse to access the internet thesmelvse probalby via iptables. but
      eventually can be improved by requiring connections initiation to require a signed packet
      that awya the router can just rjefct any stream initiationzation from non-app non-service VM types




   [!][Prevent any potential side channeling or manipualtion that could happen from HOST, via shared /dev/*]
    **Consider the idea that perhaps the /dev/* that get attached from the HOST to the**
    VMs creates side channel opportunities, and it allows for manipualtion of VMs by
    hacking the devices. We want to find a way to prvent this from happening.







   [!][Cocnept of user data, settings, clipboard, should be more distributed]

    * Files should also be able to store code, or queries. 
    * Same with folders


__________________________________________________
# Multiverse OS Controller Setup
The controller is the interface for the entire Multiverse OS cluster. The controller VM
currently is ran one-at-a-time, because it is the interface for the user. It is the 
desktop that the monitor is plugged into. It is the computer that the keyboard and
other input is plugged into.

It may seem redundant at first, but this design is not accidently, its carefully thought
out and its nested and repeated nature is a big part of the underlying securty. Multiverse
OS always relies on a layered security approach, that assumes failure of even certain
things, because certainty is an illusion. And this is we keep passwords unlike Qemu
we just come up= with a different way of simplyifying the UI for the user, but we 
never ever give p security for speed or convience, because first and foremost
Multiverse is a secure OS, it is THE secure OS. 

**The Controller VM like the Host DOES NOT ACCESS THE ITNERNET DIRECTLY EVER!** 

Instead if you wantt o use firefox you use a firefox VM. you need torrents? a torrenting service VM.


`history -a; tail -5 $HISTFILE`

## General setup, and restructuring of the Multiverse OS Controller VM 
We will be using symbolic links from the IDE hard disks to store super large files until we have a clustering 
FS that treats them all as a single drive and does duplication acrross drives (RAID - some coool number). 




_Get rid of excess in `~/.cache` and prepare the rest for easy parsing and merging into the Multiverse_
daemon that manages all user settings, logs, scripts, etc for the Controller VM. 

We DONT want this stuff just floating around, and stored in random, inconsistent formats. Instead
we will store it in a easy to understand DB format, that can print out tables, automatically
provide data via API for easy programming. and can be output in a variety of data formats, 
including the original legacy format. 

**Great example of what kind of data we are parsing and merging is `~/.cache/gnome-software/3.22/extensions/gnome.json`
whioch contantains all the information needed to build a UI to explain and download gnome JS/CSS based extensions
to GNome. These are not checked for malware and should not be trusted, so if we are modifying gnome, we need to
go ATLEAST this deep. to start over on extensions.

> __"3.2.3": {"pk": 289, "version": 5}, "3.7.5": {"pk": 2495, "version": 18}}, "description": "Add a menu for quickly navigating places in the system.\nThis extension is part of Classic Mode and is officially supported by GNOME. Please do not report bugs using the form below, use GNOME's GitLab instance instead.", "screenshot": "/extension-data/screenshots/screenshot_8_mVLeGic.png", "creator": "fmuellner", "name": "Places Status Indicator", "link": "/extension/8/places-status-indicator/", "pk": 8, "creator_url": "/accounts/profile/fmuellner", "icon": "/extension-data/icons/icon_8.png", "uuid": "places-menu@gnome-shell-extensions.gcampax.github.com"}, {"shell_version_map": {"3.3.90": {"pk": 853, "version": 8}, "3.7.90": {"pk": 2568, "version": 17}, "3.8": {"pk": 3222, "version": 21}, "3.7.92": {"pk": 2648, "version": 18}, "3.10": {"pk": 3253, "version": 22}, "3.5.5": {"pk": 1696, "version": 10}, "3.14": {"pk": 4387, "version": 25}, "3.4": {"pk": 1014, "version": 9}, "3.6": {"pk": 2217, "version": 14}, "3.15.1": {"pk": 4452, "version": 26},__

Additonally isntead of 20 files, the data will be merged into a single graph database, with different
maps of relationships. 

**WE ARE REMOVING ALL EXCESS beacuse these config files types will be on the APP VMs. Removing the excess**
 _leaves us with the basis, we will use to parse, merge, and then can have an UI for easy settings management_
 _that will be used to regenerate these files and overwrite the versions on the applicaiton VM._



========
# SETUP/CONFIGURATION/ETC

All config files are merged into a database, covnerted to a consistent format. Then an voerlay can be generated
and dumped into a system so we are legacy compatible.

=================================================================================================================
Clear top level folders, we will use symbolic links

````
rm -rf Music/
rm -rf Pictures/
rm -rf Videos/
rm -rf Downloads/
rm -rf Notes/



# scripts 
  ~/.bash_logout
  ~/.bashrc
  ~/.profile
  

````


[merge together][from ~/]

**LOOSE ~./.bash_history** is simply not acceptable. its bad security, its not as useful as an actual DB, with search. 
parsing is slow as fuck. thi0s is jsut old and we need to do better.


_.xbindkeysrc_ was where key bindiers are stored



__some of these dont hold config, they hold scripts lol, lets try to be consistent. our configs will be scripts that
pull from our DB!__





______________________
## Cleaning HOME folder, and preparing data to be parsed into Multiverse OS Controller daemon to replace old
legacy files and structure.

#### .CACHE
__Like if this cache is literally being read IO from and to the disk, this could be improved massively. look into that__
these need to be parsed and merged. 


````
rm -rf ~/.cache/gnome-calculator
rm -rf ~/.cache/evolution
rm -rf ~/.cache/folks
rm -rf ~/.cache/libgweather
````
We are left with in `~/.cache`: [`fontconfig`, `gnome-software`, `gstreamer-1.0`, `libvirt`, `thumbnails`, `tracker`, `virt-manager`, `mesa_shader_cache`]

	> [?][not sure, at the very least, should be parsed and merged into a common consistent db]
	> rm -rf ~/.cache/fontconfig



#### .CONFIG


**Libvirt files need to be merged in here. Should automatically setup files in `/var/multiverse`* structure. For now just using XML files dropped in. Later generate XML or whatever else config needed from DB.** _Look into what secrets is used for, and how we can tie that into multierse_ __And another improtant concept here is `qemu/channel/target`, which is probably how we interact wtih VirtIO.



[QEMU VM][Key concepts in `~/.config/libvirt/qemu/*`]: parse all this data, merge it into a consistent DB, capasble of outputting XML files, etc
for legacy support. 

  * **channel/target** - likely how to itneract with virtIO, which is what we will do our netstack with, or DMA

  * **dump** 

  * **lib**
  
  * **nvram** - non-violitaile memory, basically our CMOS, and other things. WE NEED HASHES OF THESE, its fucking dumb, libvirt doesnt
    do this already. [same with HDs for that matter]

  * **ram** encrypt with key for system. **Like short term or wroking memory in humans~! or actiev memory**

  * **save** 
 
  * **snapshot** will need our own snapshot system, we also wnat epeheraml systems remember that? it was a critical part of the design lol

  * **secrets** keys
  
  * **storage** xml describing storage polls, that shit is annoyuing really

````
rm -rf ~/.config/enchant
rm -rf ~/.config/evolution
rm -rf ~/.config/gnome-session
rm -rf ~/.config/goa-1.0
rm -rf ~/.config/ibus
rm -rf ~/.config/procps



````


[!][__Important ~/.config files__]
**These are a great example of what Multiverse OS wants to address to make linux easier to undrstand, easier to learn, and jsut fucking better, faster and more prodcutive** _currently kinda sucks how its organized, how its different data formats, how its hard and slow for computer to parse. THIS CAN BE BETTER!_


````[mimeapps.list] **this should be easier to see, easeir to edit, and probalby better than just a flat file. man this probalby
is slow as fuck comapred to a binary format that we just print for human readable format
````

````[Added Associations]
application/octet-stream=org.gnome.gedit.desktop;
````


````['~/.config/gtk-3.0'] has file bookmarks probablyh for nautilus
````

````[.monitors] xml with monitor info
````

````[~/.config/nautilius/search-metadata]
[directory]
nautilus-list-view-sort-column=search_relevance
nautilus-list-view-sort-reversed=true
````

````/.config/pulse [NOT SURE IF THIS IS RIGHT, dont know how best to handle clustering the sound, definltey something that can layer it all together- whcih is probalby pulse or pulse+jack] __I think it may be better to use something else, this is kinda not exaclty right__
````

````[~/.config/user-dirs.dirs]
XDG_DESKTOP_DIR="$HOME/Desktop"
XDG_DOWNLOAD_DIR="$HOME/Downloads"
XDG_TEMPLATES_DIR="$HOME/"
XDG_PUBLICSHARE_DIR="$HOME/"
XDG_DOCUMENTS_DIR="$HOME/Documents"
XDG_MUSIC_DIR="$HOME/Music"
XDG_PICTURES_DIR="$HOME/Pictures"
XDG_VIDEOS_DIR="$HOME/Videos"
````

````[~/.config/user-dirs.locale]
en_Us
````


_________________________________________________________________________________



#### .LOCAL/SHARE










Dont bother with either keyhrings or libvirt. libvirt just stores images there and they should be `/var/multiverse`

````

rm -rf ~/.local/share/applications
rm -rf ~/.local/share/libvirt 
rm -rf ~/.local/share/folks
rm -rf ~/.local/share/nautilus
rm -rf ~/.local/share/sounds
rm -rf ~/.local/share/telepathy
rm -rf ~/.local/share/xorg
rm -rf ~/.local/share/tracker

````

just delete above, not needed and any new files would be generated on APP Vm

````
rm -rf ~/.local/share/keyrings
````

[! key rings suck][mega bad design][!][home/user/.local/share/keyrings/login.keyring]
__the keyring for login, holds the LUKs passwords, probably jsut symmetric, this is terrible. these should be session
based. encrypting the data to the session key. **NEEDS A LOT OF WORK, MERGE in WITH SCRAMBLE SUIT KEY SYSETM**
  [*] just delete the key rings and rebuild it. its really bad.


[! xorg logs shouldnt be writing to disk!][they should be writing to memory, merging with all other logs in the system into a single DB. then writing to IO way alter on an unused slow disk with low priority. this group logs will generate toast type notifaction if log matches a regex/scripted matcher.w


**IMportant files to merge**


[~/.local/share/recently-used.xbel][this def needs to move to DB ssytem. add a bunch of new features after it gets in there, avoid IO]

[icc][ should find out where its getting written from and just insert into settings DB and prevent IO.] __this current `~/home` settings and files system is terrible. it waste our most valuable resource IO, is inconsistent, is designed for humans to read, so many reasosn.___



[`~/.local/share/app-info/xmls/extensions-web.xml]
  _provides an ACTIVE list of current elements in the shell in XML. so like clock at top center._
  _activites in top left, power in top right, etc_

````[`~/.local/share/gnome-shell/applicationstate][info about windows]
<?xml version="1.0"?>
<application-state>
  <context id="">
    <application id="gnome-control-center.desktop" open-window-count="0" score="4" last-seen="1535712467"/>
    <application id="gnome-display-panel.desktop" open-window-count="0" score="0" last-seen="0"/>
    <application id="org.gnome.Nautilus.desktop" open-window-count="3" score="72" last-seen="1535716979"/>
    <application id="org.gnome.gedit.desktop" open-window-count="2" score="550" last-seen="1535716968"/>
    <application id="gnome-network-panel.desktop" open-window-count="0" score="0" last-seen="0"/>
    <application id="virt-manager.desktop" open-window-count="1" score="70" last-seen="1535712082"/>
    <application id="org.gnome.Terminal.desktop" open-window-count="3" score="262" last-seen="1535716819"/>
    <application id="org.gnome.DiskUtility.desktop" open-window-count="0" score="3" last-seen="1535667376"/>
  </context>
</application-state>
````
[`~/.local/share/gnome-shell/notifications] if this how they are saved and parsed that is terrible. very very bad.
just use logging, and have a system to indicate if a log should be used as a notification. so much redundant writing
this way. 


````[~/.local/share/gsettings-data-convert][right click covnert options, not sure how they are hooking the button to the script that would be called. need a way to store scripts in a unviersal way, that can be asscoatied vay graphDB. shared, schjeduled, etc. FOR EXAMPLE, you shouldnt have to write a script in cron, you should have a script written and tested, then you just hook in the UUID to crons type scheduling.]
[State]
timestamp=1535666861
converted=gedit.convert;org.gnome.seahorse.manager.convert;mousetweaks.convert;org.gnome.crypto.pgp.convert;evolution-data-server.convert;wm-schemas.convert;file-roller.convert;org.gnome.crypto.pgp_keyservers.convert;gsettings-desktop-schemas.convert;opensubtitles.convert;brasero.convert;evolution.convert;libgnomekbd.convert;org.gnome.Vinagre.convert;gnome-settings-daemon.convert;gnome-shell-overrides.convert;totem.convert;gvfs-dns-sd.convert;org.gnome.crypto.cache.convert;org.gnome.seahorse.recipients.convert;gnome-user-share.convert;folks.convert;gvfs-smb.convert;org.gnome.seahorse.convert;eog.convert;org.virt-manager.virt-manager.convert;evince.convert;pythonconsole.convert;gnome-session.convert;nm-applet.convert;gnome-screenshot.convert;mutter-schemas.convert;
````

======================================================================================
## DRM - Direct Read Memory, in linux started for interacting with GPU, and 3D accel
------------------------------------------------------------------------------------

DRM is handled for you by the card.


**KMS** is a recent update to DRM to fix it. Moves things inside kernel space, so
display modes can be moified from the kernel.



======================================================================================
````



## Qemu Capabilities
Ever wonder how we will determine the best way to build QEMU XML or even just general qemu_system lines when we 
get rid of `virt-manager` then `libvirt`? Well we can first strat by grabbing the generated qemu caps
gerneated. learn how they generate those. and make our own for use in all future QEMU vm geenration on that hardawre.

**This is standard, it jsut says what each CPU type has which CPU features. WE can base our list on this, then**
modify it based on security reasons. Like remove bad features that are insecure, and adding them back just
isnt an option. ANd we can create new ones based on our security choices as new hardawre comes out


<qemuCaps>
  <qemuctime>1535700104</qemuctime>
  <selfctime>1535699948</selfctime>
  <selfvers>4006000</selfvers>
  <usedQMP/>
  <flag name='kvm'/>
  <flag name='no-hpet'/>
  <flag name='spice'/>
  <flag name='boot-index'/>
  <flag name='hda-duplex'/>
  <flag name='ccid-emulated'/>
  <flag name='ccid-passthru'/>
  <flag name='virtio-tx-alg'/>
  <flag name='virtio-blk-pci.ioeventfd'/>
  <flag name='sga'/>
  <flag name='virtio-blk-pci.event_idx'/>
  <flag name='virtio-net-pci.event_idx'/>
  <flag name='piix3-usb-uhci'/>
  <flag name='piix4-usb-uhci'/>
  <flag name='usb-ehci'/>
  <flag name='ich9-usb-ehci1'/>
  <flag name='vt82c686b-usb-uhci'/>
  <flag name='pci-ohci'/>
  <flag name='usb-redir'/>
  <flag name='usb-hub'/>
  <flag name='ich9-ahci'/>
  <flag name='no-acpi'/>
  <flag name='virtio-blk-pci.scsi'/>
  <flag name='scsi-disk.channel'/>
  <flag name='scsi-block'/>
  <flag name='transaction'/>
  <flag name='block-job-async'/>
  <flag name='scsi-cd'/>
  <flag name='ide-cd'/>
  <flag name='hda-micro'/>
  <flag name='dump-guest-memory'/>
  <flag name='nec-usb-xhci'/>
  <flag name='balloon-event'/>
  <flag name='lsi'/>
  <flag name='virtio-scsi-pci'/>
  <flag name='blockio'/>
  <flag name='disable-s3'/>
  <flag name='disable-s4'/>
  <flag name='usb-redir.filter'/>
  <flag name='ide-drive.wwn'/>
  <flag name='scsi-disk.wwn'/>
  <flag name='seccomp-sandbox'/>
  <flag name='reboot-timeout'/>
  <flag name='seamless-migration'/>
  <flag name='block-commit'/>
  <flag name='vnc'/>
  <flag name='drive-mirror'/>
  <flag name='usb-redir.bootindex'/>
  <flag name='usb-host.bootindex'/>
  <flag name='blockdev-snapshot-sync'/>
  <flag name='qxl'/>
  <flag name='VGA'/>
  <flag name='cirrus-vga'/>
  <flag name='vmware-svga'/>
  <flag name='device-video-primary'/>
  <flag name='usb-serial'/>
  <flag name='usb-net'/>
  <flag name='add-fd'/>
  <flag name='nbd-server'/>
  <flag name='virtio-rng'/>
  <flag name='rng-random'/>
  <flag name='rng-egd'/>
  <flag name='megasas'/>
  <flag name='tpm-passthrough'/>
  <flag name='tpm-tis'/>
  <flag name='pci-bridge'/>
  <flag name='vfio-pci'/>
  <flag name='vfio-pci.bootindex'/>
  <flag name='scsi-generic'/>
  <flag name='scsi-generic.bootindex'/>
  <flag name='mem-merge'/>
  <flag name='vnc-websocket'/>
  <flag name='drive-discard'/>
  <flag name='mlock'/>
  <flag name='device-del-event'/>
  <flag name='dmi-to-pci-bridge'/>
  <flag name='i440fx-pci-hole64-size'/>
  <flag name='q35-pci-hole64-size'/>
  <flag name='usb-storage'/>
  <flag name='usb-storage.removable'/>
  <flag name='virtio-mmio'/>
  <flag name='ich9-intel-hda'/>
  <flag name='kvm-pit-lost-tick-policy'/>
  <flag name='boot-strict'/>
  <flag name='pvpanic'/>
  <flag name='spice-file-xfer-disable'/>
  <flag name='spiceport'/>
  <flag name='usb-kbd'/>
  <flag name='msg-timestamp'/>
  <flag name='active-commit'/>
  <flag name='change-backing-file'/>
  <flag name='memory-backend-ram'/>
  <flag name='numa'/>
  <flag name='memory-backend-file'/>
  <flag name='usb-audio'/>
  <flag name='rtc-reset-reinjection'/>
  <flag name='splash-timeout'/>
  <flag name='iothread'/>
  <flag name='migrate-rdma'/>
  <flag name='ivshmem'/>
  <flag name='drive-iotune-max'/>
  <flag name='VGA.vgamem_mb'/>
  <flag name='vmware-svga.vgamem_mb'/>
  <flag name='qxl.vgamem_mb'/>
  <flag name='pc-dimm'/>
  <flag name='machine-vmport-opt'/>
  <flag name='aes-key-wrap'/>
  <flag name='dea-key-wrap'/>
  <flag name='pci-serial'/>
  <flag name='vhost-user-multiqueue'/>
  <flag name='migration-event'/>
  <flag name='ioh3420'/>
  <flag name='x3130-upstream'/>
  <flag name='xio3130-downstream'/>
  <flag name='rtl8139'/>
  <flag name='e1000'/>
  <flag name='virtio-net'/>
  <flag name='gic-version'/>
  <flag name='incoming-defer'/>
  <flag name='virtio-gpu'/>
  <flag name='virtio-keyboard'/>
  <flag name='virtio-mouse'/>
  <flag name='virtio-tablet'/>
  <flag name='virtio-input-host'/>
  <flag name='chardev-file-append'/>
  <flag name='ich9-disable-s3'/>
  <flag name='ich9-disable-s4'/>
  <flag name='vserport-change-event'/>
  <flag name='virtio-balloon-pci.deflate-on-oom'/>
  <flag name='mptsas1068'/>
  <flag name='qxl.vram64_size_mb'/>
  <flag name='chardev-logfile'/>
  <flag name='debug-threads'/>
  <flag name='secret'/>
  <flag name='pxb'/>
  <flag name='pxb-pcie'/>
  <flag name='device-tray-moved-event'/>
  <flag name='nec-usb-xhci-ports'/>
  <flag name='virtio-scsi-pci.iothread'/>
  <flag name='name-guest'/>
  <flag name='qxl.max_outputs'/>
  <flag name='spice-unix'/>
  <flag name='drive-detect-zeroes'/>
  <flag name='tls-creds-x509'/>
  <flag name='intel-iommu'/>
  <flag name='smm'/>
  <flag name='virtio-pci-disable-legacy'/>
  <flag name='query-hotpluggable-cpus'/>
  <flag name='virtio-net.rx_queue_size'/>
  <flag name='virtio-vga'/>
  <flag name='drive-iotune-max-length'/>
  <flag name='ivshmem-plain'/>
  <flag name='ivshmem-doorbell'/>
  <flag name='query-qmp-schema'/>
  <flag name='gluster.debug_level'/>
  <flag name='vhost-scsi'/>
  <flag name='drive-iotune-group'/>
  <flag name='query-cpu-model-expansion'/>
  <flag name='virtio-net.host_mtu'/>
  <flag name='nvdimm'/>
  <flag name='pcie-root-port'/>
  <flag name='query-cpu-definitions'/>
  <flag name='block-write-threshold'/>
  <flag name='query-named-block-nodes'/>
  <flag name='cpu-cache'/>
  <flag name='qemu-xhci'/>
  <flag name='kernel-irqchip'/>
  <flag name='kernel-irqchip.split'/>
  <flag name='intel-iommu.intremap'/>
  <flag name='intel-iommu.caching-mode'/>
  <flag name='intel-iommu.eim'/>
  <flag name='intel-iommu.device-iotlb'/>
  <flag name='virtio.iommu_platform'/>
  <flag name='virtio.ats'/>
  <flag name='loadparm'/>
  <flag name='vnc-multi-servers'/>
  <flag name='virtio-net.tx_queue_size'/>
  <flag name='chardev-reconnect'/>
  <flag name='virtio-gpu.max_outputs'/>
  <flag name='vxhs'/>
  <flag name='virtio-blk.num-queues'/>
  <flag name='vmcoreinfo'/>
  <flag name='numa.dist'/>
  <flag name='disk-share-rw'/>
  <flag name='iscsi.password-secret'/>
  <flag name='isa-serial'/>
  <flag name='dump-completed'/>
  <flag name='qcow2-luks'/>
  <flag name='pcie-pci-bridge'/>
  <flag name='seccomp-blacklist'/>
  <flag name='query-cpus-fast'/>
  <flag name='disk-write-cache'/>
  <flag name='nbd-tls'/>
  <flag name='tpm-crb'/>
  <flag name='pr-manager-helper'/>
  <flag name='qom-list-properties'/>
  <flag name='memory-backend-file.discard-data'/>
  <flag name='sdl-gl'/>
  <flag name='screendump_device'/>
  <flag name='hda-output'/>
  <flag name='blockdev-del'/>
  <flag name='vmgenid'/>
  <flag name='vhost-vsock'/>
  <flag name='chardev-fd-pass'/>
  <flag name='tpm-emulator'/>
  <flag name='mch'/>
  <flag name='mch.extended-tseg-mbytes'/>
  <flag name='sev-guest'/>
  <flag name='egl-headless'/>
  <flag name='vfio-pci.display'/>
  <version>2012000</version>
  <kvmVersion>0</kvmVersion>
  <microcodeVersion>1</microcodeVersion>
  <package>Debian 1:2.12+dfsg-3</package>
  <kernelVersion>4.17.0-3-amd64 #1 SMP Debian 4.17.17-1 (2018-08-18)</kernelVersion>
  <arch>x86_64</arch>
  <hostCPU type='kvm' model='base' migratability='yes'>
    <property name='phys-bits' type='number' value='0'/>
    <property name='core-id' type='number' value='-1'/>
    <property name='xlevel' type='number' value='2147483656'/>
    <property name='cmov' type='boolean' value='true' migratable='yes'/>
    <property name='ia64' type='boolean' value='false'/>
    <property name='aes' type='boolean' value='true' migratable='yes'/>
    <property name='mmx' type='boolean' value='true' migratable='yes'/>
    <property name='rdpid' type='boolean' value='false'/>
    <property name='arat' type='boolean' value='true' migratable='yes'/>
    <property name='gfni' type='boolean' value='false'/>
    <property name='pause-filter' type='boolean' value='false'/>
    <property name='xsavec' type='boolean' value='true' migratable='yes'/>
    <property name='intel-pt' type='boolean' value='false'/>
    <property name='osxsave' type='boolean' value='false'/>
    <property name='hv-frequencies' type='boolean' value='false'/>
    <property name='tsc-frequency' type='number' value='0'/>
    <property name='xd' type='boolean' value='true' migratable='yes'/>
    <property name='hv-vendor-id' type='string' value=''/>
    <property name='kvm-asyncpf' type='boolean' value='true' migratable='yes'/>
    <property name='kvm_asyncpf' type='boolean' value='true' migratable='yes'/>
    <property name='perfctr_core' type='boolean' value='false'/>
    <property name='perfctr-core' type='boolean' value='false'/>
    <property name='mpx' type='boolean' value='true' migratable='yes'/>
    <property name='pbe' type='boolean' value='false'/>
    <property name='decodeassists' type='boolean' value='false'/>
    <property name='avx512cd' type='boolean' value='true' migratable='yes'/>
    <property name='sse4_1' type='boolean' value='true' migratable='yes'/>
    <property name='sse4.1' type='boolean' value='true' migratable='yes'/>
    <property name='sse4-1' type='boolean' value='true' migratable='yes'/>
    <property name='family' type='number' value='6'/>
    <property name='vmware-cpuid-freq' type='boolean' value='true' migratable='yes'/>
    <property name='avx512f' type='boolean' value='true' migratable='yes'/>
    <property name='msr' type='boolean' value='true' migratable='yes'/>
    <property name='mce' type='boolean' value='true' migratable='yes'/>
    <property name='mca' type='boolean' value='true' migratable='yes'/>
    <property name='hv-runtime' type='boolean' value='false'/>
    <property name='xcrypt' type='boolean' value='false'/>
    <property name='thread-id' type='number' value='-1'/>
    <property name='min-level' type='number' value='13'/>
    <property name='xgetbv1' type='boolean' value='true' migratable='yes'/>
    <property name='cid' type='boolean' value='false'/>
    <property name='hv-relaxed' type='boolean' value='false'/>
    <property name='hv-crash' type='boolean' value='false'/>
    <property name='ds' type='boolean' value='false'/>
    <property name='fxsr' type='boolean' value='true' migratable='yes'/>
    <property name='xsaveopt' type='boolean' value='true' migratable='yes'/>
    <property name='xtpr' type='boolean' value='false'/>
    <property name='avx512vl' type='boolean' value='true' migratable='yes'/>
    <property name='avx512-vpopcntdq' type='boolean' value='false'/>
    <property name='phe' type='boolean' value='false'/>
    <property name='extapic' type='boolean' value='false'/>
    <property name='3dnowprefetch' type='boolean' value='true' migratable='yes'/>
    <property name='avx512vbmi2' type='boolean' value='false'/>
    <property name='cr8legacy' type='boolean' value='false'/>
    <property name='cpuid-0xb' type='boolean' value='true' migratable='yes'/>
    <property name='xcrypt-en' type='boolean' value='false'/>
    <property name='kvm_pv_eoi' type='boolean' value='true' migratable='yes'/>
    <property name='apic-id' type='number' value='4294967295'/>
    <property name='pn' type='boolean' value='false'/>
    <property name='dca' type='boolean' value='false'/>
    <property name='vendor' type='string' value='GenuineIntel'/>
    <property name='pku' type='boolean' value='false'/>
    <property name='smx' type='boolean' value='false'/>
    <property name='cmp_legacy' type='boolean' value='false'/>
    <property name='cmp-legacy' type='boolean' value='false'/>
    <property name='node-id' type='number' value='-1'/>
    <property name='avx512-4fmaps' type='boolean' value='false'/>
    <property name='vmcb_clean' type='boolean' value='false'/>
    <property name='vmcb-clean' type='boolean' value='false'/>
    <property name='3dnowext' type='boolean' value='false'/>
    <property name='hle' type='boolean' value='true' migratable='yes'/>
    <property name='npt' type='boolean' value='false'/>
    <property name='memory' type='string' value='/machine/unattached/system[0]'/>
    <property name='clwb' type='boolean' value='true' migratable='yes'/>
    <property name='lbrv' type='boolean' value='false'/>
    <property name='adx' type='boolean' value='true' migratable='yes'/>
    <property name='ss' type='boolean' value='true' migratable='yes'/>
    <property name='pni' type='boolean' value='true' migratable='yes'/>
    <property name='svm_lock' type='boolean' value='false'/>
    <property name='svm-lock' type='boolean' value='false'/>
    <property name='pfthreshold' type='boolean' value='false'/>
    <property name='smep' type='boolean' value='true' migratable='yes'/>
    <property name='smap' type='boolean' value='true' migratable='yes'/>
    <property name='x2apic' type='boolean' value='true' migratable='yes'/>
    <property name='avx512vbmi' type='boolean' value='false'/>
    <property name='avx512vnni' type='boolean' value='false'/>
    <property name='hv-stimer' type='boolean' value='false'/>
    <property name='i64' type='boolean' value='true' migratable='yes'/>
    <property name='flushbyasid' type='boolean' value='false'/>
    <property name='f16c' type='boolean' value='true' migratable='yes'/>
    <property name='ace2-en' type='boolean' value='false'/>
    <property name='pat' type='boolean' value='true' migratable='yes'/>
    <property name='pae' type='boolean' value='true' migratable='yes'/>
    <property name='sse' type='boolean' value='true' migratable='yes'/>
    <property name='phe-en' type='boolean' value='false'/>
    <property name='kvm_nopiodelay' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-nopiodelay' type='boolean' value='true' migratable='yes'/>
    <property name='tm' type='boolean' value='false'/>
    <property name='kvmclock-stable-bit' type='boolean' value='true' migratable='yes'/>
    <property name='hypervisor' type='boolean' value='true' migratable='yes'/>
    <property name='socket-id' type='number' value='-1'/>
    <property name='pcommit' type='boolean' value='false'/>
    <property name='syscall' type='boolean' value='true' migratable='yes'/>
    <property name='level' type='number' value='13'/>
    <property name='avx512dq' type='boolean' value='true' migratable='yes'/>
    <property name='svm' type='boolean' value='false'/>
    <property name='full-cpuid-auto-level' type='boolean' value='true' migratable='yes'/>
    <property name='hv-reset' type='boolean' value='false'/>
    <property name='invtsc' type='boolean' value='false'/>
    <property name='sse3' type='boolean' value='true' migratable='yes'/>
    <property name='sse2' type='boolean' value='true' migratable='yes'/>
    <property name='est' type='boolean' value='false'/>
    <property name='avx512ifma' type='boolean' value='false'/>
    <property name='tm2' type='boolean' value='false'/>
    <property name='kvm-pv-eoi' type='boolean' value='true' migratable='yes'/>
    <property name='cx8' type='boolean' value='true' migratable='yes'/>
    <property name='kvm_mmu' type='boolean' value='false'/>
    <property name='kvm-mmu' type='boolean' value='false'/>
    <property name='sse4_2' type='boolean' value='true' migratable='yes'/>
    <property name='sse4.2' type='boolean' value='true' migratable='yes'/>
    <property name='sse4-2' type='boolean' value='true' migratable='yes'/>
    <property name='pge' type='boolean' value='true' migratable='yes'/>
    <property name='fill-mtrr-mask' type='boolean' value='true' migratable='yes'/>
    <property name='avx512bitalg' type='boolean' value='false'/>
    <property name='nodeid_msr' type='boolean' value='false'/>
    <property name='pdcm' type='boolean' value='false'/>
    <property name='movbe' type='boolean' value='true' migratable='yes'/>
    <property name='model' type='number' value='85'/>
    <property name='nrip_save' type='boolean' value='false'/>
    <property name='nrip-save' type='boolean' value='false'/>
    <property name='kvm_pv_unhalt' type='boolean' value='true' migratable='yes'/>
    <property name='ssse3' type='boolean' value='true' migratable='yes'/>
    <property name='sse4a' type='boolean' value='false'/>
    <property name='invpcid' type='boolean' value='false'/>
    <property name='pdpe1gb' type='boolean' value='false'/>
    <property name='tsc-deadline' type='boolean' value='true' migratable='yes'/>
    <property name='fma' type='boolean' value='true' migratable='yes'/>
    <property name='cx16' type='boolean' value='true' migratable='yes'/>
    <property name='de' type='boolean' value='true' migratable='yes'/>
    <property name='enforce' type='boolean' value='false'/>
    <property name='stepping' type='number' value='4'/>
    <property name='xsave' type='boolean' value='true' migratable='yes'/>
    <property name='clflush' type='boolean' value='true' migratable='yes'/>
    <property name='skinit' type='boolean' value='false'/>
    <property name='tsc' type='boolean' value='true' migratable='yes'/>
    <property name='tce' type='boolean' value='false'/>
    <property name='fpu' type='boolean' value='true' migratable='yes'/>
    <property name='ibs' type='boolean' value='false'/>
    <property name='ds_cpl' type='boolean' value='false'/>
    <property name='ds-cpl' type='boolean' value='false'/>
    <property name='host-phys-bits' type='boolean' value='false'/>
    <property name='fma4' type='boolean' value='false'/>
    <property name='la57' type='boolean' value='false'/>
    <property name='osvw' type='boolean' value='false'/>
    <property name='check' type='boolean' value='true' migratable='yes'/>
    <property name='hv-spinlocks' type='number' value='-1'/>
    <property name='pmu' type='boolean' value='false'/>
    <property name='pmm' type='boolean' value='false'/>
    <property name='apic' type='boolean' value='true' migratable='yes'/>
    <property name='spec-ctrl' type='boolean' value='false'/>
    <property name='min-xlevel2' type='number' value='0'/>
    <property name='tsc-adjust' type='boolean' value='true' migratable='yes'/>
    <property name='tsc_adjust' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-steal-time' type='boolean' value='true' migratable='yes'/>
    <property name='kvm_steal_time' type='boolean' value='true' migratable='yes'/>
    <property name='kvmclock' type='boolean' value='true' migratable='yes'/>
    <property name='l3-cache' type='boolean' value='true' migratable='yes'/>
    <property name='lwp' type='boolean' value='false'/>
    <property name='ibpb' type='boolean' value='false'/>
    <property name='xop' type='boolean' value='false'/>
    <property name='avx' type='boolean' value='true' migratable='yes'/>
    <property name='ospke' type='boolean' value='false'/>
    <property name='ace2' type='boolean' value='false'/>
    <property name='avx512bw' type='boolean' value='true' migratable='yes'/>
    <property name='acpi' type='boolean' value='false'/>
    <property name='hv-vapic' type='boolean' value='false'/>
    <property name='fsgsbase' type='boolean' value='true' migratable='yes'/>
    <property name='ht' type='boolean' value='false'/>
    <property name='nx' type='boolean' value='true' migratable='yes'/>
    <property name='pclmulqdq' type='boolean' value='true' migratable='yes'/>
    <property name='mmxext' type='boolean' value='false'/>
    <property name='vaes' type='boolean' value='false'/>
    <property name='popcnt' type='boolean' value='true' migratable='yes'/>
    <property name='xsaves' type='boolean' value='true' migratable='yes'/>
    <property name='tcg-cpuid' type='boolean' value='true' migratable='yes'/>
    <property name='lm' type='boolean' value='true' migratable='yes'/>
    <property name='umip' type='boolean' value='false'/>
    <property name='pse' type='boolean' value='true' migratable='yes'/>
    <property name='avx2' type='boolean' value='true' migratable='yes'/>
    <property name='sep' type='boolean' value='true' migratable='yes'/>
    <property name='pclmuldq' type='boolean' value='true' migratable='yes'/>
    <property name='x-hv-max-vps' type='number' value='-1'/>
    <property name='nodeid-msr' type='boolean' value='false'/>
    <property name='kvm' type='boolean' value='true' migratable='yes'/>
    <property name='misalignsse' type='boolean' value='false'/>
    <property name='min-xlevel' type='number' value='2147483656'/>
    <property name='kvm-pv-unhalt' type='boolean' value='true' migratable='yes'/>
    <property name='bmi2' type='boolean' value='true' migratable='yes'/>
    <property name='bmi1' type='boolean' value='true' migratable='yes'/>
    <property name='realized' type='boolean' value='false'/>
    <property name='tsc_scale' type='boolean' value='false'/>
    <property name='tsc-scale' type='boolean' value='false'/>
    <property name='topoext' type='boolean' value='false'/>
    <property name='hv-vpindex' type='boolean' value='false'/>
    <property name='xlevel2' type='number' value='0'/>
    <property name='clflushopt' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-no-smi-migration' type='boolean' value='false'/>
    <property name='monitor' type='boolean' value='false'/>
    <property name='avx512er' type='boolean' value='false'/>
    <property name='pmm-en' type='boolean' value='false'/>
    <property name='pcid' type='boolean' value='true' migratable='yes'/>
    <property name='3dnow' type='boolean' value='false'/>
    <property name='erms' type='boolean' value='true' migratable='yes'/>
    <property name='lahf-lm' type='boolean' value='true' migratable='yes'/>
    <property name='lahf_lm' type='boolean' value='true' migratable='yes'/>
    <property name='vpclmulqdq' type='boolean' value='false'/>
    <property name='fxsr-opt' type='boolean' value='false'/>
    <property name='hv-synic' type='boolean' value='false'/>
    <property name='xstore' type='boolean' value='false'/>
    <property name='fxsr_opt' type='boolean' value='false'/>
    <property name='kvm-hint-dedicated' type='boolean' value='false'/>
    <property name='rtm' type='boolean' value='true' migratable='yes'/>
    <property name='lmce' type='boolean' value='true' migratable='yes'/>
    <property name='hv-time' type='boolean' value='false'/>
    <property name='perfctr-nb' type='boolean' value='false'/>
    <property name='perfctr_nb' type='boolean' value='false'/>
    <property name='ffxsr' type='boolean' value='false'/>
    <property name='rdrand' type='boolean' value='true' migratable='yes'/>
    <property name='rdseed' type='boolean' value='true' migratable='yes'/>
    <property name='avx512-4vnniw' type='boolean' value='false'/>
    <property name='vmx' type='boolean' value='false'/>
    <property name='vme' type='boolean' value='true' migratable='yes'/>
    <property name='dtes64' type='boolean' value='false'/>
    <property name='mtrr' type='boolean' value='true' migratable='yes'/>
    <property name='rdtscp' type='boolean' value='true' migratable='yes'/>
    <property name='pse36' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-pv-tlb-flush' type='boolean' value='true' migratable='yes'/>
    <property name='tbm' type='boolean' value='false'/>
    <property name='wdt' type='boolean' value='false'/>
    <property name='pause_filter' type='boolean' value='false'/>
    <property name='sha-ni' type='boolean' value='false'/>
    <property name='model-id' type='string' value='Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz'/>
    <property name='abm' type='boolean' value='true' migratable='yes'/>
    <property name='avx512pf' type='boolean' value='false'/>
    <property name='xstore-en' type='boolean' value='false'/>
  </hostCPU>
  <hostCPU type='tcg' model='base' migratability='yes'>
    <property name='phys-bits' type='number' value='0'/>
    <property name='core-id' type='number' value='-1'/>
    <property name='xlevel' type='number' value='2147483658'/>
    <property name='cmov' type='boolean' value='true' migratable='yes'/>
    <property name='ia64' type='boolean' value='false'/>
    <property name='aes' type='boolean' value='true' migratable='yes'/>
    <property name='mmx' type='boolean' value='true' migratable='yes'/>
    <property name='rdpid' type='boolean' value='false'/>
    <property name='arat' type='boolean' value='true' migratable='yes'/>
    <property name='gfni' type='boolean' value='false'/>
    <property name='pause-filter' type='boolean' value='false'/>
    <property name='xsavec' type='boolean' value='false'/>
    <property name='intel-pt' type='boolean' value='false'/>
    <property name='osxsave' type='boolean' value='false'/>
    <property name='hv-frequencies' type='boolean' value='false'/>
    <property name='tsc-frequency' type='number' value='0'/>
    <property name='xd' type='boolean' value='true' migratable='yes'/>
    <property name='hv-vendor-id' type='string' value=''/>
    <property name='kvm-asyncpf' type='boolean' value='false'/>
    <property name='kvm_asyncpf' type='boolean' value='false'/>
    <property name='perfctr_core' type='boolean' value='false'/>
    <property name='perfctr-core' type='boolean' value='false'/>
    <property name='mpx' type='boolean' value='true' migratable='yes'/>
    <property name='pbe' type='boolean' value='false'/>
    <property name='decodeassists' type='boolean' value='false'/>
    <property name='avx512cd' type='boolean' value='false'/>
    <property name='sse4_1' type='boolean' value='true' migratable='yes'/>
    <property name='sse4.1' type='boolean' value='true' migratable='yes'/>
    <property name='sse4-1' type='boolean' value='true' migratable='yes'/>
    <property name='family' type='number' value='6'/>
    <property name='vmware-cpuid-freq' type='boolean' value='true' migratable='yes'/>
    <property name='avx512f' type='boolean' value='false'/>
    <property name='msr' type='boolean' value='true' migratable='yes'/>
    <property name='mce' type='boolean' value='true' migratable='yes'/>
    <property name='mca' type='boolean' value='true' migratable='yes'/>
    <property name='hv-runtime' type='boolean' value='false'/>
    <property name='xcrypt' type='boolean' value='false'/>
    <property name='thread-id' type='number' value='-1'/>
    <property name='min-level' type='number' value='13'/>
    <property name='xgetbv1' type='boolean' value='true' migratable='yes'/>
    <property name='cid' type='boolean' value='false'/>
    <property name='hv-relaxed' type='boolean' value='false'/>
    <property name='hv-crash' type='boolean' value='false'/>
    <property name='ds' type='boolean' value='false'/>
    <property name='fxsr' type='boolean' value='true' migratable='yes'/>
    <property name='xsaveopt' type='boolean' value='true' migratable='yes'/>
    <property name='xtpr' type='boolean' value='false'/>
    <property name='avx512vl' type='boolean' value='false'/>
    <property name='avx512-vpopcntdq' type='boolean' value='false'/>
    <property name='phe' type='boolean' value='false'/>
    <property name='extapic' type='boolean' value='false'/>
    <property name='3dnowprefetch' type='boolean' value='false'/>
    <property name='avx512vbmi2' type='boolean' value='false'/>
    <property name='cr8legacy' type='boolean' value='true' migratable='yes'/>
    <property name='cpuid-0xb' type='boolean' value='true' migratable='yes'/>
    <property name='xcrypt-en' type='boolean' value='false'/>
    <property name='kvm_pv_eoi' type='boolean' value='false'/>
    <property name='apic-id' type='number' value='4294967295'/>
    <property name='pn' type='boolean' value='false'/>
    <property name='dca' type='boolean' value='false'/>
    <property name='vendor' type='string' value='AuthenticAMD'/>
    <property name='pku' type='boolean' value='true' migratable='yes'/>
    <property name='smx' type='boolean' value='false'/>
    <property name='cmp_legacy' type='boolean' value='false'/>
    <property name='cmp-legacy' type='boolean' value='false'/>
    <property name='node-id' type='number' value='-1'/>
    <property name='avx512-4fmaps' type='boolean' value='false'/>
    <property name='vmcb_clean' type='boolean' value='false'/>
    <property name='vmcb-clean' type='boolean' value='false'/>
    <property name='3dnowext' type='boolean' value='true' migratable='yes'/>
    <property name='hle' type='boolean' value='false'/>
    <property name='npt' type='boolean' value='false'/>
    <property name='memory' type='string' value='/machine/unattached/system[0]'/>
    <property name='clwb' type='boolean' value='true' migratable='yes'/>
    <property name='lbrv' type='boolean' value='false'/>
    <property name='adx' type='boolean' value='true' migratable='yes'/>
    <property name='ss' type='boolean' value='true' migratable='yes'/>
    <property name='pni' type='boolean' value='true' migratable='yes'/>
    <property name='svm_lock' type='boolean' value='false'/>
    <property name='svm-lock' type='boolean' value='false'/>
    <property name='pfthreshold' type='boolean' value='false'/>
    <property name='smep' type='boolean' value='true' migratable='yes'/>
    <property name='smap' type='boolean' value='true' migratable='yes'/>
    <property name='x2apic' type='boolean' value='false'/>
    <property name='avx512vbmi' type='boolean' value='false'/>
    <property name='avx512vnni' type='boolean' value='false'/>
    <property name='hv-stimer' type='boolean' value='false'/>
    <property name='i64' type='boolean' value='true' migratable='yes'/>
    <property name='flushbyasid' type='boolean' value='false'/>
    <property name='f16c' type='boolean' value='false'/>
    <property name='ace2-en' type='boolean' value='false'/>
    <property name='pat' type='boolean' value='true' migratable='yes'/>
    <property name='pae' type='boolean' value='true' migratable='yes'/>
    <property name='sse' type='boolean' value='true' migratable='yes'/>
    <property name='phe-en' type='boolean' value='false'/>
    <property name='kvm_nopiodelay' type='boolean' value='false'/>
    <property name='kvm-nopiodelay' type='boolean' value='false'/>
    <property name='tm' type='boolean' value='false'/>
    <property name='kvmclock-stable-bit' type='boolean' value='false'/>
    <property name='hypervisor' type='boolean' value='true' migratable='yes'/>
    <property name='socket-id' type='number' value='-1'/>
    <property name='pcommit' type='boolean' value='true' migratable='yes'/>
    <property name='syscall' type='boolean' value='true' migratable='yes'/>
    <property name='level' type='number' value='13'/>
    <property name='avx512dq' type='boolean' value='false'/>
    <property name='svm' type='boolean' value='true' migratable='yes'/>
    <property name='full-cpuid-auto-level' type='boolean' value='true' migratable='yes'/>
    <property name='hv-reset' type='boolean' value='false'/>
    <property name='invtsc' type='boolean' value='false'/>
    <property name='sse3' type='boolean' value='true' migratable='yes'/>
    <property name='sse2' type='boolean' value='true' migratable='yes'/>
    <property name='est' type='boolean' value='false'/>
    <property name='avx512ifma' type='boolean' value='false'/>
    <property name='tm2' type='boolean' value='false'/>
    <property name='kvm-pv-eoi' type='boolean' value='false'/>
    <property name='cx8' type='boolean' value='true' migratable='yes'/>
    <property name='kvm_mmu' type='boolean' value='false'/>
    <property name='kvm-mmu' type='boolean' value='false'/>
    <property name='sse4_2' type='boolean' value='true' migratable='yes'/>
    <property name='sse4.2' type='boolean' value='true' migratable='yes'/>
    <property name='sse4-2' type='boolean' value='true' migratable='yes'/>
    <property name='pge' type='boolean' value='true' migratable='yes'/>
    <property name='fill-mtrr-mask' type='boolean' value='true' migratable='yes'/>
    <property name='avx512bitalg' type='boolean' value='false'/>
    <property name='nodeid_msr' type='boolean' value='false'/>
    <property name='pdcm' type='boolean' value='false'/>
    <property name='movbe' type='boolean' value='true' migratable='yes'/>
    <property name='model' type='number' value='6'/>
    <property name='nrip_save' type='boolean' value='false'/>
    <property name='nrip-save' type='boolean' value='false'/>
    <property name='kvm_pv_unhalt' type='boolean' value='false'/>
    <property name='ssse3' type='boolean' value='true' migratable='yes'/>
    <property name='sse4a' type='boolean' value='true' migratable='yes'/>
    <property name='invpcid' type='boolean' value='false'/>
    <property name='pdpe1gb' type='boolean' value='true' migratable='yes'/>
    <property name='tsc-deadline' type='boolean' value='false'/>
    <property name='fma' type='boolean' value='false'/>
    <property name='cx16' type='boolean' value='true' migratable='yes'/>
    <property name='de' type='boolean' value='true' migratable='yes'/>
    <property name='enforce' type='boolean' value='false'/>
    <property name='stepping' type='number' value='3'/>
    <property name='xsave' type='boolean' value='true' migratable='yes'/>
    <property name='clflush' type='boolean' value='true' migratable='yes'/>
    <property name='skinit' type='boolean' value='false'/>
    <property name='tsc' type='boolean' value='true' migratable='yes'/>
    <property name='tce' type='boolean' value='false'/>
    <property name='fpu' type='boolean' value='true' migratable='yes'/>
    <property name='ibs' type='boolean' value='false'/>
    <property name='ds_cpl' type='boolean' value='false'/>
    <property name='ds-cpl' type='boolean' value='false'/>
    <property name='host-phys-bits' type='boolean' value='false'/>
    <property name='fma4' type='boolean' value='false'/>
    <property name='la57' type='boolean' value='true' migratable='yes'/>
    <property name='osvw' type='boolean' value='false'/>
    <property name='check' type='boolean' value='true' migratable='yes'/>
    <property name='hv-spinlocks' type='number' value='-1'/>
    <property name='pmu' type='boolean' value='false'/>
    <property name='pmm' type='boolean' value='false'/>
    <property name='apic' type='boolean' value='true' migratable='yes'/>
    <property name='spec-ctrl' type='boolean' value='false'/>
    <property name='min-xlevel2' type='number' value='0'/>
    <property name='tsc-adjust' type='boolean' value='false'/>
    <property name='tsc_adjust' type='boolean' value='false'/>
    <property name='kvm-steal-time' type='boolean' value='false'/>
    <property name='kvm_steal_time' type='boolean' value='false'/>
    <property name='kvmclock' type='boolean' value='false'/>
    <property name='l3-cache' type='boolean' value='true' migratable='yes'/>
    <property name='lwp' type='boolean' value='false'/>
    <property name='ibpb' type='boolean' value='false'/>
    <property name='xop' type='boolean' value='false'/>
    <property name='avx' type='boolean' value='false'/>
    <property name='ospke' type='boolean' value='true' migratable='yes'/>
    <property name='ace2' type='boolean' value='false'/>
    <property name='avx512bw' type='boolean' value='false'/>
    <property name='acpi' type='boolean' value='true' migratable='yes'/>
    <property name='hv-vapic' type='boolean' value='false'/>
    <property name='fsgsbase' type='boolean' value='true' migratable='yes'/>
    <property name='ht' type='boolean' value='false'/>
    <property name='nx' type='boolean' value='true' migratable='yes'/>
    <property name='pclmulqdq' type='boolean' value='true' migratable='yes'/>
    <property name='mmxext' type='boolean' value='true' migratable='yes'/>
    <property name='vaes' type='boolean' value='false'/>
    <property name='popcnt' type='boolean' value='true' migratable='yes'/>
    <property name='xsaves' type='boolean' value='false'/>
    <property name='tcg-cpuid' type='boolean' value='true' migratable='yes'/>
    <property name='lm' type='boolean' value='true' migratable='yes'/>
    <property name='umip' type='boolean' value='false'/>
    <property name='pse' type='boolean' value='true' migratable='yes'/>
    <property name='avx2' type='boolean' value='false'/>
    <property name='sep' type='boolean' value='true' migratable='yes'/>
    <property name='pclmuldq' type='boolean' value='true' migratable='yes'/>
    <property name='x-hv-max-vps' type='number' value='-1'/>
    <property name='nodeid-msr' type='boolean' value='false'/>
    <property name='kvm' type='boolean' value='true' migratable='yes'/>
    <property name='misalignsse' type='boolean' value='false'/>
    <property name='min-xlevel' type='number' value='2147483658'/>
    <property name='kvm-pv-unhalt' type='boolean' value='false'/>
    <property name='bmi2' type='boolean' value='true' migratable='yes'/>
    <property name='bmi1' type='boolean' value='true' migratable='yes'/>
    <property name='realized' type='boolean' value='false'/>
    <property name='tsc_scale' type='boolean' value='false'/>
    <property name='tsc-scale' type='boolean' value='false'/>
    <property name='topoext' type='boolean' value='false'/>
    <property name='hv-vpindex' type='boolean' value='false'/>
    <property name='xlevel2' type='number' value='0'/>
    <property name='clflushopt' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-no-smi-migration' type='boolean' value='false'/>
    <property name='monitor' type='boolean' value='true' migratable='yes'/>
    <property name='avx512er' type='boolean' value='false'/>
    <property name='pmm-en' type='boolean' value='false'/>
    <property name='pcid' type='boolean' value='false'/>
    <property name='3dnow' type='boolean' value='true' migratable='yes'/>
    <property name='erms' type='boolean' value='true' migratable='yes'/>
    <property name='lahf-lm' type='boolean' value='true' migratable='yes'/>
    <property name='lahf_lm' type='boolean' value='true' migratable='yes'/>
    <property name='vpclmulqdq' type='boolean' value='false'/>
    <property name='fxsr-opt' type='boolean' value='false'/>
    <property name='hv-synic' type='boolean' value='false'/>
    <property name='xstore' type='boolean' value='false'/>
    <property name='fxsr_opt' type='boolean' value='false'/>
    <property name='kvm-hint-dedicated' type='boolean' value='false'/>
    <property name='rtm' type='boolean' value='false'/>
    <property name='lmce' type='boolean' value='false'/>
    <property name='hv-time' type='boolean' value='false'/>
    <property name='perfctr-nb' type='boolean' value='false'/>
    <property name='perfctr_nb' type='boolean' value='false'/>
    <property name='ffxsr' type='boolean' value='false'/>
    <property name='rdrand' type='boolean' value='false'/>
    <property name='rdseed' type='boolean' value='false'/>
    <property name='avx512-4vnniw' type='boolean' value='false'/>
    <property name='vmx' type='boolean' value='false'/>
    <property name='vme' type='boolean' value='false'/>
    <property name='dtes64' type='boolean' value='false'/>
    <property name='mtrr' type='boolean' value='true' migratable='yes'/>
    <property name='rdtscp' type='boolean' value='true' migratable='yes'/>
    <property name='pse36' type='boolean' value='true' migratable='yes'/>
    <property name='kvm-pv-tlb-flush' type='boolean' value='false'/>
    <property name='tbm' type='boolean' value='false'/>
    <property name='wdt' type='boolean' value='false'/>
    <property name='pause_filter' type='boolean' value='false'/>
    <property name='sha-ni' type='boolean' value='false'/>
    <property name='model-id' type='string' value='QEMU TCG CPU version 2.5+'/>
    <property name='abm' type='boolean' value='true' migratable='yes'/>
    <property name='avx512pf' type='boolean' value='false'/>
    <property name='xstore-en' type='boolean' value='false'/>
  </hostCPU>
  <cpu type='kvm' name='max' usable='yes'/>
  <cpu type='kvm' name='host' usable='yes'/>
  <cpu type='kvm' name='base' usable='yes'/>
  <cpu type='kvm' name='qemu64' usable='yes'/>
  <cpu type='kvm' name='qemu32' usable='yes'/>
  <cpu type='kvm' name='phenom' usable='no'>
    <blocker name='mmxext'/>
    <blocker name='fxsr-opt'/>
    <blocker name='pdpe1gb'/>
    <blocker name='3dnowext'/>
    <blocker name='3dnow'/>
    <blocker name='sse4a'/>
    <blocker name='npt'/>
  </cpu>
  <cpu type='kvm' name='pentium3' usable='yes'/>
  <cpu type='kvm' name='pentium2' usable='yes'/>
  <cpu type='kvm' name='pentium' usable='yes'/>
  <cpu type='kvm' name='n270' usable='yes'/>
  <cpu type='kvm' name='kvm64' usable='yes'/>
  <cpu type='kvm' name='kvm32' usable='yes'/>
  <cpu type='kvm' name='coreduo' usable='yes'/>
  <cpu type='kvm' name='core2duo' usable='yes'/>
  <cpu type='kvm' name='athlon' usable='no'>
    <blocker name='mmxext'/>
    <blocker name='3dnowext'/>
    <blocker name='3dnow'/>
  </cpu>
  <cpu type='kvm' name='Westmere' usable='yes'/>
  <cpu type='kvm' name='Westmere-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='Skylake-Server' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='pdpe1gb'/>
  </cpu>
  <cpu type='kvm' name='Skylake-Server-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
    <blocker name='pdpe1gb'/>
  </cpu>
  <cpu type='kvm' name='Skylake-Client' usable='no'>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='kvm' name='Skylake-Client-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='SandyBridge' usable='yes'/>
  <cpu type='kvm' name='SandyBridge-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='Penryn' usable='yes'/>
  <cpu type='kvm' name='Opteron_G5' usable='no'>
    <blocker name='pdpe1gb'/>
    <blocker name='sse4a'/>
    <blocker name='misalignsse'/>
    <blocker name='xop'/>
    <blocker name='fma4'/>
    <blocker name='tbm'/>
  </cpu>
  <cpu type='kvm' name='Opteron_G4' usable='no'>
    <blocker name='pdpe1gb'/>
    <blocker name='sse4a'/>
    <blocker name='misalignsse'/>
    <blocker name='xop'/>
    <blocker name='fma4'/>
  </cpu>
  <cpu type='kvm' name='Opteron_G3' usable='no'>
    <blocker name='sse4a'/>
    <blocker name='misalignsse'/>
  </cpu>
  <cpu type='kvm' name='Opteron_G2' usable='yes'/>
  <cpu type='kvm' name='Opteron_G1' usable='yes'/>
  <cpu type='kvm' name='Nehalem' usable='yes'/>
  <cpu type='kvm' name='Nehalem-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='IvyBridge' usable='yes'/>
  <cpu type='kvm' name='IvyBridge-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='Haswell' usable='no'>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='kvm' name='Haswell-noTSX' usable='no'>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='kvm' name='Haswell-noTSX-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='Haswell-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='EPYC' usable='no'>
    <blocker name='sha-ni'/>
    <blocker name='mmxext'/>
    <blocker name='fxsr-opt'/>
    <blocker name='pdpe1gb'/>
    <blocker name='cr8legacy'/>
    <blocker name='sse4a'/>
    <blocker name='misalignsse'/>
    <blocker name='osvw'/>
  </cpu>
  <cpu type='kvm' name='EPYC-IBPB' usable='no'>
    <blocker name='sha-ni'/>
    <blocker name='mmxext'/>
    <blocker name='fxsr-opt'/>
    <blocker name='pdpe1gb'/>
    <blocker name='cr8legacy'/>
    <blocker name='sse4a'/>
    <blocker name='misalignsse'/>
    <blocker name='osvw'/>
    <blocker name='ibpb'/>
  </cpu>
  <cpu type='kvm' name='Conroe' usable='yes'/>
  <cpu type='kvm' name='Broadwell' usable='no'>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='kvm' name='Broadwell-noTSX' usable='no'>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='kvm' name='Broadwell-noTSX-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='Broadwell-IBRS' usable='no'>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='kvm' name='486' usable='yes'/>
  <cpu type='tcg' name='max' usable='yes'/>
  <cpu type='tcg' name='host' usable='no'>
    <blocker name='kvm'/>
  </cpu>
  <cpu type='tcg' name='base' usable='yes'/>
  <cpu type='tcg' name='qemu64' usable='yes'/>
  <cpu type='tcg' name='qemu32' usable='yes'/>
  <cpu type='tcg' name='phenom' usable='no'>
    <blocker name='fxsr-opt'/>
    <blocker name='npt'/>
  </cpu>
  <cpu type='tcg' name='pentium3' usable='yes'/>
  <cpu type='tcg' name='pentium2' usable='yes'/>
  <cpu type='tcg' name='pentium' usable='yes'/>
  <cpu type='tcg' name='n270' usable='yes'/>
  <cpu type='tcg' name='kvm64' usable='yes'/>
  <cpu type='tcg' name='kvm32' usable='yes'/>
  <cpu type='tcg' name='coreduo' usable='yes'/>
  <cpu type='tcg' name='core2duo' usable='yes'/>
  <cpu type='tcg' name='athlon' usable='yes'/>
  <cpu type='tcg' name='Westmere' usable='yes'/>
  <cpu type='tcg' name='Westmere-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='Skylake-Server' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='avx512f'/>
    <blocker name='avx512dq'/>
    <blocker name='rdseed'/>
    <blocker name='avx512cd'/>
    <blocker name='avx512bw'/>
    <blocker name='avx512vl'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='Skylake-Server-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='avx512f'/>
    <blocker name='avx512dq'/>
    <blocker name='rdseed'/>
    <blocker name='avx512cd'/>
    <blocker name='avx512bw'/>
    <blocker name='avx512vl'/>
    <blocker name='spec-ctrl'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='Skylake-Client' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='rdseed'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='Skylake-Client-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='rdseed'/>
    <blocker name='spec-ctrl'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='SandyBridge' usable='no'>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
  </cpu>
  <cpu type='tcg' name='SandyBridge-IBRS' usable='no'>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='Penryn' usable='yes'/>
  <cpu type='tcg' name='Opteron_G5' usable='no'>
    <blocker name='fma'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='misalignsse'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xop'/>
    <blocker name='fma4'/>
    <blocker name='tbm'/>
  </cpu>
  <cpu type='tcg' name='Opteron_G4' usable='no'>
    <blocker name='avx'/>
    <blocker name='misalignsse'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='xop'/>
    <blocker name='fma4'/>
  </cpu>
  <cpu type='tcg' name='Opteron_G3' usable='no'>
    <blocker name='misalignsse'/>
  </cpu>
  <cpu type='tcg' name='Opteron_G2' usable='yes'/>
  <cpu type='tcg' name='Opteron_G1' usable='yes'/>
  <cpu type='tcg' name='Nehalem' usable='yes'/>
  <cpu type='tcg' name='Nehalem-IBRS' usable='no'>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='IvyBridge' usable='no'>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
  </cpu>
  <cpu type='tcg' name='IvyBridge-IBRS' usable='no'>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='Haswell' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
  </cpu>
  <cpu type='tcg' name='Haswell-noTSX' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
  </cpu>
  <cpu type='tcg' name='Haswell-noTSX-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='Haswell-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='spec-ctrl'/>
  </cpu>
  <cpu type='tcg' name='EPYC' usable='no'>
    <blocker name='fma'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='rdseed'/>
    <blocker name='sha-ni'/>
    <blocker name='fxsr-opt'/>
    <blocker name='misalignsse'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='osvw'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='EPYC-IBPB' usable='no'>
    <blocker name='fma'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='rdseed'/>
    <blocker name='sha-ni'/>
    <blocker name='fxsr-opt'/>
    <blocker name='misalignsse'/>
    <blocker name='3dnowprefetch'/>
    <blocker name='osvw'/>
    <blocker name='ibpb'/>
    <blocker name='xsavec'/>
  </cpu>
  <cpu type='tcg' name='Conroe' usable='yes'/>
  <cpu type='tcg' name='Broadwell' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='rdseed'/>
    <blocker name='3dnowprefetch'/>
  </cpu>
  <cpu type='tcg' name='Broadwell-noTSX' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rdseed'/>
    <blocker name='3dnowprefetch'/>
  </cpu>
  <cpu type='tcg' name='Broadwell-noTSX-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rdseed'/>
    <blocker name='spec-ctrl'/>
    <blocker name='3dnowprefetch'/>
  </cpu>
  <cpu type='tcg' name='Broadwell-IBRS' usable='no'>
    <blocker name='fma'/>
    <blocker name='pcid'/>
    <blocker name='x2apic'/>
    <blocker name='tsc-deadline'/>
    <blocker name='avx'/>
    <blocker name='f16c'/>
    <blocker name='rdrand'/>
    <blocker name='hle'/>
    <blocker name='avx2'/>
    <blocker name='invpcid'/>
    <blocker name='rtm'/>
    <blocker name='rdseed'/>
    <blocker name='spec-ctrl'/>
    <blocker name='3dnowprefetch'/>
  </cpu>
  <cpu type='tcg' name='486' usable='yes'/>
  <machine name='pc-i440fx-2.12' alias='pc' hotplugCpus='yes' maxCpus='255'/>
  <machine name='isapc' hotplugCpus='yes' maxCpus='1'/>
  <machine name='pc-1.1' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-1.2' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-1.3' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.8' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-1.0' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.9' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.6' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.7' hotplugCpus='yes' maxCpus='255'/>
  <machine name='xenfv' hotplugCpus='yes' maxCpus='128'/>
  <machine name='pc-i440fx-2.3' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.4' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.5' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.1' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.2' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.0' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.11' hotplugCpus='yes' maxCpus='288'/>
  <machine name='pc-q35-2.12' alias='q35' hotplugCpus='yes' maxCpus='288'/>
  <machine name='xenpv' maxCpus='1'/>
  <machine name='pc-q35-2.10' hotplugCpus='yes' maxCpus='288'/>
  <machine name='pc-i440fx-1.7' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.9' hotplugCpus='yes' maxCpus='288'/>
  <machine name='pc-0.15' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-1.5' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.7' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-1.6' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.11' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.8' hotplugCpus='yes' maxCpus='288'/>
  <machine name='pc-0.13' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-0.14' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.4' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.5' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-q35-2.6' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-1.4' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-i440fx-2.10' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-0.11' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-0.12' hotplugCpus='yes' maxCpus='255'/>
  <machine name='pc-0.10' hotplugCpus='yes' maxCpus='255'/>
</qemuCaps>














