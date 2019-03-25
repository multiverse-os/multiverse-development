# NBD Exmaples with KVM



[DEBOOTRAP + NBD]
Tools:

  * debootstrap

  * networ block device (nbd)
    [enablings nbd kernel module][:] `modprobe nbd max_part=16`




Creating QEMU compliant QCOW2 images: 
  
  `kvm-img create -f qcow2 debian.qcow2 2G`

Using new QCOW2 for NBD: 

  `kvm-nbd -c /dev/nbd0 debian.qcow2`


Then we can partition the image and mount it:

````
sfdisk /dev/nbd0 -D -uM << EOF
,512,82
;
EOF
````


Using `debootstrap` to install debian on new HD:

````
deboostrap --include=less,locales-all,vim,sudo,openssh-ssh-server stable /mnt
http://ftp.us.debian.org/
````



Now wait for it to install, then we chroot to the new system, require all the special filesystems by mounting everything:

````
mount --bind /dev/ /mnt/dev
LANG=C chroot /mnt/ /bin/bash
mount -t proc none /proc
mount -t sysfs none /sys
````


Next install a kernel and grub 

````
apt-get install linux-image-amd64 grub
grub-install /dev/nbd0
update-grub
````

then setup root pass: `passwd root`


unload the image, fix grub, and done!

````
umount /proc/ /sys/ /dev/
exit
grub-install /dev/nbd0
```` 


finally edit `/mnt/boot/grub/grub.cfg` and replace `nbd0p2` for instances of `sda2`.


`umount /mnt`
`kvm -d /dev/nbd0`
