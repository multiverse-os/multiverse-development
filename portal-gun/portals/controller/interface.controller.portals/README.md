# Controller Virtual Machine

The **Controller VM** is the machine that functions as the interface for user interaction because in Multiverse OS the bare-metal host machine is completely inaccessible. All user interaction takes place within a virtual machine.

#### Filesystem Passthrough
The Controller VM requires two kernel modules to boot without errors:


#### Provisioning Controller VM
The foundation of every controlleris the same, then from the foundation, modifications are made to fulfill the specific requirements of the type of controller, for example a gaming controller would have steam installed by default.

To begin creating a foundational Debian/Ubuntu ControllerVM template some base packages need to be setup on the controller VM

```
sudo apt-get install vim tor ovmf virt-manager xpra python-netifaces
```

Then nano should be removed to shift vim to the default

```
sudo apt-get remove nano
```

Because the Controller VM runs nested Application VMs inside of it, we need some basics setup for launching virtual machines:

```
virbr0 with ip address range 10.10.10.0/24
virbr1 with ip address range 10.11.11.0/24
virbr2 with ip address range 10.12.12.0/24
```

And storage pools need to be setup:

```
'images' for HD images
'os-images' for os images
```

For mounting a 9p folder without erroring at the startup, modify 
`/etc/initramfs-tools/modules` and add the following modules:


```
9p
9pnet
9pnet_virtio
```

#### Setting up Xpra
Seamless access to applications launched within Application VMs is done using currently using `xpra` (until we can build a solution that uses /dev/shm between VMs). 

We want to use the beta version of `xpra`, the best way to obtain this is using the `xpra` apt repositories. 

```
wget -qO - https://xpra.org/gpg.asc | apt-key add -
cd /etc/apt/sources.list.d/
wget https://xpra.org/repos/buster/xpra.list
```

Once the repositories are installed run `sudo apt-get update` then install `xpra` and other packages necessary that are dependencies but don't get installed automatically. 

```
sudo apt-get install xpra python-netifaces python-pip python-cups

pip install numpy
pip install pyinotify
pip install opencv-python
pip install pyopengl
pip install pyopengl-accelerate
```
#### GPU Passthrough
If the Multivesre OS host is not being used as a sattalite Multiverse OS install then the GPU must be passed through to the Controller VM. 

**NOTE** A Satallite Multiverse OS install is a headless install, or an install on a computer that does not have a keyboard/mouse plugged in, or a server, or in other words not directly used but instead provides resources and joins a Multiverse OS super computer cluster to empower normal Multiverse OS installations. Sattalite installations can provide resources to both you and your friends, the resources can be shared and the system is designed to isolate and secure your software completely from any other users sharing resources on the satallite machine.

Keep in mind, that when doing GPU passthrough, you must install the video card driver **before** you actually pass through the card (at least the open source version, such as `nouveau` or `amd-gpu`). 

After the driver is installed, then you can passthrough the GPU PCI devices and isolate all your user activity inside of the Multiverse OS Controller VM.

**DEVELOPMENT** Infact, we should just detect what is the make of the card and ensure the appropraite `nouveau` driver, `amd-gpu` or available `nvidia-*` driver is (in the case of Ubuntu based Controller VMs). 

##### AMD
First step is adding `non-free` to the `/etc/apt/source.list` file

```
deb http://ftp.us.debian.org/debian/ buster main non-free
deb-src http://ftp.us.debian.org/debian/ buster main non-free

deb http://security.debian.org/debian-security buster/updates main non-free
deb-src http://security.debian.org/debian-security buster/updates main non-free

# buster-updates, previously known as 'volatile'
deb http://ftp.us.debian.org/debian/ buster-updates main non-free
deb-src http://ftp.us.debian.org/debian/ buster-updates main non-free

```

Then the firmware can be installed:

```
sudo apt-get update
sudo apt-get install firmware-amd-graphics xserver-xorg-video-amdgpu
```

Finally shut down the machine, remove the QXL/Spice devices and switch the the AMD passed through PCI GPU.


#### Feature Brainstorming 
Below are a list of potential features that can be implemented to improve and extend the functionality of the Controller VM.

  * Gnome extenions to interact with the **Router VMs** seamlessly.

  * Fallback version that functions with very limited resources to repair any issues with the other Controller VMs



## NOTES

Should be debian9-7
