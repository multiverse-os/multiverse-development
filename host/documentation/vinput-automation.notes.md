# Automate VMs through Virtual Devices
Until we can move the host to later versions one way of doing this is:

```
<qemu:commandline>
  <qemu:arg value='-object'/>
  <qemu:arg value='input-linux,id=mouse1,evdev=/dev/input/by-id/{MOUSE_NAME}'/>
  <qemu:arg value='-object'/>
  <qemu:arg value='input-linux,id=kbd1,evdev=/dev/input/by-id/{KEYBOARD_NAME},grab_all=on,repeat=on'/>
</qemu:commandline>
```

Later versions makes this much easier with libvirt, but in general we are moving away from libvirt altogether and just using it as a stop gab solution. 

We are already testing with custom version of firecracker and qemu. 

