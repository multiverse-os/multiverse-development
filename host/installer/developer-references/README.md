## Multiverse OS: Installer
*This README.md is out of date and will either be removed or updated to help explain the procses of upgrading the manual installation guides developers used throughout research and development of Multiverse OS, the eventual shell scripts to simplify the process and the compilation and conversion to Go and an addition of a UI to produce the first Multiverse alpha installer.*

Multiverse OS originally was going to rely on the Debian9 installer but as the complexity of the project grew it became clear it would be better to simply implement a installer that would be consistent with the rest of the primary UI components and fit in with the rest of the machine building security precuations.

This would allow for example, simplified installation from ISO images, automate the process of validating checksums and signatures, to ensure that these steps are not skipped at any part of the installation process, but this is just one of many security examples that would benefit from consistency across the entire install and setup process.


## Packages
A list and description of various packages or repoistories to organize development of the Installer.

  * *base-files* - often found in OS git repositories, a repository containing folders and configuration files, like, configuration files that are put into folders like `/etc/*`. Like `motd`, or `issue`, but also other important files that are not just cosmetic.

  * *bin-utils* and *coreutils* - a collection of binaries needed for the host machine for installation, general operation, maintainence, and updating. Also a set for self-desctruction, to remove the install and revert to a normal default Debian installation.

  * *libpcap* - low level interaction with packets. *A modified version of this, maybe implemented as a kernel module could prevent packets from leaving or tampering with them to ensure networking does not work on the HOST*.


**This is likely best to move into controller instead of being on the host. The host should be locked down as much as possible and literally have no ability to control virtual machines beyond force closing their procceses, and not even that if possible.** 
  * *libmad* - mpeg audio decoder library, 24-bit pcm output, 100% fixed-point (integer) computation, based on new ISO/IEC standards, GNU. *Could be the way we link sound betwen devices.*


