# Expressive Go 

```
class Account < ActiveRecord::Base
  validates :subdomain, :name, :email_address, :password, presence: true
  validates :subdomain, uniqueness: true
  validates :terms_of_service, acceptance: true, on: :create
  validates :password, :email_address, confirmation: true, on: :create
end
```
______
Beginning to brainstorm about the concept of taking the Go source code and
cutting out packages I do not use, adding in packages that are missing (uart,
jtag, spi, etc), linux kernel tools, MVC, CLI, Shell. (None of those are decided
just brainstorming) 

The idea of modifying the Go source has been an interest for a while now, and
ideally I should be interacting with it so I can contribute to the project to
build reputation with the community. There is definitely work to be done. 

But the idea is to make it more Ruby like now, adding ? ! = and other symbols to
function names. Creating standard methods, like Blank? Nil? etc.. 

Making struct called class, Len Length, support aliasing by default, etc. 

THe result would ideally be minimalistic, linux focused, expressive language
with all the core functionality of Go. 

Interators like Each Times would be great over the current range. Builtins for
Blank == "", Zero == "" constants. 

**Better support for calling packages from local folders over remote libraries**
with a Bundle like system for downloading gems to the local folder and locoking
the current git commit. 

**A standard folder structure that is not system wide but rather project wide.
All information automatically stored in a .GoFile style file that would have all
the Makefile type information. 

Builtin support for embeding binary data into the compiler 

Builtin support for runtime interpreting 

Builtin ssupport for embedding executables like Ruby executable for scripting
for example or qcow-cli for making qcows without implementing ourselves. 

Print functions by adding files to it, can print to multiple files at once, like
os.Stdout and a log file builtin. Support for regex based coloring and
stripping before printing to output. 

BUILTIN SUPPORT FOOR VALIDATION OF DATA! So int type can define limitations
directly on the variable. If the variable ever tries to defy it, it refuses and
returns and error.


______ 

What I'm really curious about, is if I make these changes above, make the
language WAY easier to use, easier to read, easier to understand. How much
preformance lost will we experience for this cost. If its not substnational, I
would considering finding a way to build it as a patch on Go so we can stay up
to date with underlying core changes. 
