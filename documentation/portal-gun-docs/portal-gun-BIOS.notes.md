##
##  Multiverse OS: BIOS, Bootloader, and Initramfs
===========================================================


=====
https://github.com/cloudius-systems/osv/wiki/OSv-early-boot-(MBR)

What happens after the computer is turned on? How do you
write the very first lines of code the CPU would execute
right after the computer starts?

Here, again, OSv is supplying a nice, self contained answer.
Let's limit our answer to the `x86_64` architecture, even though OSv
supports ARM as well.

What happens right after an CPU is started? The CPU instruction pointer is
initialized to its _reset vector_. The x86 CPU sets their `CS:EIP` address
to the fixed address `0xF000:0xFFF0`, this address contains some
code leading you eventually to the BIOS.
See [coreboot](http://www.coreboot.org/Coreboot_v3#How_coreboot_starts_after_Reset)'s
documentation for additional information.

Let's see that in action

    $ # Start paused (-S) QEMU with debugger (-s)
    $ qemu-system-x86_64 -s -S -nographic
    $ # on another tab
    $ gdb
    ...
    (gdb) set architecture i8086 
    warning: A handler for the OS ABI "GNU/Linux" is not built into this configuration
    of GDB.  Attempting to continue with the default i8086 settings.

    The target architecture is assumed to be i8086
    (gdb) target remote localhost:1234
    Remote debugging using localhost:1234
    0x0000fff0 in ?? ()
    (gdb) info registers eip cs
    eip            0xfff0	0xfff0
    cs             0xf000	61440
    (gdb) 

As we start, the CPU start on `0xf000:0xfff0`. To see what the CPU is about to execute,
we'll have to translate the segmented address `0xf000:0xfff0` to a
linear address. Since we start in real mode, we simply have to multiply the segment
address by `0x10`, and add it to the IP, see [Wikipedia](http://wiki.osdev.org/Segmentation#Real_mode)
for additional details. Let's print the first instruction the CPU which should execute
at `0x10*0xf000+0xfff0=0xffff0`:

    (gdb) x/i 0xffff0
    0xffff0:	ljmp   $0xf000,$0xe05b

It's probably a jump to the BIOS code at `0xf000:0xe05b`. Let's step to the next instruction

    (gdb) si
    0x0000e05b in ?? ()
    (gdb) i r eip cs
    eip            0xe05b	0xe05b
    cs             0xf000	61440

Indeed, we jumped to `0xf000:0xe05b` as expected.

The BIOS then loads the MBR from the disk at address `0x7c00`, and executes
it.

What's in the MBR? In order to understand that, let's check what does
the bare OSv image contains. Looking at `build.mk` we can find the files
generating the basic OSv image, `loader.img`.
Here is a simplified version of the instruction that creates it from `build.mk`:

    loader.img: boot.bin lzloader.elf
        # first block, ie, 512 bytes, are simply boot.bin
        dd if=boot.bin of=$@ > /dev/null 2>&1, DD $@ boot.bin
        # Then, after 128 blocks of 512 bytes, ie, 64K, the lzloader.elf
        dd if=lzloader.elf of=$@ conv=notrunc seek=128 > /dev/null 2>&1
        # set number of blocks boot16.S fetches to lzloader.elf's size
    	$(src)/scripts/imgedit.py setsize $@ $(image-size), IMGEDIT $@)
        # write the command line parameters right after the MBR, after 512 bytes
    	$(src)/scripts/imgedit.py setargs $@ $(cmdline), IMGEDIT $@)

The first 512 bytes, the MBR, is `build/debug/boot.bin`. Then `imgedit.py` inserts
the command line parameter for OSv. Next we have zeros until the 64Kth byte.
After that, we put `lzloader.elf`, the loader. In ASCII art sketch:

    0......512.................65536=64k.....
    [boot.bin][cmdline]00000000[lzloader.elf]

Let's see MBR code in action.
Let's start the virtual machine, and tell it to stop at load time:

    $ ./scripts/run.py -d --wait

On a different terminal, let's connect with gdb, and break at `0x7c00`, the
address the BIOS loads the MBR to. Note that we're specifying commands that `gdb`
should run at startup with the `-ex` switch, so that when `gdb` is started,
it'll connect to QEMU, and the correct architecture is set:
    
    $ gdb -ex 'set architecture i8086' -ex 'target remote localhost:1234'
    ...
    (gdb) hbr *0x7c00
    Hardware assisted breakpoint 1 at 0x7c00
    (gdb) c
    Continuing.

    Breakpoint 1, 0x00007c00 in ?? ()

