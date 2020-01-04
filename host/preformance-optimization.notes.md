# Preformance Optimization Notes & Research

## Use Q35 for the Controller

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

## CPU Pinning

## Isolate CPUs

## Define memory modules and make them unshared and private

## Specify and Tune NUMA

## CPU Throttling 

The value 146500 is the max rate allowed, and the value 0 works on some systems and if this does not improve the preformance, try setting it to 100.

```
kernel.perf_cpu_time_max_percent=0
kernel.perf_event_max_sample_rate=146500
```
