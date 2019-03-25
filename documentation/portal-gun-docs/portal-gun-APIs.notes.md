# Portal Gun APIs
# ===============
Multiverse OS provisioning and VM management system is named `portal-gun`, this component will provide multiple APIs to users, to simplify provisioning, cluster expansion, and so on. 


## Memory, assignment and automatic adjustment over time.
  [*][Direct memory access (DMA)]
  _[FOR DMA TO WORK]_ [we must have pre-allocated memory] which is critical for security anyways. SO PORTAL GUN SHOULD PRE_ALLOCATE MEMROY, dont over commit, divide by usage (use over time stats to adjust over time too, lets make this shti smart). 


  [*][do not use memory baloon]






____________________________________________________________________
## Timers and Clocks in QEMU

[Timer/Clock Reviews]
  * _pit_: programmable interval timer - a timer with periodic interrupts. [how is this used by the vm?]=

  * _rtc_ - real-time-clock - continiously running timer with preiodic interrupts

  * _tsc_ - time stamp counter -0 counts number of ticks since reset, no interrupts. [whats this used for?]


  * kvmclock - reocommend clock source for KVM guest, virtual machines KVM pvclock or kvmclock lets guest machine reads thost pshyyscial machines wall clock time. [!!]
   

