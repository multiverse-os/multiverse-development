# Libvirt XML Configuration Notes

define qmp ... unix:/run/qmp-sock


-m q36,accel=kvm


-mempath /dev/hugepages

<model fallback='forbid'>
<vendor>Intel</vendor>

Q35 is PCI-E native machine! Multiverse should likely not even support i440fx, and rip it out of qemu to lighten it up. 
