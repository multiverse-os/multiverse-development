# Preformance Optimization Notes & Research

## Use Q35 for the Controller

## Set CPU Scale from powersave to performance

echo "performance" > /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor

**Ideally write a script that caculautes the number, then iterates through each and sets it. Quick research has not found a simple way to set this in sysctl.conf but thats the preferred location, so our hack quick fix is to just add this to our rc.local.*

Next we want to sent current to a higher number:

```
cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq 
```

Outputs 0 and we want a better starting point for our VMs.


## Nesting PCI Passhtrough Functionality


Add CPU feature vmx (for intel)


Under devices add 

```
<iommu model='intel'>
  <driver intremap='on' caching_mode='on' iotlb='on'/>
</iommu>
```

Add the following to grub kernel line (of the Controller):

```
iommu=pt intel_iommu=on
```

## Enable L3 Caching
This one is a major preformance increase

```
<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>

```

Allows for addition of below, which enables l3 cache. Despite what stackoverflow says, this is not impossible and this is very important for getting decent preformance.

```
  <qemu:commandline>
    <qemu:arg value='-cpu'/>
    <qemu:arg value='host,l3-cache=on'/>
  </qemu:commandline>
```


## CPU Pinning
Pinning is actually very simple but because libvirt is poorly designed it does not take of this simple task for you. So for the time being, and to better understand it first start with:

```
sudo lscpu -e
```
Results in the ouput:

```
CPU NODE SOCKET CORE L1d:L1i:L2:L3 ONLINE MAXMHZ    MINMHZ
0   0    0      0    0:0:0:0       yes    4500.0000 1200.0000
1   0    0      1    1:1:1:0       yes    4500.0000 1200.0000
2   0    0      2    2:2:2:0       yes    4500.0000 1200.0000
3   0    0      3    3:3:3:0       yes    4500.0000 1200.0000
4   0    0      4    4:4:4:0       yes    4500.0000 1200.0000
5   0    0      5    5:5:5:0       yes    4500.0000 1200.0000
6   0    0      6    6:6:6:0       yes    4500.0000 1200.0000
7   0    0      7    7:7:7:0       yes    4500.0000 1200.0000
8   0    0      8    8:8:8:0       yes    4500.0000 1200.0000
9   0    0      9    9:9:9:0       yes    4500.0000 1200.0000
10  0    0      0    0:0:0:0       yes    4500.0000 1200.0000
11  0    0      1    1:1:1:0       yes    4500.0000 1200.0000
12  0    0      2    2:2:2:0       yes    4500.0000 1200.0000
13  0    0      3    3:3:3:0       yes    4500.0000 1200.0000
14  0    0      4    4:4:4:0       yes    4500.0000 1200.0000
15  0    0      5    5:5:5:0       yes    4500.0000 1200.0000
16  0    0      6    6:6:6:0       yes    4500.0000 1200.0000
17  0    0      7    7:7:7:0       yes    4500.0000 1200.0000
18  0    0      8    8:8:8:0       yes    4500.0000 1200.0000
19  0    0      9    9:9:9:0       yes    4500.0000 1200.0000
```

This means that you use 0,10 because these are both within CORE 0. 1,11 because these are both within CORE 1. And so on. This will ensure your pairing of cores will be the nearest cores and will result in a massive preformance increase over just pinning random cores to a VM.

Also, it may be better to skip the CORE 0 0,10 cores and leave those to the host machine and assing starting from the second CORE 1.


## Memory Backing
This is for security improvements and speed. Huge pages improves preformance, while locked and nosharedpages are improvements to security.

```
<memoryBacking>
  <hugepages/>
  <nosharepages/>
  <locked/>
  <source type="file|anonymous|memfd"/>
  <access mode="shared|private"/>
  <allocation mode="immediate|ondemand"/>
  <discard/>
</memoryBacking>
```
Or huge pages can also be defined as:

```
<memoryBacking>
  <hugepages>
    <page size="1" unit="G" nodeset="0-3,5"/>
  </hugepages>
</memoryBacking>
```

In addition, there are many other options specified in libvirt, and while we are going to deprecate libvirt soon as possible, it can sometimes provide good ideas on what features to provide or what features QEMU provides. 



## KSM (Kernel Same Page Merging)
Ensure KSM is disabled, it is a security risk. But this is not included by default on Debian.

## Isolate CPUs
Another way to obtain major preformance gains regarding CPUs, is isolating specific CPU cores so that they will be utilized solely by the VM. This is important because it will ensure they are only used by the VM and not constantly switch context which cost a lot of overhead.

This is accomplished by modifying the grub kernel line, adding the CPUs you pinned, it can be some or all: 

```
BOOT_IMAGE="... intel_iommu=on isolcpus=1,11,2,12,3,13,4,14,5,15,6,16,7,17,8,18,9,19
```

## Enabling VHOST-NET Zero Copy
While we are working on our own better way either using FPGAs or other methods to do Guest-to-guest connections that completely bypass the host, an improvement fix that promises this functionality can be accomplished using the following in `/etc/modprobe.d/multiverse.conf`:

```
options vhost_net experimental_zcopytx=1
```

## Emulator Pins 
# NOTE: No idea exactly what this does, and haven't found a great explanation yet.

Another one is iothreadpin with iothread. And there are various tuning settings, preformance quotas, and several other things. 

Review libvirt, but in the end we want to abandon this and work directly with QEMU and our new firecracker fork. 

## Domain vCPU Threads & Sepcific NUMA Settings

## Using Cache Allocation Technology To Improve Preformance

## Define memory modules and make them unshared and private

## Specify and Tune NUMA

## CPU Throttling 

The value 146500 is the max rate allowed, and the value 0 works on some systems and if this does not improve the preformance, try setting it to 100.

```
kernel.perf_cpu_time_max_percent=0
kernel.perf_event_max_sample_rate=146500
```
