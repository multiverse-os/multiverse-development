# Portal Gun: Custom PCI devices 

Implemented possibly as a kernel module, or implemented in userspace. 

## Use something like VirtIO UDP Channel, possibly multicasing or using kcp
This would work out very well in combination with kernelspace devices

## Sound Device
Build a device that forward sound to Controller Sound device that syncs the sounds and plays them all together like SONOS

## vhost_net experimental_zcopytx=1
modprobe -r vhost_net