Now we finally started to run OSv code, the Master Boot Record, or the MBR.

Let's verify that this is the case.

Let's define a simple alias, that would display files
in `gdb`'s binary format:

    $ alias gdbdump='hexdump -e '\''"0x%04_ax: " 8/1 "0x%02x\t" "\n"'\'''

Now let's verify that the BIOS is indeed loading `boot.bin`:

    (gdb) x/32b $eip
    0x7c00:	0xea	0x5e	0x7c	0x00	0x00	0x00	0x00	0x00
    0x7c08:	0x00	0x00	0x00	0x00	0x00	0x00	0x00	0x00
    0x7c10:	0xdb	0x00	0x10	0x00	0x40	0x00	0x00	0x00
    0x7c18:	0x00	0x80	0x80	0x00	0x00	0x00	0x00	0x00
    $ gdbdump -n 32 build/release/boot.bin 
    0x0000:	0xea    0x5e    0x7c    0x00    0x00    0x00    0x00    0x00
    0x0008:	0x00    0x00    0x00    0x00    0x00    0x00    0x00    0x00
    0x0010:	0x00    0x10    0x10    0x00    0x40    0x00    0x00    0x00
    0x0018:	0x00    0x80    0x80    0x00    0x00    0x00    0x00    0x00

Looks like this is the case.

How is `boot.bin` generated?
The answer is, again, in `build.mk`. First `boot16.S` is compiled to `boot16.o`
with regular compilation, by the `%.o: %.S` rule. Then, `boot.bin`
is being linked from `boot16.o` and `boot16.ld` linker script.

What does `boot16.ld` do? At first it defines a memory section of the
available memory at boot time. The first 1MB.

    MEMORY { BOOTSECT : ORIGIN = 0, LENGTH = 0x10000 }

Then, it'll take all text sections from the input, and relocate them to
`0x7c00`, where the MBR would be loaded in to. It would verity that the text
section fits to the first megabyte of memory we defined previously. That
way, if `boot16.S` accidently surpass 1MB, the linker would complain.

    SECTIONS { .text 0x7c00 : { *(.text) } > BOOTSECT }

For example, if we'll add a megabyte of data at the end of `boot16.S`,
we'll get the following:

    $ echo .fill 0x10000, 1, 0 >> arch/x64/boot16.S
    $ make mode=debug
    ...
    make[1]: Entering directory `/home/elazar/dev/osv/build/debug'
      GEN gen/include/osv/version.h
      AS arch/x64/boot16.o
      LD boot.bin
    ld: address 0x17e00 of boot.bin section `.text' is not within region `BOOTSECT'

Finally, we'll instruct `ld` to output the raw assembly instructions,
without any ELF headers.

    OUTPUT_FORMAT(binary)

What does `boot16.S` do?

