##
##  ScrambleShell: Security 
##
=================================================================
Scramble shell provides many small features, tools and default 
settings that add up to a more secure desktop environment.

**Features**
	* Auto-remove clipboard contents after __x__ seconds


=================================================================
## Deterministic Complex Password Generation 
-----------------------------------------------------------------
The concept is essentially to use a deterministic but secret 
generation technique, randomized and encrypted with your master
key that can be used (Domain being used, username, and the users
simple and easy to remember password) is used to then generate
a very complex 32 character password. 

This would allow users to not need to remember very complex
passwords and to deterministically regenerate/recover the
passwords from their simple password. 

*I have seen this implemented but the method of generating the*
*passwords is always static between all users, but it could be*
*vastly improved by addig in determistic and unique randomness*
*so that bruteforcing would not be possible*
	[Example/Reference Software]
		https://extensions.gnome.org/extension/825/password-calculator/


=================================================================
## Clipboard Manager
-----------------------------------------------------------------
A full featured clipboard manager is requried for Multiverse OS
largely because Multiverse OS is an operating system that ties
several networked machines in a cluster together to function as
a single computer and for this to work in a fluid way that is 
natural to users.

Basic features would include allowing one-way or bi-directional
copying of clibboard contents:
		Application VM <-> Controller VM and vice-versa 

The core functionality would be acheived using a pattern matching
system that would be configured using Ruby, that would encrypt
contents of clipboard and decrypt it as needed using a special
session clipboard key that is ephemeral and generated per
session. 

Since our passwords will be auto-generated and 32 characters
it should be relatively easy to reliably detect passwords in
the clibboard and that will allow us to deal with passwords in
a safe manner, like immeidate removal after the first paste. 

	[Example/Reference Software]
	  **CopyQ** is a great clipboard manager with editing and 
	  scripting features and its feature list should be reviewed
	  because it provides many great features. 
	  
	  Such as automatically detecting *images*, *links* and then
	  providing tabs to review and use these items later which
	  is an outstanding idea.
	  
	  Allowing scripting allows the user to prevent caching
	  or saving history of sensntive data that was once in 
	  the clipboard which many clipboard managers fail to account
	  for. Considering applications like `pass` that encourage the
	  use of 32+ character passwords by loading the password into
	  the clipboard buffer for pasting into the password field
	  this is very important.
	  
	  **Clipman** enables you to store and recall X selections, as
	  well as GTK+ clipboard content. You can also define actions
	  to be triggered by selecting some text pattern.

**Features**
	* Auto-remove clipboard contents after __x__ seconds
  * Pattern based scripting over clipboard contents
  	[>][Example] Detect passwords and remove after first paste and 
  	control flow between VMs
  
=================================================================
## Automatic Pax (grSEC) 
-----------------------------------------------------------------



