##
##  ScrambleShell: Shell Extensions Manager
=================================================================
[gnome-shell-extensions](https://github.com/GNOME/gnome-shell-extensions)
GNOME Shell Extensions is a collection of extensions providing 
additional and optional functionality to GNOME Shell.

**Since GNOME Shell is not API stable, extensions work only**
**against a very specific version of the shell, usually the**
**same as this package (see "configure --version").**

Gnome Shell **DOES NOT** provide a method to manage extensions, 
its either hidden, deep inside the gnome tweak tool, or deep
inside `gnome-software`. This is among several other MAJOR 
problems relating to both security and functionality in the 
Gnome Shell Extension system.


  **PROBLEMS with Gnome Shell Extensions**
    [*] Gnome Shell does not provide a easy to use, or consistent
    way to manage Gnome Shell Extensions.
    
    [*] Gnome Shell does not verify versions will work before 
    installing.
    
    [*] Gnome Shell does not require extensions to register what
    functional requirements (networking, writing files, etc) the
    gnome shell extensions are using
    
    
    
  **SECURITY improvements for Gnome Shell Extensions**
    [*] Gnome shell extensions should be run inside of Multiverse
    OS clear containers and provided with EXACTLY the explicitly
    requested.
        [*] FORBID INTERNET ACCESS BY DEFAULT
        
				[*] Using a virtual file system (VFS), all extensions MUST
				be contained inside a fake and VERY limited VFS, only
				providing access to only the EXACT files granted permission
				to modify.
				
    [*] Upon install requested functionality should be presented
    to the user (think APKs in android) to verify if the program
    should ahve access to those features. 
   
    [*] Lint, determine if there is obvious malwave _(Not as_
    _important or may not be neccessary at all)_

    [*] DO NOT USE JS DIRECTLY, using either Go language's
    GopherJS or Rust's Servo JS rendering, implement extensions
    in Go, or Rust.
    

    
  **FUNCTIONAL improvements to Gnome Shell Extensions**
    [*] Provide a easy to access, easy to use, simple Gnome
    extension manager, with the ability to install from a list of
    community approved extensions.
				[*] Add INVENTORY system for Gnome Shell Extensions
				(include __Gnome Terminal__ functionality, like Gnome
				Terminal Prompt)

	  [*] Refuse to install versions that do not match.

    [IF] The Gnome Shell Extension Management proves to work well
    implement all custom features to Gnome for Multiverse OS
    through the improved extension system.

=================================================================
## Related Software 
----------------------------------------------------------------- 
[gnome-shell-sass](https://github.com/GNOME/gnome-shell-sass)
GNOME Shell Sass is a project intended to allow the sharing of
the theme sources in sass between gnome-shell and other projects
like gnome-shell-extensions.


=================================================================
## Gnome Extensions
-----------------------------------------------------------------
[toggle_term](https://github.com/edipox/toggle_term)
A simple script to transform your boring terminal into
A DROP DOWN boring terminal 
_New and nice drop down terminal, rest in `gnome-software`_ 
_fucking suck_

[remove-accessibility](https://github.com/tomasz-oponowicz/remove-accessibility)




