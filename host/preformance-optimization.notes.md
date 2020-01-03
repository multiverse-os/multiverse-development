# Preformance Optimization Notes & Research


## CPU Throttling 

The value 146500 is the max rate allowed, and the value 0 works on some systems and if this does not improve the preformance, try setting it to 100.

```
kernel.perf_cpu_time_max_percent=0
kernel.perf_event_max_sample_rate=146500
```
