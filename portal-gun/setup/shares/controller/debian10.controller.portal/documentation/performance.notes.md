# Multiverse OS Controller VM Performance
Multiverse OS is capable of having Controller VM performance near bare-metal speeds. After doing the below performance tuning, Multiverse OS developers were able to play modern FPS (Deus Ex: Human Revolution and Stellaris were both tested with perfect results) on their dedicated gaming Controller VMs.

There are even potential paths for development that will lead to further performance increases on all Multiverse OS virtual machines, especially Controller VMs, including but not limited to: memory-based screen relaying of Application VMs to the Controller VM, custom Multiverse OS tcp/ip network stack to avoid interaction with host machine kernelspace, and other planned improvements to the Multiverse OS system. 

The configuration changes discussed below will be applied by editing the libvirt XML:

```
virsh edit debian9.controller.multiverse
```

### Performance Tuning
Performance tuning gains come primarily from two main topics: (1) CPU Pinning, where we specify explicitly what CPU cores are dedicated to the VM, along with explicitly defining what CPU features we want passed through to the Controller VM (which enhances both security and performance); and (2) Clock/Tick configuration, where we define how the VM CPU should handle clock/ticks in a way that does not affect other VM performance. 

In addition to the above topics there are other enhancements that will be listed but the main boosts come from the above two configuration changes. This includes exposing `l3-cache` if the bare-metal host machine hardware supports it and other similar changes.

#### CPU Pinning
CPU Pinning is where we define explicitly what CPUs are configured to be used by the VM along with explicitly define what CPU features will be passed.

A key concept with CPU pinning is ensuring that correctly paired CPU cores are passed together. Because of `hyperthreading` provided by essentially all modern CPUs, each core has an associated pair, and very importantly the cores are not simply paired (1,2), (3,4), (5,6)... 

The physical core mapping can be found in the data output from `cat /proc/cpuinfo` but to save time the following command can be used to cut out just the relevant information:

```
paste <(cat /proc/cpuinfo | grep "core id") <(cat /proc/cpuinfo | grep "processor") | sort
```
The result will show the `core id` number along with the `processor` number, below is the output for an `i5` intel processor: 

```
core id : 0 processor : 0 
core id : 1 processor : 1
core id : 2 processor : 2
core id : 3 processor : 3
core id : 0 processor : 4
core id : 1 processor : 5
core id : 2 processor : 6
core id : 3 processor : 7
```

On this example `i5` machine the pairs look like:

```
# The first core is made up of two sub-cores numbered
0,4
# The second core is made up of two sub-cores numbered
1,5
# The third core is made of up two sub-cores numbered
2,6
# The fourth core is made of up two sub-cores numbered
3,7
```
When CPU pinning for Multiverse with the above `i5` we use three physical cores and their associated processors and leave the last for the routers and the host machine to use. Grouping by physical cores is the most efficient, otherwise it would pair cores in different physical locations causing lag. 

This results in the following XML configuration, placed below the `<currentMemory>` element. Note in this instance we are giving the VM two composite vCPUs (vCPU 6 and 7) that can use any CPU in the cpuset specified. This is optional, and may help with performance of VMs on machines that don't have many physical cores.

```
<vcpu placement='static' cpuset='1-3,5-7'>8</vcpu>
<cputune>
  <vcpupin vcpu='0' cpuset='3'/>
  <vcpupin vcpu='1' cpuset='7'/>
  <vcpupin vcpu='2' cpuset='1'/>
  <vcpupin vcpu='3' cpuset='5'/>
  <vcpupin vcpu='4' cpuset='2'/>
  <vcpupin vcpu='5' cpuset='6'/>
  <vcpupin vcpu='6' cpuset='1-3,5-7'/>
  <vcpupin vcpu='7' cpuset='1-3,5-7'/>
</cputune>
```

CPU Pinning is where Multiverse OS developers found the most significant performance boost outside of GPU passthrough and is an essential part of the Multiverse OS experience. 

