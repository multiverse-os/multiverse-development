# WOrking with Grub

## Grub Env file
just a file that like holds a speicifc size for grub to work with by defualt is 1024 kb or b or something



_______________________________________________________________
# Multiverse OS, the hidden OS

The original design of Multiverse was that it was hidden and would
require special sequence or action, or key inserted (that would
be a kill switch/dead-man-switch when pulled out).

This can be done using grub. It could be super hidden and any mistep would boot a small vanilla qcow debian



_______________________________________________________________


## Grub Modules

Using customized GRUB and GRUB modules, we can implement a super low level FALLBACK Controller VM. That generates a list of available CONTROLLER VMs and displays them in a grub style menu. Choosing one would boot to that VM. This grub ONLY  CONTROLLER VM could be booted into, if a CONTROLLER VM fails to complete the boot processes after X amount of time.


## CPIO - Create files/archives that can be accessible by grub


The following command takes all files in directory and puts them into a CPIO
`ls | cpio -ov > ~/output-archive.cpio`


## Examples of functionality of grub modules
Legend for below [Z] is potentially useful for zerg, modifying boot record. [MV] is potentially use for Multiverse [MV:C] is Multiverse CONTROLLER VM, [MV:A] Multiverse APPLICATION VM, etc.



  * [MV:C] CPUID - this module tests various CPU features *(good potentailly for testing custom CPUs built within VM using selected disabled/enabled CPU features and CPU passthrough)*

  
  * acpi - power related fucnityonality

  * [Z] memrw - read/write to physical memory

  * [Z] cmosdump - raw CMOS data

  * cryptodisk - needed to mount cryptodisk

  * [Z][MV] http - GRUB can boot that resides on NFS (!)

  * [MV] hwmatch - whitelist/blacklist hardware 

  * [MV][Z] jpeg/png/etc -can display background image in GRUB
    Zerg could use this by taking existing background image which may be large and hiding data in the existing PNG/JPEG and replacing it with the new version.

  * [Z][MV] net - provides networking modules needed to access the net (AT GRUB!)

  * [Z] ntfs 



   * regex - 

   * RomFS - rom filesystem simple file system lacking many features intended for burning files onto eeprom

## PCI - could be useful
  * [MV] setpci - configure PCI devices

  * [MV] pcidump - get all pci devices

  * [MV] probe - get device info

  * [MV] pci - use PCI devices, like video card? could be necesssary for seeing the password prompt!

[video -trying to get visuals on password entry]

  * allvideo - provides additonal graphics suipport

## LIVE CDs (epehermerals)

  * iso9660 and biosdisk modules provies ability to boot off "LiveDisks"

## Controlling computer FROM grub (for super low level provisioning needed by portal-gun) [Z][MV]

  * serial - this module provides serial devices (like then we can VirtIO in and PXE, use this to install OS, then combine with sendkey to control shit!) [Z][MV]

  * sendkey - send a key command to be used to emualte key presses to grub


[mega weird maybe way to communicate][via speakers, bios speaks, but fake like VM doing it and transfering data over sound] 

  * spkmodem - generic speaker driver
  * spkmodem-recv

    

## Chain Loading and PXE (Preboot execution envrionemnt) [Z][MV]


  * [MV][Z] pxe - preboot execution environemnt used to boot an operating system idnepdently of local storage units [PWN][!]
    [MV] also how we would do a controll flaback that could boot to other controllers after failure.


  * [Z][MV] chain, pxechain - Chain loading 


   
## Zerg - aka cryptozergology, gonna fuck shit up beyond belief


  * use this to boot to OWN kernel, with patches, premake these for Ubuntu, deibn, etc.


_______________
## GRUB LIST FILES




/boot/grub/x86_64-efi$ cat terminal.lst 
iat_keyboard: at_keyboard
iserial: serial
iserial_*: serial
oaudio: morse
ocbmemc: cbmemc
ogfxterm: gfxterm
oserial: serial
oserial_*: serial
ospkmodem: spkmodem


 fs.lst 
affs
afs
bfs
btrfs
cbfs
cpio
cpio_be
exfat
ext2
fat
hfs
hfsplus
iso9660
jfs
minix
minix2
minix2_be
minix3
minix3_be
minix_be
newc
nilfs2
ntfs
odc
procfs
reiserfs
romfs
sfs
squash4
tar
udf
ufs1
ufs1_be
ufs2
xfs
zfs




cat video.lst 
efi_gop
efi_uga
video_bochs
video_cirrus


