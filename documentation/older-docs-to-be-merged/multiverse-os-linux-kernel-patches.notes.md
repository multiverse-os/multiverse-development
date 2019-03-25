##
##  Multiverse OS: Linux Kernel Modifications
===================================================================
Changes the the Linux Kernel should be applied as a patch. The patches
should be systematically built using a patching tool that allows
easy creation of patches using a Ruby configruation system. 

After extensive research on the Linux Kernel during the Development of Multiverse OS, multiple significant patches have been developed to significantly improve the security of the Linux Kernel but additionally modify the core functionality of the monolithic Linux kernel to meet the design needs of Multiverse OS.
## RTOS 
Emulation and real RTOS
## Interesting Libs
https://github.com/deniswernert/udev - FULL udev replacement

## Random number generation
https://github.com/takatoh/boxmuller
## Coop Mode
systemd-logind.service **HANDLES MULTISEAT MANAGEMNT**
       /usr/lib/systemd/systemd-logind

## Proposed Multiverse OS Linux Kernel Patches
[1][Encrypted Stdout]
Initramfs and the Linux kernel are to be patched to encrypt all Stdout, Stderr and all other output with the active user's ephemeral session key.

http://man7.org/linux/man-pages/man5/limits.conf.5.html

[2][Signed Stdin & Systemd Interaction & udevadm]
Initramfs and the Linux kernel are to be patched to require all Stdin to be signed with the active user's ephemeral session key.


