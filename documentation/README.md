# Multiverse OS 
**Isolated, compartmentalized secure general-use operating system**

The Multiverse Operating System is intended to be an easy-to-use, general use operating system with reliable security provided by compartmentalized ephemeral VMs nested within isolated networks with secure microkernel VM routers. As much as possible, the complexity of the security model is transparently hidden within the UI of a full featured desktop environment based on Gnome or i3, capable of per-application pseudoanonymous full-VM sandboxing within ephemeral VMs. The network of virtual machines that support Multiverse OS can be deployed across several diverse bare-metal servers to create clusters of secure VMs.</style>

*Multiverse OS draws inspiration from previous secure operating systems such as Tails OS, Whonix and QubesOS. An easy way to conceptualize Multiverse OS design is that it is like running multiple Debian based QubesOS dom0 hosts in virtual machines in parallel. The user interacts with Multiverse OS from within a virtual machine with PCI devices passed, resulting in absolutely zero user interaction with the bare-metal host hypervisor. Each `controller-vm` running on a bare-metal host owns multiple customized, isolated networks. Like QubesOS, the user experience is designed to be like a traditional linux desktop environment experience, with as much of the complexity as possible hidden away from users.*


## Why 
Multiverse OS was created after realizing the need for compartmentalizing several open source development projects and to defend against the growing number of sophisticated vectors such as XSS frameworks combined with metasploit-like frameworks. 

After experimenting for months with Qubes OS, I realized this style of OS was the best available solution to the problem described above. However I decided it did not fit many of my requirements and could be vastly improved. So I started developing an alternative solution. My focus would be on making an operating system that fit the use cases of typical computer users, games, torrents, development, local file sharing and more. Some of the major criticisms I had of Qubes OS was the use of Fedora for routers and on the host, inability to pass through GPUs, passwordless setup, and most importantly how the host and virtual machines interact. I do not claim to be a security expert, this is an experiment but for my use case I believe it is better.

Use this at your own risk, do not trust this with your life or freedom. 

