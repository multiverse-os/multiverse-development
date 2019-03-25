# Epeheramlity of VMs
Originally we were going to rely on `snapshots` supplied by the `QCOW2` VM image type. But using the `OVMF` BIOS prevents snapshots. So we can use `RAW` VM images when applicable becuse they are MUCH more preformant. 

This led to the idea of manually copying the ISO images, then splitting the VM image up into layers, that way VM snapshots could be generated in waves. Which would work decently. 

### Researching Two (2) Techniques to provide reliable ephemerality to Multiverse OS VMs
Further work needs to be done on reseaching the two strategies detailed below; determine the main differences betwen the two strategies, the pros/cons and how community feels about these.


[1][QCOW/RAW disk images, connected via VirtIO, minimum OS size, read-only] 
The most recent strategy is to create a READ-ONLY base image, for example, an APP VM requiring debian 9 desktop, would be 10 GB, or 6 GB for service VM. Then the second layer can be added to provide identity or user level settings and files. (This strategy is still being tested for efficacy and general ease of use and so on)

[2][Epehermality based portal gun building LIVE CDs for each VMs]
make PORTAL GUN produce "portals" or LIVE CDs that are ehemeral and transportable. Making these buildable and trustless makes them epic. It also ensures that the BIOS, initramfs, etc is all packaged with the system and that it is READONLY.

Some of the main benefits incude transportability, it would be easy to use these on remote machines.

This was a plan at one point but the exact way it could be accomplished was never traced out:

````
# 


````
