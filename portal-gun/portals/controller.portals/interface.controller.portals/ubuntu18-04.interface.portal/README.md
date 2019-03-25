# Application Virtual Machine (VM)
The Multiverse OS application virtual machine are nested virtual machines (VMs) **run inside the Controller VM.** Application VMs are similar to Multiverse OS Service VMs, but instead of running in the background without the need for user iteraction, Application VMs have their active window passed to the Controller VM.

Currently, `xpra` is used to pass the active window of the Application VM to the Controller VM. Indicators can be added to the windows to segregate windows by identity to make it clear what identity each VM is associated to. 

Furthermore, shared storage connects folders on the Application VMs to folders within the Controller VM to provide seamless interaction to Applications while keeping their operation completely segregated in their own ephemeral operating systems. This ensures that any malware infections are deleted after reset. For example, the Downloads folder is mounted to the `firefox` Application VM, so that any downloads in the Firefox VM appear in the Downloads folder of the Controller VM seamlessly. 

### Provisioning
Application VMs can be run ontop of Debian, Alpine, Ubuntu and Fedora. Curently we are support Debian and Alpine and support for other base templates will be expanded as development continues.

#### Debian Provisioning
Below are the setups to set up Debian Buster for use as an Application VM.

##### 1) Upgrade stretch to buster
In `/etc/apt/sources.list` change all repositories except "security" from `stretch` to `buster` and upgrade to buster.

```
sudo apt-get update
sudo apt-get dist-upgrade
```

**NOTE:** From debian.org buster release info page [](https://www.debian.org/releases/buster/):
> Please note that security updates for "testing" distribution are not yet managed by the security team. Hence, "testing" does not get security updates in a timely manner. You are encouraged to switch your sources.list entries from testing to stretch for the time being if you need security support. See also the entry in the Security Team's FAQ for the "testing" distribution.

##### 2) Install Xpra
The following must be done on the Controller VM as well as the Application VM.

Add the Xpra repo:

```
wget -qO - https://xpra.org/gpg.asc | apt-key add -
cd /etc/apt/sources.list.d/
wget https://xpra.org/repos/buster/xpra.list
```

Currently, we have to enable Debians's sid (unstable) repo for one dependency (`python-gtkglext1`), which was removed from the testing repo. (`python-gtkglext1` exists in stretch and sid, but not currently buster [](https://tracker.debian.org/news/927623/python-gtkglext1-removed-from-testing/).) Add the following lines to `/etc/apt/sources.list` (editing for your prefered mirror if necessary):

```
deb http://ftp.debian.org/debian/ sid main
deb-src http://ftp.debian.org/debian/ sid main
```

To only allow installs from `sid` if the dependency is missing from all other repos, pin sid's priority in `/etc/apt/preferences`. Higher number = higher priority, default is 500 and if two repos tie, the higher version will be installed. We don't need to explicitly set priority for all enabled repos, just ones that you want higher or lower than 500.

```
Package: *
Pin: release a=unstable
Pin-Priority: 400
```

(Note: Can check that sid is not the prefered repo for a given packge with `apt-cache policy packagename`)

```
sudo apt-get update
sudo apt-get install xpra python-netifaces python-pip python-cups python-numpy python-opengl python-opencv python-avahi
usermod -a -G xpra user

pip install numpy
pip install pyinotify
pip install opencv-python
pip install pyopengl
pip install pyopengl-accelerate
```

To test that xpra is sucessfully installed, from the Controller VM run `xpra start ssh:SERVERHOSTNAME:100 --start-child=xterm`, with SERVERHOSTNAME replaced by an SSH style user@host style address pointing to the Application VM. You shoud have a window pop up running an xterm instance from the Application VM.


##### 3) Enable Multiverse mounts
Add lines for `9p`,`9pnet`, and `9pnet_vfio` to `/etc/modules`

Add 9p mounts to `/etc/fstab`:

```
multiverse-downloads /media/user/Downloads 9p trans=virtio,9p2000.L,rw,posixacl,cache=none 0 0
#... and others if necessary
```
### Using Xpra
Windows running on the Multiverse OS application VMs are seamlessly transfered to the Controller VM using `xpra`. 

The command to open `gnome-terminal` is:

```
 xpra start ssh:app.firefox.debian:100 --start-child=gnome-terminal
```

With `gnome-terminal` you can open `firefox` or other programs for now, and the whole process will be hooked into Icons in the near future as development continues.


In libvirt (virt-manager, virsh edit, etc), edit the Application VM to add Filesystems pointing to the Multiverse folders desired.

### Development
Management of VMs will be handled by `portal-gun`, it will handle provisioning VMs, downloading new templates, updating, mounting `/var/multiverse/shared-storage` folders, and handling xpra window hand off. 

Portal gun will need to manage:

  * Setting the hostname for each Application VM in the controller hostfile

  * Starting xpra on client and server

  * Handle random generation of 32 char passwords for BOTH user and root, store in pass-store

  * Handle shared storage moutning, editing `/etc/fstab` and doing symbolic links to relevant locations
