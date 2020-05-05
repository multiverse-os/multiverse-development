# VM Concealment
While it may be benficial to not hide the fact that the VMs are not actual machines because some viruses will avoid activating on machines they detect as VMs because they may fear they are in an environment for testing. The best way to do this is to at least attempt to not appear to be a VM, because it will look even more like this environment, and in some conditions it will appear as if its not a VM. We can do this while still using `virtio`. 


One method is converting our disks using virtio to `sdX`.


Hiding KVM. 

And we will have to run software that dtects VMs and ensure we trick them.


Then in the end this can be an option provided to users but by default we should try to hide our layers upon layers of VMs. 