[systemd.system](http://man7.org/linux/man-pages/man7/systemd.special.7.html)
     A few units are treated specially by systemd. Many of them have
       special internal semantics and cannot be renamed, while others simply
       have a standard meaning and should be present on all systems.


[SOCKETS AND FIFOS](http://man7.org/linux/man-pages/man1/systemd.1.html)
       /run/systemd/notify
           Daemon status notification socket. This is an AF_UNIX datagram
           socket and is used to implement the daemon notification logic as
           implemented by sd_notify(3).
       /run/systemd/private
           Used internally as communication channel between systemctl(1) and
           the systemd process. This is an AF_UNIX stream socket. This
           interface is private to systemd and should not be used in
           external projects.
       /dev/initctl
           Limited compatibility support for the SysV client interface, as
           implemented by the systemd-initctl.service unit. This is a named
           pipe in the file system. This interface is obsolete and should
           not be used in new applications.

 A unit configuration file encodes information about a service, a
       socket, a device, a mount point, an automount point, a swap file or
       partition, a start-up target, a watched file system path, a timer
       controlled and supervised by systemd(1), a resource management slice
       or a group of externally created processes. The syntax is inspired by
       XDG Desktop Entry Specification[1].desktop files, which are in turn
       inspired by Microsoft Windows .ini files.

       This man page lists the common configuration options of all the unit
       types. These options need to be configured in the [Unit] or [Install]
       sections of the unit files.

       In addition to the generic [Unit] and [Install] sections described
       here, each unit may have a type-specific section, e.g. [Service] for
       a service unit. See the respective man pages for more information:
       systemd.service(5), systemd.socket(5), systemd.device(5),
       systemd.mount(5), systemd.automount(5), systemd.swap(5),
       systemd.target(5), systemd.path(5), systemd.timer(5),
       systemd.slice(5), systemd.scope(5).

[udevadm]
udevadm - udev management tool

[udev - Dynamic device management](http://man7.org/linux/man-pages/man7/udev.7.html)

SYNOPSIS

       udevadm [--debug] [--version] [--help]

       udevadm info options

       udevadm trigger [options]

       udevadm settle [options]

       udevadm control command

       udevadm monitor [options]

       udevadm test [options] devpath

       udevadm test-builtin [options] command devpath

DESCRIPTION

       udevadm expects a command and command specific options. It controls
       the runtime behavior of systemd-udevd, requests kernel events,
       manages the event queue, and provides simple debugging mechanisms.

OPTIONS

       --debug
           Print debug messages to standard error.

       --version
           Print version number.

       -h, --help
           Print help text.

   udevadm info [options] [devpath|file]
       Queries the udev database for device information stored in the udev
       database. It can also query the properties of a device from its sysfs
       representation to help creating udev rules that match this device.

       -q, --query=TYPE
           Query the database for the specified type of device data. It
           needs the --path or --name to identify the specified device.
           Valid TYPEs are: name, symlink, path, property, all.

       -p, --path=DEVPATH
           The /sys path of the device to query, e.g.
           [/sys]/class/block/sda. Note that this option usually is not very
           useful, since udev can guess the type of the argument, so udevadm
           --devpath=/class/block/sda is equivalent to udevadm
           /sys/class/block/sda.

       -n, --name=FILE
           The name of the device node or a symlink to query, e.g.
           [/dev]/sda. Note that this option usually is not very useful,
           since udev can guess the type of the argument, so udevadm
           --name=sda is equivalent to udevadm /dev/sda.

       -r, --root
           Print absolute paths in name or symlink query.

       -a, --attribute-walk
           Print all sysfs properties of the specified device that can be
           used in udev rules to match the specified device. It prints all
           devices along the chain, up to the root of sysfs that can be used
           in udev rules.

       -x, --export
           Print output as key/value pairs. Values are enclosed in single
           quotes.

       -P, --export-prefix=NAME
           Add a prefix to the key name of exported values.

       -d, --device-id-of-file=FILE
           Print major/minor numbers of the underlying device, where the
           file lives on.

       -e, --export-db
           Export the content of the udev database.

       -c, --cleanup-db
           Cleanup the udev database.

       --version
           Print version.

       -h, --help
           Print help text.

       In addition, an optional positional argument can be used to specify a
       device name or a sys path. It must start with /dev or /sys
       respectively.

   udevadm trigger [options] [devpath|file...]
       Request device events from the kernel. Primarily used to replay
       events at system coldplug time.

       -v, --verbose
           Print the list of devices which will be triggered.

       -n, --dry-run
           Do not actually trigger the event.

       -t, --type=TYPE
           Trigger a specific type of devices. Valid types are: devices,
           subsystems. The default value is devices.

       -c, --action=ACTION
           Type of event to be triggered. The default value is change.

       -s, --subsystem-match=SUBSYSTEM
           Trigger events for devices which belong to a matching subsystem.
           This option can be specified multiple times and supports shell
           style pattern matching.

       -S, --subsystem-nomatch=SUBSYSTEM
           Do not trigger events for devices which belong to a matching
           subsystem. This option can be specified multiple times and
           supports shell style pattern matching.

       -a, --attr-match=ATTRIBUTE=VALUE
           Trigger events for devices with a matching sysfs attribute. If a
           value is specified along with the attribute name, the content of
           the attribute is matched against the given value using shell
           style pattern matching. If no value is specified, the existence
           of the sysfs attribute is checked. This option can be specified
           multiple times.

       -A, --attr-nomatch=ATTRIBUTE=VALUE
           Do not trigger events for devices with a matching sysfs
           attribute. If a value is specified along with the attribute name,
           the content of the attribute is matched against the given value
           using shell style pattern matching. If no value is specified, the
           existence of the sysfs attribute is checked. This option can be
           specified multiple times.

       -p, --property-match=PROPERTY=VALUE
           Trigger events for devices with a matching property value. This
           option can be specified multiple times and supports shell style
           pattern matching.

       -g, --tag-match=PROPERTY
           Trigger events for devices with a matching tag. This option can
           be specified multiple times.

       -y, --sysname-match=PATH
           Trigger events for devices with a matching sys device path. This
           option can be specified multiple times and supports shell style
           pattern matching.

       --name-match=NAME
           Trigger events for devices with a matching device path. This
           option can be specified multiple times.

       -b, --parent-match=SYSPATH
           Trigger events for all children of a given device.

       -h, --help
           Print help text.

       In addition, optional positional arguments can be used to specify
       device names or sys paths. They must start with /dev or /sys
       respectively.

   udevadm settle [options]
       Watches the udev event queue, and exits if all current events are
       handled.

       -t, --timeout=SECONDS
           Maximum number of seconds to wait for the event queue to become
           empty. The default value is 120 seconds. A value of 0 will check
           if the queue is empty and always return immediately.

       -E, --exit-if-exists=FILE
           Stop waiting if file exists.

       -h, --help
           Print help text.

   udevadm control command
       Modify the internal state of the running udev daemon.

       -e, --exit
           Signal and wait for systemd-udevd to exit.

       -l, --log-priority=value
           Set the internal log level of systemd-udevd. Valid values are the
           numerical syslog priorities or their textual representations:
           emerg, alert, crit, err, warning, notice, info, and debug.

       -s, --stop-exec-queue
           Signal systemd-udevd to stop executing new events. Incoming
           events will be queued.

       -S, --start-exec-queue
           Signal systemd-udevd to enable the execution of events.

       -R, --reload
           Signal systemd-udevd to reload the rules files and other
           databases like the kernel module index. Reloading rules and
           databases does not apply any changes to already existing devices;
           the new configuration will only be applied to new events.

       -p, --property=KEY=value
           Set a global property for all events.

       -m, --children-max=value
           Set the maximum number of events, systemd-udevd will handle at
           the same time.

       --timeout=seconds
           The maximum number of seconds to wait for a reply from
           systemd-udevd.

       -h, --help
           Print help text.

   udevadm monitor [options]
       Listens to the kernel uevents and events sent out by a udev rule and
       prints the devpath of the event to the console. It can be used to
       analyze the event timing, by comparing the timestamps of the kernel
       uevent and the udev event.

       -k, --kernel
           Print the kernel uevents.

       -u, --udev
           Print the udev event after the rule processing.

       -p, --property
           Also print the properties of the event.

       -s, --subsystem-match=string[/string]
           Filter kernel uevents and udev events by subsystem[/devtype].
           Only events with a matching subsystem value will pass.

       -t, --tag-match=string
           Filter udev events by tag. Only udev events with a given tag
           attached will pass.

       -h, --help
           Print help text.

   udevadm test [options] [devpath]
       Simulate a udev event run for the given device, and print debug
       output.

       -a, --action=string
           The action string.

       -N, --resolve-names=early|late|never
           Specify when udevadm should resolve names of users and groups.
           When set to early (the default), names will be resolved when the
           rules are parsed. When set to late, names will be resolved for
           every event. When set to never, names will never be resolved and
           all devices will be owned by root.

       -h, --help
           Print help text.

   udevadm test-builtin [options] [command] [devpath]
       Run a built-in command COMMAND for device DEVPATH, and print debug
       output.

       -h, --help
           Print help text.



[3][Upgrade Syscall `passwd` hashing function]
The password hash used by the Linux Kernel for user passwords in no longer adequately secure and should be switched with a cryptographic hashing function like `bcrypt` that is better suited for secure password hashing.

https://github.com/coreos/bcrypt-tool

			"The security of a password depends upon the strength of the
			encryption algorithm and the size of the key space. The legacy UNIX
			System encryption method is based on the NBS DES algorithm. More
			recent methods are now recommended (see ENCRYPT_METHOD). The size of
			the key space depends upon the randomness of the password which is
			selected.

			Compromises in password security normally result from careless
			password selection or handling. For this reason, you should not
			select a password which appears in a dictionary or which must be
			written down. The password should also not be a proper name, your
			license number, birth date, or street address. Any of these may be
			used as guesses to violate system security.

			You can find advice on how to choose a strong password on
			http://en.wikipedia.org/wiki/Password_strength" - [linux manual: passwd](http://man7.org/linux/man-pages/man1/passwd.1.html)

Additional modification to the PAM module to support two factor authentication through TOTP, HOTP, OTP, decrypting a secret message or signing with a registered key would all provide significant increased security.

https://github.com/twitchyliquid64/pamtls

"passwd uses PAM to authenticate users and to change their passwords."

[PAM Resources]
http://man7.org/linux/man-pages/man5/pam.conf.5.html
http://man7.org/linux/man-pages/man3/pam_acct_mgmt.3.html
http://man7.org/linux/man-pages/man3/pam_chauthtok.3.html
http://man7.org/linux/man-pages/man3/pam_open_session.3.html
http://man7.org/linux/man-pages/man3/pam_sm_open_session.3.html
http://man7.org/linux/man-pages/man3/pam_authenticate.3.html
http://man7.org/linux/man-pages/man3/pam_sm_authenticate.3.html
http://man7.org/linux/man-pages/man3/pam_set_data.3.html
http://man7.org/linux/man-pages/man3/pam_get_data.3.html
http://man7.org/linux/man-pages/man8/pam_filter.8.html
http://man7.org/linux/man-pages/man3/pam_get_user.3.html
http://man7.org/linux/man-pages/man3/pam_get_item.3.html
http://man7.org/linux/man-pages/man8/PAM.8.html [Important]


[4][Modify keyrings Linux kernel module to conform to scramble suit identity protocol]
http://man7.org/linux/man-pages/man1/login.1.html
http://man7.org/linux/man-pages/man7/environ.7.html
http://man7.org/linux/man-pages/man7/locale.7.html
       /etc/security/pam_env.conf
           Default configuration file
       /etc/environment
           Default environment file
       $HOME/.pam_environment
           User specific environment file
       /etc/pam.conf
           the configuration file
       /etc/pam.d
           the Linux-PAM configuration directory. Generally, if this
           directory is present, the /etc/pam.conf file is ignored.
       /usr/lib/pam.d
           the Linux-PAM vendor configuration directory. Files in /etc/pam.d
           override files with the same name in this directory.

       pam_unix - Module for traditional password authentication
       pam_unix.so
  unix_chkpwd is a helper program for the pam_unix module that verifies
       the password of the current user.
http://man7.org/linux/man-pages/man7/keyrings.7.html
