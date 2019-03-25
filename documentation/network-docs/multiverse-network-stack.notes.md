# Multiverse OS: Network Stack
Multiverse OS uses a custom network stack for both networking between virtual machines (vms), host/virtual machine communication, and even vm/wan communication.

Multiverse OS uses this design in ordert completely bypass the Host Machine kernelspace when routing virtaul machine network packets. A significant portion of known hypervisor breakout attacks were implemented using a specially crafted packet sent through the Host Machine networking kernel space portions. By implementing the entire virtual machine network stack in userspace Multiverse OS adds moredepth to all the current defenses, limits the attack surface, enables rich networking features, and even obtains increased networking speed for virtaul machines.

## Multiverse Networking

## Onion Networking

