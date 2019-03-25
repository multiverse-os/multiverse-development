# Controller VM preformance tweaking
Geting some jittering with the new setup, its capable of running high end games but there may be some final tweaking required.

- JS is ridiculous, it blocks the UI of its own websites, its getting used more than ever, its super dangerous because of XSS, it seemingly has no limit on resources usage and people who program it are often novices who are not concerned about resource usage, or at least finish their project and need to be complained to before it even crosses their mind thats a problem to consider. They may never discover it relates to power usage or more importantly: battery drain. This is obvious considering the trend has been to push logic off the server and into the JS, meaning mobile devices use their battery faster than they need to if the design was more concious. That said: javascript usage can be confined and controlled in their amount of total resources allowed creating a hard limit on their ability to drag down a whole client computer. 

* Huge pages could still be used to squeeze out more power

* In the same way the CPUs are isolated, if we can isolate the memory used by the Host machine from the memory of the Controller and two router VMs. Then you remove another major vector for breakout attacks.

* A UI to select CPU features, name the CPU, and control it would be wundervoll, maybe even a full UI showing the virtual computer. Letting the user remove features like power button so for example the host cant turn it off or on. Clicking on the parts then being able to switch out the features. Letting people understand they are basically able to program the physical aspects of their computer.

* Control over cpufreq, maybe like when certain apps are turned on known to have profiles of larger usage, change the settings. When they are turned off, fallback. THis can save energy.

* Reflash bios on every boot, keep stored bios in a hidden drive that has read only access.

* The point is to extend the concept of multiple hop proxies, like with onion routing in Tor. Except with computers in a network, forcing a proxy between, each USB drive, between each Ethernet card, between any direct access to the Controller VM. Allowing ephemeral machines to function as proxies, so if they become infected, the infection is deleted when they are reset.

* The controller should have custom BIOS, custom CPU. Exposing only secure features needed, leaving the rest. Then a custom kernel, only providing features needed.

* Take steps (including testing with hypervisor detection tools) to make the VM undecteable and the HOST not detectable as a hypervisor. Includes (1) change MAC on network cards, (2) Change information on virtual components

* I found a scientific article explaining why Xen clocks are better than other VM software. Then they explained how to make an even better one than Xen. Implement this clock in Multiverse.

* ufw like tool for port forwding, opening ports on VMs, doing tunneling, multihop connecting

* The controller VM should not have any applications, anything that has not been established as a permentant Application VM, should open up in a disposable VM. The disposable VM would be configured similarly to the controller VM, so one could modify files and such, but the system files would all be thrown away when the X was hit. So any program like terminal or anything else.

* UI to visualize the networks. 

* auto wifi hacking- use this: https://github.com/xtr4nge/FruityWifi as a reference, kinda lame its php, make it better 

* [cool unique idea] Arm, and other micro controller VMs, built to mimic standard devices. Possible to insert sensors and so on, capable of feeding it random data. Simplyfing the process of programming robots and make the process cheaper. Once the code is done, you can buy all the pieces you need and have the code ready to run on the assembled bot.
  Also androids and other mobiesl

### Controller VM

Controller VM is setup to get the PCI graphics card the usb slots and so on. It also gets isolated CPUs. 

This not only vastly improves the preformance but it also lets you receive massive security benefits not seen in other VM/Container based compartmentalization operating systems. By telling the Host machine kernel not to schedule the Cores dedicated to the controller and by ensuring you correctly paired the hyperthreads when doing pinning. You ensure that the host and the controller use separate CPUs removing one of the biggest vectors of VM breakout. Exploits that take advantage of running code on the same CPU as the host and leveraging their shared access will not work shrinking the attack surface massively. 

One may ask why even use VMs, the primary reason is because we create a logical shim between the hardware and the virtual machine. Allowing programmatic control over BIOS flashing/updating/configuration, over CPU (you can just the CPUID features you want, leave out dangerous ones (watching CVE you can remove these making your CPU secure despite running a chip with known problems. like leave out management features, but keep all preformance features.) This promatic control over a baremetal machine would require robotic control to do everything one can with Multieverwse controller VM. This fine grain control provides a completely different user experience, now one does not need to wait for Intel, and gets more power than Intel in deciding what features their CPU has and is not forced to use the features sold. This is a powerful concept and fundamentally changes the security of general use computing.

Attacks against the machine would most likely infect the VM BIOS/firmware and not the host. And the BIOS can literally be reflashed on every boot, providing unpresented security. By walling off the HOST machine, ripping out everything extra and giving it no control over the Controller VM, malware in the host BIOS should have difficulty in attacking Multiverse OS.



## Utility VM Ideas
* more general but wget for whole folder?

* Backup server

* Interal DNS, ensure an onion request never leaves the network so it can be accidently sent over the internet and reveal dns information in the failed lookup

* USB proxy vms

* Electronic disobedience, easy set and forget DOS tools. Should apply attacks and check effect, use variety of tools, spit out all the recon it can do automatically and suggest attacks. Routes traffic through multihops to hide origin.

* Tor Non-Exit/Exit Relay

* Automatically convert pdfs to open formats
  - https://github.com/Debian/gpdftext


* Home web hosting, linking up with friends and even freinds of friends from contacts DB. If any site gets a ton of usage, people can share their resources and provide additional hosts, scaling up when use is larger and sclaing down when its not needed. This way people host form home but they also get the benefits of sclaing up in cloud centers. Lets actually ahve the internet not just be hosted by 1 provider creating a giant central point of weakness and essentially breaking the entire concept that makes the internet good

* [important to have early] setup to provide packages for OS upgrade of the host machine

https://github.com/strothj/alpine-debtool 


* Reproducible build VM SERVER not fucking container. Continious integration testing and so on.

## Key Ideas

* No logs on any computer in the cluster BUT the controller, all logs should be output to this server so its all centralized.
https://github.com/clustree/journalbeat-deb
 - https://github.com/mbenkmann/garcon
 - maybe a middleware for caddy!

# https://github.com/systemd/mkosi
**This is critical, in Multiverse, we want to build all the isos for the App and Utility VMs. This would allow for custom building of images. If this is not used it should be the guide. It even supports adding luks and other stuff to the iamge and multiple images.**

Unlike most bullshit like docker or even Qubes, we want the Apps to be distributed WITHOUT fucking prebuilt images that you have to trust. Removal of all trust is critical to security. So the images should be built from auditable buildscripts

*This may be even just useful for building images for VMs inside of Multiverse because it can make a wide variety of custom iso images.*

A fancy wrapper around dnf --installroot, debootstrap, pacstrap and zypper that may generate disk images with a number of bells and whistles.

In theory, any distribution may be used on the host for building images containing any other distribution, as long as the necessary tools are available. Specifically, any distro that packages debootstrap may be used to build Debian or Ubuntu images. Any distro that packages dnf may be used to build Fedora images. Any distro that packages pacstrap may be used to build Arch Linux images. Any distro that packages zypper may be used to build openSUSE images.

Additionally, bootable GPT disk images (as created with the --bootable flag) work when booted directly by EFI systems, for example in KVM via:

*Generated images are legacy-free. This means only GPT disk labels (and no MBR disk labels) are supported, and only systemd based images may be generated. Moreover, for bootable images only EFI systems are supported (not plain MBR/BIOS).*

qemu-kvm -m 512 -smp 2 -bios /usr/share/edk2/ovmf/OVMF_CODE.fd -drive format=raw,file=image.raw

