##
##  ScrambleShell: Screen Relaying
##
=================================================================
In the earliest prototypes of Multiverse OS, `xpra` was used to
relay the screens. This worked very well, but over time other
solutions for this were found, such as `looking glass` which is
designed to relay KVM screens at a low level. This is designed
around Windows but the concept could be applied to linux possibly
using something like Vulcan to achieve best results.

Things like VNC and Spice are not low level enough to provide
adequate speed.

-----

Instead of using SSH or TCP, we would likely get much better 
preformance by using a shared memory or FIFO type custom device
to provide much faster connection, or at the very least a Unix 
socket.

=================================================================
## Working Solutions 
-----------------------------------------------------------------
The current Solution that is being used in Multiverse OS during
development is **xpra**. Ideally it needs to be reworked, 
removing the python overlay and replace it with a Go or Rust 
solution.

Eventually rewriting and refining the lower level components to
get even better speed and improved identity segregation.
  
#### Other Options
  * **Looking Glass** though relies on directx type screen sharing
  the concept is good, it is a low level 


=================================================================
## Known Remote Desktop Clients
-----------------------------------------------------------------
A list of known remote desktop clients which may be useful for
a variety of reasons from UI, unique features, code or other 
reasons that may not be immediately obvious: 

* **vino/vinagre** I don't knowe why its installed as vino but
the program is viagre. I suspect this is a *spice* remote desktop
client.
















=================================================================
## 
-----------------------------------------------------------------





=================================================================
## 
-----------------------------------------------------------------

