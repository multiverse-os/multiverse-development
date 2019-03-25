# Development Notes
Below are a collection of ntoes, ideas and research related to on-going development of Multiverse OS.

  * Support use of encrypting LUKs key, hiding the result in an image, and use this to unlock the Controller VM and Host VM. This can then serve the dual purpose as a kill-switch USB drive, that clears the memory and powers down if removed.

  * Each VM including Host machine should be generating SSH compatible ecdsa keys even if not used for SSH they will be useful for other tasks in inter-VM communication
	* echo -e $subheader"Unpriviledged User Creation"$reset

	* [!] Alpine does not have ssh-keygen when you do not install SSH, so the scramble keygen should provide GPG and SSH key generation 

  * Still trying to figure out the best way to handle machine images, where they should go, I like the idea of having each in its own folder but then each fodler would need to be a storage pool. The folder `Machines/Images` has been created as an option, or even just `./Images`.

  * A script that is capable of scaffolding/generating the base code for building a Multiverse OS machine of a specified type: [Controller VM, Service VM, Application VM, Router VM]

  * Move content in `./Data` to `./Machines` and move `\*.qcow2` images currently located in `./Machines` into their corresponding folder and symbolically link them into a top level folder like `./Storage` (name not yet determined).
    * When the folders are changed, downcase all the folder names

  * Build agents: long-running background management services that track resource stats, preform maintainence, and provide minimal control through a consistent API.
    * Router VM Agent
    * Controller VM Agent
    * Service VM Agent
    * Host Machine (bare-metal) Agent

  * Convert `provision.sh` shell scripts to `go` scripts and convert collected configuration files to templates that can change configuration content based on architecture and options selected in an eventual installer.

  * Generate onion address keys from (private_key + VM name (universe.router.multiverse)) to deterministically generate onion addresses.
 
## User Account
A minimum amount of Multiverse OS user account functionality must be implemented before the alpha installer can be completed:

  * Key generation and management
    * Scramble suit key
    * SSH key
    * GPG key

  * `pass-store` (The pass-store will hold all VM passwords)
    * Eventually will be replaced with an improved version that has an improved API that is consistent with other Multiverse OS APIs and supports newer password requirements like OTP/HOTP/TOTP. 

  * Write a `/dev/udev` file to properly set the permissions on `/dev/vfio/` files.

  * Write a script that pulls all `xml` configuration files from the **Host Machine** and organize them within the Multiverse OS structure to back up VM configurations. The reverse to install configurations would be a large part of installation automation.  

  * Develop a custom installer (`alpine-setup`) to simplify provisioning

______
## Name Brainstorming
Below are arrays containing used names, and potential names for various software that mesk up Multiverse OS. Names that make sense within the context of the Multiverse metaphor are preferred.

used_names = ["Scramble Suit", "Portal Gun", "Gravity", "Universe", "Gravity", "Portal", "Singularity", "c137"]
potential_names = ["Blackhole", "Quasar", "Singularity", "Sol", "Everett", "Black Star", "Metaverse", "Slider", "Wormhole", "Clark Kent Glasses", "Sun", "stealthboy", "ghost", "zerg", "zergling", "overlord", "overseer", "inventory"]