At first we can see some x86 bookkeeping. Setting up the
[A20 line](http://www.win.tue.nl/~aeb/linux/kbd/A20.html), and the segment
registers.

Then, it'll try to load the command line arguments from disk, using
[interrupt 13](http://wiki.osdev.org/ATA_in_x86_RealMode_%28BIOS%29#LBA_in_Extended_Mode)

    int1342_boot_struct:
    .byte 0x10 # size of packet (16 bytes)
    .byte 0 # should always be 0
    .short 0x3f   # fetch 0x3f sectors = 31.5k
    .short cmdline # fetch to address $cmdline
    .short 0 # fetch to segment 0
    .quad 1 # start at LBA 1.
    # That is, fetch the first 31.5k from the disk
    ...
    lea int1342_boot_struct, %si
    mov $0x42, %ah
    mov $0x80, %dl
    int $0x13

Indeed after the interrupt, we can see something in `cmdline=0x7e00`

    (gdb) hbr *0x7c81
    ...
    (gdb) x/32b 0x7e00
    0x7e00:	0x00	0x00	0x00	0x00	0x00	0x00	0x00	0x00
    0x7e08:	0x00	0x00	0x00	0x00	0x00	0x00	0x00	0x00
    0x7e10:	0x00	0x00	0x00	0x00	0x00	0x00	0x00	0x00
    0x7e18:	0x00	0x00	0x00	0x00	0x00	0x00	0x00	0x00
    (gdb) si
    0x00007c8c in ?? ()
    (gdb) x/32b 0x7e00
    0x7e00:	0x2f	0x75	0x73	0x72	0x2f	0x6d	0x67	0x6d
    0x7e08:	0x74	0x2f	0x68	0x74	0x74	0x70	0x73	0x65
    0x7e10:	0x72	0x76	0x65	0x72	0x2e	0x73	0x6f	0x26
    0x7e18:	0x6a	0x61	0x76	0x61	0x2e	0x73	0x6f	0x20

Those bytes are indeed the command line arguments given to OSv

    $ gdbdump -n 32 build/debug/cmdline
    0x0000: 0x2f	0x75	0x73	0x72	0x2f	0x6d	0x67	0x6d
    0x0008: 0x74	0x2f	0x68	0x74	0x74	0x70	0x73	0x65
    0x0010: 0x72	0x76	0x65	0x72	0x2e	0x73	0x6f	0x26
    0x0018: 0x6a	0x61	0x76	0x61	0x2e	0x73	0x6f	0x20

Let's move on. Now `boot16.S` would load OSv's loader from disk.

We have a problem here. On the one hand, we need to be in real mode
in order to use `int 0x13h` and access the disk with the BIOS. On the other
hand, we need to be in protected mode in order to access more than the first
1 MB of memory. What `boot16.S` does, is, switch to real mode, fetch
a few KB from the disk, move to protected mode, and copy them to memory, back
to real mode, rinse and repeat.

In order to do that, we have to have a GDT that supports both 16-bit and 32-bit
segments. Let's see how the GDT is configured:

    gdt:
    .short gdt_size - 1
    .short gdt
    .long 0
    #             
    #     base flag limit type  base  limit
    .quad 0x00 c    f     9b   000000 ffff # 32-bit code segment
    .quad 0x00 c    f     93   000000 ffff # 32-bit data segment
    .quad 0x00 0    0     9b   000000 ffff # 16-bit code segment
    .quad 0x00 0    0     93   000000 ffff # 16-bit data segment
    ...
    # set the gdt
    cli
    lgdtw gdt

The first GDT entry is the zero descriptor, then
two 32 bit flat selectors `limit = 0xfffff, base=0x0` whose flag
have the *size* and *granularity* bits on in `flag`. Next two
identical flat 16 bit segments, so that we'll be able to jump back
to real mode. See [GDT](http://wiki.osdev.org/GDT) section in OSDev
for addition details.

Now let's see the snippets that takes us back and forth from protected
mode to real mode:

    # set protected mode bit in cr0
    mov $0x11, %ax
    lmsw %ax
    # move to 32 bit code segment (0x8 = first) - protected mode
    # ljmp to flush prefetch queue http://goo.gl/JBOnZ5
    ljmp $8, $1f
    ...
    # move to 16 bit code segment (0x18 = third)
    # then set real mode in cr0
    ljmpw $18, $1f
    1:
    .code16
    # clear protected mode bit
    mov $0x10, %eax
    mov %eax, %cr0
    ljmpw $0, $1f

Finally let's see the process of moving memory from the disk. First `read_disk`
is used to read `0x8000` bytes from the disk to `tmp`:

    read_disk:
    lea int1342_struct, %si
    mov $0x42, %ah
    mov $0x80, %dl
    int $0x13

Then in protected mode we copy them to `xfer`, and increment the value
at `xfer` by `0x8000`:

    mov $0x10, %ax
    mov %eax, %ds
    mov %eax, %es
    mov $tmp, %esi
    mov xfer, %edi
    mov $0x8000, %ecx
    rep movsb
    mov %edi, xfer

Now back in real mode, we read more from the disk, unless we already read
`count32` bytes. The `count32` memory location is being set to the loader's
size by OSv build process.

    xor %ax, %ax
    mov %ax, %ds
    mov %ax, %es
    sti
    addl $(0x8000 / 0x200), lba
    decw count32
    jnz read_disk

Finally, the loader would save the memory map to `mb_mmap_addr` with
[int 15h e820](http://www.uruk.org/orig-grub/mem64mb.html).

The very last step is, jump into the `lzloader.elf` code.
We go back to protected mode, jump to a predefined addresses
that would decompress the loader code, and another call to the
decompressed loader code.

From now on, we're salvated from the x86 assembly land, and most
of the code would be in C++.

The first part of the journey is done, from reset to the OS loader.

