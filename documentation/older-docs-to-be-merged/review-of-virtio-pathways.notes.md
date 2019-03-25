## Types of VirtIO VM-to-VM communication
**Review 23.18 the "Devices" portion of the Red Hat documentation, which is by far the best VirtIO/Libvirt documentation.**

[VirtIO Channel: UDP Network Pathway]

[VirtIO Channel: Using HD acceleration via kernelspace channeling]

[VirtIO Channel: Named Pipe]

[VirtIO Channel: Spice Port, a type of VDI (virtual device ]
````
<deviecs>
  <serial type="spiceport">
    <source channel="qemu.console.serial.0"/>
    <target port="1" />
  </serial> 
</devices>
````


[VirtIO SCSI][Is not avaialble in virt-manager, so noobs havent played with it]

<devices>
  <controller type="scsi" index="0" model="virtio-scsi">
    <driver iothread="4">
    <address type="pci" domain="0x0000" bus="0x00" slot="0x0b" function="0x0" />
  </controller>
</devices>

</devices>

[VirtIO Console]





[VirtIO Serial]


[VirtIO Parallel]
