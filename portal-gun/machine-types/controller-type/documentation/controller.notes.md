## Preformance


### L3 Cache
Now we dont need to pass `qemu-kvm` command-line arguments through `<qemu:commandline>` element, and instead we can define L3, L2 and L1 CPU cache passthrough directly in the `<cpu>` element. 

kvm=off can be acheived using kvm hidden=true. 

````
<qemu:commandline>
  <qemu:arg value='-cpu'/>
  <qemu:arg value='host,kvm=off,l3-cache=on'/>
</qemu:commandline>
````



