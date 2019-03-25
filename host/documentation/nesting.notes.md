## Nested Virtualization
Nested virtualization is a still experimental virtual machine feature that provides a lot of very interesting possibilities and is leveraged by Multiverse to encapsulate complexity and provide a unique form of isolation that we believe provides an unparalled level of security. 

**New Multiverse OS Operating Modes**
The Multiverse OS design now has two operating modes:

  1) The original, and preferred mode due to maximum isolation, modularity and enhanced security utilizing nested virtualization for application VMs running inside the controller VM. 

  2) The new "lighter weight" operating mode combines the responsibility of the controller VM, and functions more like Qubes OS. The host+controller responsibilities remain on the host and application VMs are run direclty on the host instead of running nested inside the controller VM. This is being designed specifically for running Multiverse on a wider variety of devices, intended for use with machines with less CPUs but still providing all the security enhanced features and functionlaity that comes with Multiverse OS by subtracting features that are either not available or not supported by older or more affordable hardware. 

Currently the alpha-installer is being built to support automated installation for both operating modes, in addition simple standalone software will be provided in conjuction with the installer to easily and automatically determine the best operating mode for a machine during installation. 

#### Preformance
To acheive expected bare-metal like preformance inside of application VMs running nested within the controller VM, more than just passing `nested=1` to the kvm kernel module is required. 

Below are efforts to consolidate notes from different developers and installation attempts into a multi-architecture guide that should cover all work up until this point and any new additions based on our growing understanding of the various level level aspects of the Multiverse OS design. 

**Kernel Configuration: Modules and Associated Options**





**SystemFS based configuration**
In addition to providing kernel module options to important VM related kernel modules, to acheive bare-metal like preformance some, arguably undesirable changes must be made to the host machine system configuration. The easiest way to do this is through sysfs.

One important reason the guests or nested guests run slow is because the host and guest by default slowly increments resource allocation up and down. In this specific case, we are referring to CPU resources, and explicitly CPU frequency scaling and power usage. CPU scaling is a featured designed to save power by limiting resource allocation, scaling up power usage and resources allocation as a device needs it. The amount of lag this creates on a single computer running a typical OS is neglible and by default this feature is desirable. 

However due to the nature of virtual machines being nested, the ability to change values *rapidly* from high-to-low and low-to-high is required, but it is not the default behavior .

Instead, by default, each tick the value is incremented up or down, towards the desired goal value. Importantly, the rate of change of the value is throttled by a per tick maximum increment/decrement limit. So if the value is needed to be 100 because the CPU is being used to run 2 VMs, and one of those 4 VMs nested inside, and currently the value is 15, with a maximum rate of change limit set to 5, it takes unnecessary wait time in order to allocate the CPU resources required to preform a given intense operation. This functionality is throttled by default because it is related to power usage, if the value is 100 all the time, then even when you are doing little the amount of power being used is greater. 

**We need to find a balance, a way to allocate resources quickly without needing to pre-allocate the CPU resource for quick preformance. We don't want to unnecssarily waste power for preformance boosts.**



**QEMU/NEMU Customization And Custom Multiverse OS**
This has not yet been required or implemented but customizing the software will provide a lot of functionality we desire and this can be experimented with and started by creating stand alone software implementation of desired changes that can run alongside the existing vanilla QEMU. 
