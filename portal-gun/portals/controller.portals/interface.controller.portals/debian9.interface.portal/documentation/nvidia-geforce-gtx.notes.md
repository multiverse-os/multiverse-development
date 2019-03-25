# GeForce GTX Passthrough
The first successful Multiverse OS install was done on an `i5` with a earlier GTX series and later successfully upgraded to the GTX 970. 

```
01:00.0 VGA compatible controller [0300]: NVIDIA Corporation GM204 [GeForce GTX 970] [10de:13c2] (rev a1)
01:00.1 Audio device [0403]: NVIDIA Corporation GM204 High Definition Audio Controller [10de:0fbb] (rev a1)
```

Due to changes made to the Nvida drivers by Nvida developers to push consumers to buy `cloud/server` GPU cards to work with KVM/QEMU some modifications to the libvirt VM xml file must be made for the GPU PCI passthrough to work successfully.

## Required XML Changes
We start by loading the xml file using virsh into vim, with the following command:


#### Adding Qemu Command-Line Arguments
```
virsh edit debian9.controller.multiverse
```

Then start with editing the opening `<domain>` element, adding in a custom schema to support passing qemu command-line arguments using the libvirt xml defintion. The resulting tag will be:

```
<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
```

Unfortunately they do not provide secure `https` access, and are required to only access the schema via `http`.

**DEVELOPMENT** If Multiverse OS provides its own schema it can be served over `https` to improve the security and prevent tampering with schemas when downloading. 

The above can not be saved without adding the following, otherwise it will automatically be removed.

```
<qemu:commandline>
  <qemu:arg value='-cpu'/>
  <qemu:arg value='host,kvm=off'/>
</qemu:commandline>
```

#### Adding hidden KVM
The last step to getting the consumer/gaming GPU to work with PCI passthrough is adding the `<kvm>` feature to the `<features>` section. The resulting xml will look like:

```
<features>
  <acpi/>
  <apic/>
  <kvm>
    <hidden state='on'/>
  </kvm>
  <vmport state='off`/>
</features>
```

The other items located within the `<features>` element are dependent on your hardware and configuration, the important element for the purpose of PCI passthrough is just the following:

```
<kvm>
  <hidden state='on'/>
</kvm>
```

### Summary
With these two changes any modern Nvidia GPU, and not just the server/cloud GPUs will support PCI passthrough which is required for a primary Multiverse OS installation (the exception being satallite installations which only provide VM resources, and are not actively used by the user with a monitor, mouse and keyboard). 