cat partmap.lst 
part_acorn
part_amiga
part_apple
part_bsd
part_dfly
part_dvh
part_gpt
part_msdos
part_plan
part_sun
part_sunpc


:/boot/grub/x86_64-efi$ cat moddep.lst 
squash4: xzio gzio lzopio fshelp
search_fs_uuid:
lssal:
legacycfg: gcry_md5 crypto password normal
date: datetime normal
bfs: fshelp
uhci: usb
multiboot2: boot video net acpi relocator mmap
gcry_twofish: crypto
cpio_be: archelp
priority_queue:
gcry_rijndael: crypto
echo: extcmd
cpio: archelp
xzio: crypto
part_sun:
hfspluscomp: gzio hfsplus
gcry_sha512: crypto
gcry_cast5: crypto
efi_gop: video video_fb
ctz_test: functional_test
boot:
setjmp_test: setjmp functional_test
odc: archelp
ls: extcmd normal
gzio:
cbmemc: cbtable terminfo normal
video:
test_blockarg: extcmd normal
linuxefi: boot
gfxterm: video font
cbtable:
mul_test: functional_test
gcry_tiger: crypto
gcry_serpent: crypto
aout:
macbless: disk
gcry_blowfish: crypto
appleldr: boot
trig:
extcmd:
at_keyboard: boot keylayouts
videoinfo: video
testspeed: extcmd normal
minix:
keylayouts:
xnu_uuid: gcry_md5
usbtest: usb
usbms: usb scsi
reboot:
morse:
help: extcmd normal
part_msdos:
http: net
gcry_rsa: mpi verify
cbtime: cbtable
blocklist:
probe: extcmd
pbkdf2: crypto
lsefisystab:
gcry_rfc2268: crypto
ufs1_be:
nativedisk:
gcry_camellia: crypto
fat: fshelp
exfctest: functional_test
parttool: normal
lzopio: crypto
linux: boot video relocator mmap
gcry_md4: crypto
zfsinfo: zfs
usb_keyboard: keylayouts usb
gcry_md5: crypto
fshelp:
ehci: boot usb cs5536
datetime:
bitmap_scale: bitmap
ata: scsi
usbserial_common: usb serial
syslinuxcfg: extcmd normal
net: priority_queue boot datetime bufio
gcry_des: crypto
div_test: div functional_test
time:
reiserfs: fshelp
dm_nv: diskfilter
datehook: datetime normal
mdraid09_be: diskfilter
cmp_test: functional_test
backtrace:
ahci: boot ata
kernel:
video_cirrus: video video_fb
part_plan:
gcry_seed: crypto
minix_be:
crypto:
video_colors:
test:
terminal:
part_dvh:
lsacpi: extcmd acpi
jpeg: bufio bitmap
bsd: boot video aout extcmd gcry_md5 crypto cpuid elf relocator serial mmap
memdisk:
gfxmenu: video gfxterm trig bitmap_scale video_colors bitmap normal font
cmp:
acpi: extcmd mmap
xfs: fshelp
elf:
div:
cpuid: extcmd
affs: fshelp
usb:
videotest: video gfxmenu font
tr: extcmd
testload:
relocator: mmap
play:
gfxterm_menu: video_fb functional_test procfs normal font
cbfs: archelp
adler32: crypto
progress: normal
password: crypto normal
part_sunpc:
video_fb:
tftp: priority_queue net
sleep: extcmd normal
serial: extcmd terminfo
search_fs_file:
gcry_sha256: crypto
gcry_rmd160: crypto
exfat: fshelp
search: search_fs_uuid extcmd search_fs_file search_label
mdraid09: diskfilter
chain: boot net efinet
mpi: crypto
memrw: extcmd
cs5536:
password_pbkdf2: gcry_sha512 pbkdf2 crypto normal
mdraid1x: diskfilter
linux16: boot video relocator mmap
gcry_crc: crypto
configfile: normal
zfscrypt: gcry_rijndael extcmd pbkdf2 crypto zfs gcry_sha1
signature_test: functional_test procfs
raid5rec: diskfilter
pcidump: extcmd
gcry_arcfour: crypto
part_dfly:
minix2_be:
gettext:
pbkdf2_test: pbkdf2 gcry_sha1 functional_test
hello: extcmd
usbserial_pl2303: usbserial_common usb serial
hashsum: extcmd crypto normal
xnu_uuid_test: functional_test
regexp: extcmd normal
part_gpt:
ohci: boot usb cs5536
gptsync: disk
zfs: gzio
part_apple:
hdparm: extcmd
bufio:
btrfs: gzio lzopio
bitmap:
true:
terminfo: extcmd
romfs: fshelp
ntfscomp: ntfs
lsefi:
hfs: fshelp
gcry_dsa: mpi verify
cmdline_cat_test: video_fb functional_test procfs normal font
ufs1:
offsetio:
legacy_password_test: legacycfg functional_test
setjmp:
ufs2:
nilfs2: fshelp
lsmmap: mmap
gcry_sha1: crypto
mmap:
tar: archelp
png: bufio bitmap
lspci: extcmd
hfsplus: fshelp
cbls: cbtable
tga: bufio bitmap
random: hexdump
minix2:
setpci: extcmd
scsi:
pata: ata
minix3:
lvm: diskfilter
functional_test: video extcmd video_fb btrfs
eval: normal
iso9660: fshelp
crc64: crypto
udf: fshelp
search_label:
raid6rec: diskfilter
msdospart: parttool disk
archelp:
procfs: archelp
minix3_be:
halt: acpi
xnu: boot video extcmd bitmap_scale relocator bitmap mmap random macho
read:
multiboot: boot video net relocator mmap
keystatus: extcmd
cryptodisk: extcmd crypto procfs
shift_test: functional_test
normal: boot extcmd crypto terminal gettext bufio
lsefimmap:
loadbios:
geli: gcry_sha512 pbkdf2 crypto gcry_sha256 cryptodisk
spkmodem: terminfo
gcry_idea: crypto
bswap_test: functional_test
video_bochs: video video_fb
verify: extcmd crypto mpi gcry_sha1
sfs: fshelp
part_amiga:
luks: pbkdf2 crypto cryptodisk
loopback: extcmd
jfs:
gfxterm_background: gfxterm video extcmd bitmap_scale video_colors bitmap
efifwsetup: reboot
usbserial_usbdebug: usbserial_common usb serial
part_acorn:
newc: archelp
macho:
iorw: extcmd
cat: extcmd
afs: fshelp
sleep_test: datetime functional_test
ldm: part_msdos part_gpt diskfilter
hexdump: extcmd
efinet: net
disk:
usbserial_ftdi: usbserial_common usb serial
minicmd:
loadenv: extcmd disk
gcry_whirlpool: crypto
fixvideo:
part_bsd: part_msdos
font: video bufio
ext2: fshelp
diskfilter:
videotest_checksum: video_fb functional_test font
file: extcmd elf offsetio macho
ntfs: fshelp
efi_uga: video video_fb
all_video: efi_gop efi_uga video_bochs video_cirrus









