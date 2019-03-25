# Multiverse Shell & Account System
======
Beyond the collection of packages and binaries running ontop of Debian to provide ephemeral compartmentalized VMs. Multiverse OS also must provide complete system for selecting and managing basic shell interaction, bash/terminal features and all user settings.

------

## General System Configuration
Systemwide Multiverse OS configuration that is not dependent on Account configuration

  1. Updating host-\* systems Lockdown Scripts, system packages and core files using a proxy VM

  2. General Bash Scripts
    * Lockdown scripts, a collection of lockdown scripts run when the user is created, or each boot
    * Patch scripts to be run before and after installation, and update
   

## Account System Configuration
Multiverse OS should make parts of the shell, both desktop and bash terminals, which are typically only known and used by by advanced users easy to customize and use. This will empower even novice users to take full advantage of community Bash scripting while providing a better entry point into understanding these features of Linux. 

#### Dot-file Configurations
Accounts within Multiverse OS are made up of the following files, packagable, deployable, even remote deployment on rented hosting.

  1. Boot Configuration
    * Two factor authentication
      (a) Password based luksOpen for 'host-controller-vm' 
      (b) USB based luksOpen for 'controller-vm'
    * Securely delete Multiverse OS by modifying cryptsetup to begin booting decoy OS based on a *special alternative password*
    * Decoy OS booted accessed by modifying cryptsetup to begin booting decoy OS based on a *special alternative password*

  1. USB Management
    * USB Kill Script Settings
    * Allow/Deny list for USB serial #'s with defining scripts responses

  1. Repositories, for plugin or script manager similar to npm, gems, etc
    * VIM Plugins
    * Bash Scripts

  2. Bash Components
    * Bash Prompt
    * Bash Coloring

  4. General Bash Scripts
    * script repository and plugin management system
    * A way to quickly save and recall simple 1 line scripts
    * Categories and tree system (osint, conversions, development, ui-framework etc)
    * Lockdown scripts, a collection of lockdown scripts run when the user is created, or each boot
    * Logout scripts, a collection of scripts to clean up logs and remove ones presence
    * Delete user scripts, a collection of scripts run when a user is being removed

  6. Cryptographic Keys
    * Onion host keys (generate based on GPG key?)
    * SSH Key
    * SSL/TLS Key
      (a) Automatic lets encrypt retrieval and setup
    * GPG Key

  8. Preferred packages
    1. Which is started on boot?
    2. Software configuration repository:
      * VIM configuration
      * tmux configuration
      * mpd configuration
      * gnome-terminal configuration

  9.  Is developer?
    * What programming languages are to be added/removed by default

  10. Wallpaper options

  11. Desktop Environment
    * I3 plugins, configs repository
    * Gnome plugins repository

  12. Backup settings yaml, identifying items to be backed up (VMs, folders) and backup endpoints
  and method of connecting.

  14. End-point management: Manage automatic reverse SSH forwarding to registered remote end-points. Remote endpoints can be rented VPS servers or off-site servers. This allows interal VMs which route through isolating proxies to obtain firewall protected internet accessible ports providing essentially the same security as Cloudflare. 
    (a) A port forwarding UI very similar to the one found on a router but instead provides port forwarding from available endpoints, including onion services. 


------

## Multivberse Scripts & Binaries
List of unique scripts and binaries to improve the default terminal, shell, user interface functionality.


| **Completed** | **Scripts & Binaries**                                    |
|---------------------------------------------------------------------------|
|               | **Open Git Info**                                         |
|               | Open relevant git website, by url or git folder           |
|---------------------------------------------------------------------------|
|               | **Search from terminal**                                  |
|               | Open wiki, google, ddg, wolfram from CLI, either as lynx, |
|               | Tor, or Firefox.                                          |
|---------------------------------------------------------------------------|
|               | **CURL+sh check**                                         |
|               | Detect when using curl | sh, download and checksum files  |
|               | before running sh.                                        |
|---------------------------------------------------------------------------|
|               | **Onion Text Pass**                                       |
|               | Using Tor hidden service and a collection of contacts,    |
|---------------------------------------------------------------------------|
|               | **Title**                                                 |
|               | Description of script                                     |



