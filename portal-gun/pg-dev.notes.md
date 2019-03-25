# Portal Gun Development Notes & Research
===============================================================================
Portal gun development has begun around the provisioning system which will
provide extremely low level access so that instead of relying on a third-party
or additional piece of software to produce images like vagrant or docker, that
often require trust of a third party who developed the images.

Multiverse OS virtual machine management system is built around the concept of
trustless images with secure booting sequence built in using the users keypairs
to sign the builds and produce the base for the ephemeral virtual machine images
that will be used until an update is releated and they need to be rebuilt.

Using full hardware virtualization and a boot sequence starting with open source
BIOS, the Multiverse OS virtual machines will provide a fully secure boot from
signed BIOS, to signed bootloader kernel, to signed and verified initramfs and
finally Multiverse OS host kernel. This provides a trustless platform for all
Multiverse OS virtual machines to be built on, providing a uniquely secure
experience requiring no trust of third parties.

The base of the ephemeral Multiverse OS images used by virtual machines are built
from vanilla popular linux install media that is verified with developer keys and
built ontop of the standard disk using simple and easy to understand provisioning
defined the the virtual machine configuration. Virtual machine configurations
can call in other simpler base configurations and build in complexity, using 
a tree of checkpoints that can be used as the base of a given virtual machine.

This provides a way for Multiverse OS users to share work on virtual machines
without requiring trust between anonymous Multiverse OS users since the images
are built locally, verified with checksums, built in a reproducible way that 
is verifiable at completion of the provisioning process. 

## Provisioning Development
===============================================================================
Development on the provisioning system will begin with creation of custom 
virtual keyboard and mouse packaged with console and video output scanning
to verify sucess of each step of the virtual device input, instead of relying
on flimsy sleeps and waits as used in services like Gnome Boxes. Instead a
basic console output verification system provides reliable and retryable 
interactions for each step of a image provisioning. 

This allows multiple retries to prompt the user, allowing on-the-fly
correcting of any provisoning which has been made invalid due to any sort
of updates. Changes to these provisioning steps can then be shared across
Multiverse OS users so that self-healing changes can be broadcasted to
trusted peers resulting in collaboritive updates to provisoning. 

**Uinput Based Virtual Devices**
Virtual devices will be built in pure Go implementations so that they do
not require any C libraries and have no dependencies. The devices will
function as a prototype for highly secure open souuce hardware being developed
by Multiverse OS lead developers. 


...

## VSOCK based VM-to-VM Networking
The next major priority is providing VSOCK based VM-to-VM networking
bypassing the need for host based switches/bridges; at the very least
for alpha release userspace fully unprivilidged software based bridges
that preferably run on the routers will be implemented to avoid use
of host controlled and connected bridges to ensure the host has no
ability to even accidently access the internet or be reached from
any other machine in the Multiverse OS internal cluster.