/boot/grub/x86_64-efi$ cat command.lst 
*acpi: acpi
*all_functional_test: functional_test
*background_image: gfxterm_background
*cat: cat
*cpuid: cpuid
*crc: hashsum
*cryptomount: cryptodisk
*echo: echo
*extract_syslinux_entries_configfile: syslinuxcfg
*extract_syslinux_entries_source: syslinuxcfg
*file: file
*functional_test: functional_test
*gettext: gettext
*hashsum: hashsum
*hdparm: hdparm
*hello: hello
*help: help
*hexdump: hexdump
*inb: iorw
*inl: iorw
*inw: iorw
*keystatus: keystatus
*kfreebsd: bsd
*knetbsd: bsd
*kopenbsd: bsd
*list_env: loadenv
*load_env: loadenv
*loopback: loopback
*ls: ls
*lsacpi: lsacpi
*lspci: lspci
*md5sum: hashsum
*menuentry: normal
*pcidump: pcidump
*probe: probe
*read_byte: memrw
*read_dword: memrw
*read_word: memrw
*regexp: regexp
*save_env: loadenv
*search: search
*serial: serial
*setpci: setpci
*sha1sum: hashsum
*sha256sum: hashsum
*sha512sum: hashsum
*sleep: sleep
*submenu: normal
*syslinux_configfile: syslinuxcfg
*syslinux_source: syslinuxcfg
*terminfo: terminfo
*test_blockarg: test_blockarg
*testspeed: testspeed
*tr: tr
*trust: verify
*verify_detached: verify
*xnu_splash: xnu
*zfskey: zfscrypt
.: configfile
[: test
appleloader: appleldr
authenticate: normal
background_color: gfxterm_background
backtrace: backtrace
badram: mmap
blocklist: blocklist
boot: boot
break: normal
cat: minicmd
cbmemc: cbmemc
chainloader: chain
clear: normal
cmp: cmp
configfile: configfile
continue: normal
coreboot_boottime: cbtime
cutmem: mmap
date: date
distrust: verify
dump: minicmd
eval: eval
exit: minicmd
export: normal
extract_entries_configfile: configfile
extract_entries_source: configfile
extract_legacy_entries_configfile: legacycfg
extract_legacy_entries_source: legacycfg
fakebios: loadbios
false: true
fix_video: fixvideo
fwsetup: efifwsetup
gptsync: gptsync
halt: halt
help: minicmd
hexdump_random: random
initrd16: linux16
initrd: linux
initrdefi: linuxefi
keymap: keylayouts
kfreebsd_loadenv: bsd
kfreebsd_module: bsd
kfreebsd_module_elf: bsd
knetbsd_module: bsd
knetbsd_module_elf: bsd
kopenbsd_ramdisk: bsd
legacy_check_password: legacycfg
legacy_configfile: legacycfg
legacy_initrd: legacycfg
legacy_initrd_nounzip: legacycfg
legacy_kernel: legacycfg
legacy_password: legacycfg
legacy_source: legacycfg
linux16: linux16
linux: linux
linuxefi: linuxefi
list_trusted: verify
loadbios: loadbios
loadfont: font
lscoreboot: cbls
lsefi: lsefi
lsefimmap: lsefimmap
lsefisystab: lsefisystab
lsfonts: font
lsmmap: lsmmap
lsmod: minicmd
lssal: lssal
macppcbless: macbless
mactelbless: macbless
module2: multiboot2
module: multiboot
multiboot2: multiboot2
multiboot: multiboot
nativedisk: nativedisk
net_add_addr: net
net_add_dns: net
net_add_route: net
net_bootp6: net
net_bootp: net
net_del_addr: net
net_del_dns: net
net_del_route: net
net_get_dhcp_option: net
net_ipv6_autoconf: net
net_ls_addr: net
net_ls_cards: net
net_ls_dns: net
net_ls_routes: net
net_nslookup: net
normal: normal
normal_exit: normal
outb: iorw
outl: iorw
outw: iorw
parttool: parttool
password: password
password_pbkdf2: password_pbkdf2
play: play
read: read
reboot: reboot
return: normal
rmmod: minicmd
search.file: search_fs_file
search.fs_label: search_label
search.fs_uuid: search_fs_uuid
setparams: normal
shift: normal
source: configfile
terminal_input: terminal
terminal_output: terminal
test: test
testload: testload
time: time
true: true
usb: usbtest
videoinfo: videoinfo
videotest: videotest
write_byte: memrw
write_dword: memrw
write_word: memrw
xnu_devprop_load: xnu
xnu_kernel64: xnu
xnu_kernel: xnu
xnu_kext: xnu
xnu_kextdir: xnu
xnu_mkext: xnu
xnu_ramdisk: xnu
xnu_resume: xnu
xnu_uuid: xnu_uuid
zfs-bootfs: zfsinfo
zfsinfo: zfsinfo



at crypto.lst 
RIJNDAEL: gcry_rijndael
RIJNDAEL192: gcry_rijndael
RIJNDAEL256: gcry_rijndael
AES128: gcry_rijndael
AES-128: gcry_rijndael
AES-192: gcry_rijndael
AES-256: gcry_rijndael
ADLER32: adler32
CRC64: crc64
ARCFOUR: gcry_arcfour
BLOWFISH: gcry_blowfish
CAMELLIA128: gcry_camellia
CAMELLIA192: gcry_camellia
CAMELLIA256: gcry_camellia
CAST5: gcry_cast5
CRC32: gcry_crc
CRC32RFC1510: gcry_crc
CRC24RFC2440: gcry_crc
DES: gcry_des
3DES: gcry_des
DSA: gcry_dsa
IDEA: gcry_idea
MD4: gcry_md4
MD5: gcry_md5
RFC2268_40: gcry_rfc2268
AES: gcry_rijndael
AES192: gcry_rijndael
AES256: gcry_rijndael
RIPEMD160: gcry_rmd160
RSA: gcry_rsa
SEED: gcry_seed
SERPENT128: gcry_serpent
SERPENT192: gcry_serpent
SERPENT256: gcry_serpent
SHA1: gcry_sha1
SHA224: gcry_sha256
SHA256: gcry_sha256
SHA512: gcry_sha512
SHA384: gcry_sha512
TIGER192: gcry_tiger
TIGER: gcry_tiger
TIGER2: gcry_tiger
TWOFISH: gcry_twofish
TWOFISH128: gcry_twofish
WHIRLPOOL: gcry_whirlpool