#### Clock/Tick Configuration
The next boost comes from correctly configuring clocks/ticks. Eventually Multiverse OS developers plan to implement a customized version of the method used by `Xen` VMs to synchronize/update ticks, but until this can be completed, the following is the best configuration for Qemu/KVM clocks. The key point here is QEMU avoiding delays and correcting clock skew after the fact in combination with a script run on the guest VM. 

Below is optimized for the same example `Intel i5` used in the previous section:

```
<clock offset='utc'>
  <timer name='rtc' tickpolicy='catchup' track='guest'>
    <catchup threshold='123' slew='120' limit='10000'/>
  </timer>
  <timer name='pit' tickpolicy='discard'/>
  <timer name='hpet' present='no'/>
  <timer name='hypervclock' present='no'/>
</clock>
```

Multiverse OS developers always encourage learning more about the topic and refining the configuration to your device specifications. It is always better to understand configuration options over mindlessly setting them.

#### CPU Passthrough
One key step in setting up a Multiverse OS Controller is configuring CPU `host-passthrough` which allows us to explicitly define what CPU features we want passed to the controller VM. This allows for the Multiverse OS CPU to essentially be a custom set of CPU features, enabling us to not pass through CPU features that are potential security risks and only passing through features that are known to be secure and add performance. The net result is that this step increases both the performance and security of our Multiverse OS Controller VM. 

Below is optimized for the `Intel i5`:

```
<cpu mode='host-passthrough'>
  <topology sockets='1' cores='4' threads='2'/>
  <feature policy='disable' name='lahf_lm'/>
  <feature policy='require' name='fpu'/>
  <feature policy='require' name='pse'/>
  <feature policy='require' name='pse36'/>
  <feature policy='require' name='bmi2'/>
  <feature policy='require' name='rtm'/>
  <feature policy='require' name='lm'/>
  <feature policy='require' name='avx2'/>
  <feature policy='require' name='apic'/>
  <feature policy='require' name='mmx'/>
  <feature policy='require' name='aes'/>
  <feature policy='require' name='nx'/>
  <feature policy='require' name='pdpe1gb'/>
  <feature policy='require' name='clflush'/>
  <feature policy='require' name='vme'/>
  <feature policy='require' name='ss'/>
  <feature policy='require' name='avx'/>
  <feature policy='require' name='hle'/>
  <feature policy='require' name='erms'/>
  <feature policy='require' name='xsave'/>
  <feature policy='require' name='hypervisor'/>
  <feature policy='require' name='cx16'/>
  <feature policy='require' name='popcnt'/>
  <feature policy='require' name='movbe'/>
  <feature policy='require' name='sse2'/>
  <feature policy='require' name='ssse3'/>
  <feature policy='require' name='sse4.1'/>
  <feature policy='require' name='sse4.2'/>
</cpu>
```

#### Adding Features
Below are features added on an `i5` cpu to increase performance, review available features for your processor and add accordingly.

```
<features>
  ...
  <hap state='on'/>
  <pmu state='off'/>
  ...
</features>
```

**NOTE** The Multiverse OS developer who determined these features useful notes are mixed in with the old notes, so these need to be explained thoroughly so we are not mindlessly adding features we do not know. Everyone should always be fully informed what they are doing when configuring Multiverse OS machines.

#### Enabling L3 Cache
If your CPU supports L3 cache, you will benefit significantly by enabling it in the Multiverse OS controller VM. This is done by adding custom Qemu command-line arguments which requires the use of a custom schema.

The schema is changed by modifying the first `<domain>` element to:

```
<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
```

Unfortunately they do not provide secure `https` access, and are required to only access the schema via `http`.

This enables the use of the `<qemu:command-line>` tag used to pass custom command-line arguments to Qemu when launching the VM.

**DEVELOPMENT** If Multiverse OS provides its own schema it can be served over `https` to improve the security and prevent tampering with schemas when downloading. 

The above can not be saved without adding the following, otherwise it will automatically be removed.

```
<qemu:commandline>
  <qemu:arg value='-cpu'/>
  <qemu:arg value='host,l3-cache=on'/>
</qemu:commandline>
```

