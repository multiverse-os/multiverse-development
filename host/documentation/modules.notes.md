If hardware *like video card passthrough* doesn't support
*remapping of interruptions*, then the undesired and 
insecure option `allow_unsafe_assigned_interrupts=1` with
the kvm kernel module declaration. 

This is not required at boot time but probably should be in
`/etc/modules` over `/etc/modprobe.d/multiverse.conf`.

Ideally we have a stand-alone software to check various
important features, hardware, configurations. Then the 
output can be used to guide a rescue mode, installation, 
configuration, etc. ...

```
options kvm allow_unsafe_assigned_interrupts=1
```
