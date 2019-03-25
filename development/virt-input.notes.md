# Useful Tools For VIRTs things
===============================================================================
## URI / PATH

**go-pathmatch** cool concept because its a router that works on
  paths in general. but not sure it uses radix. it would need
  to use radix to be ideal.

_special URL types!_ [we should come up with one for multiverse, and vms]

**go-dataurl** a library provides work with data URLs defined by
  RFC 2397
  
  data:,Hello%20World!
  data:text/html;base64,xx...==
  data:text/plain;charset=utf-8,Thisisatest
  data:;charset=utf-8,Thisisadifftest
  data:text/plain,Thisisatest
  dataLtext.plain;charset-utf-8,sdfsdf

**go-hashuri**
  hash://sha256/{HAHS}
  hash://scramble/hash?action=thing
    maybe for scramble we can add like tree data 


## Text Encoding & Formats & Validations
go-ascii= tools for working with ascii cahracters
go-whitespace, to work with the 26 UNICODE whitespace cahracters, and 7 mandatory breaks
go-utf8s
go-numeric - utf8 numeric library 





#
# Virtual Machines VirtContainers
===============================================================================
 * Hypervisors
   * QEMU / NEMU
   * Kapsule / NoVM






#
# Virtual Block Device (and maybe FIlesystem?)
===============================================================================

  * Support snappy or zstandard

# Virtual SOUND-over-PCI card
===============================================================================
  kdevb0x/netsound - one option, may work, but think rather lower levels like
  either `pulseaudio` or `alsa`

**PCM**

**PCI Audio?**

**'alsa' or 'pulseaudio'**


# Virtual X-over-virtual PCI card
===============================================================================

  * Support snappy or zstandard

  * Working with video codecs, compression formats, colors, and pixel format:
    reiver/go-fourcc - four character code 4CC to parse video data

  * Framebuffer
    * reiver/go-fbdev
	framebuffer for X11
	framebuffer can be accessed as MEMORY (woot!)
	framebuffer can be accessed as FILE (woot)
	framebuffer can be accessed as VIRTUAL CONSOLE! (you know CTL+ALT+(F1-F9) )
      
  * **v4l2** 
    * charleye/v4l2-go - not the best, that is refiner but this has the best
      API

    * Quickcam uses v4l2, and interacts with it via mmaps, useful for viewing. 
      when experimenting with streaming on twitch, and with rtmp in general

      ___________________________________________________________________________
      [__we discovered that v4dl is what we used to capture video from webcams,__]
      [__digital cameras, and even for streaming our screen, as either a window__]
      [__or entire scren__]

    * GETTING IMPROTANT DATA FROM v4l2:
      This is awesome, can get all the card details, then check the cabiltiies
      like capabiolity streaming. So this at the very least may be used to
      interact with our virtual PCI card.


# Virtual Keyboard Input
===============================================================================
 * Encryption key derived from session


 * After inplmenting basic uinput stuff
   * implement a very basic TypeString("then you can type stuff <enter>")

   * Itd also be nice to do something chaining:
	* PressKey and maybe a spearate PressKeys:
            * Would be a fix amount of time, pretty normal amount but still
              randomized for avoiding signatures. 
        * HoldKey and HoldKeys: 
	  * Specify amount of time to Hold, if nothing specified then, maybe
            reasonable amount of time.
        * TypeString("then a string to type")
        * Return().Enter().Space().Backspace()... and so on. All chainablex
            
     Keyboard.PresKey(ALT+TAB).PressKey)

     **READING CONSOLE**
     WE would not read the screen via x or whatever, or use event. The way
     you get the response is by reading the console. 
     There is an example of this in "hasicorp/packer/blob/master/provisoner/shell/unix_reader.go": func scanUnixLine(data, atEOF bool) 
  


   * Mouse.MoveTo(Position{x: 90, y: 220}).RightClick(Hold)
     * An easy way to select from the right click menu without having to
       specify X/Y would be great. 

     * In fact the BEST way to handle mouse, would be to scan for buttons, menu
       items, FILE.. menu, taskbar, system tray, etc. All tracked so the X/Y
       is known by the softwawre and we don't have to know all that information.
   
       * GETTING IMPROTANT DATA FROM v4l2:
       github.com/reiver/go-v4l2

 
       * GETTING IMPORTANT WINDOW DATA FROM WAYLAND:
       We just need to know what exists on the screen. and if we are tracking it       then we can ask the program what is available to click on. This would
       be quite great. Just need reliable way to scan the UI OR even better
       get this data DIRECTLY FROM WAYLAND or XORG. 
      

