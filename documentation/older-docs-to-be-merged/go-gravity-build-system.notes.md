##
##  Multiverse OS: Gravity Build System
===========================================================

Multiverse OS provides a build system, currently with the codename `gravity`. This tool is included to provide developers with a consistent, reliable and simplified way to compile all the components that make up Multiverse OS. The build system provides a simple way to compile components individually for independent use of sub-components or altogether for use with the operating system.
### Examples / Samples
https://github.com/nikogura/gomason - A tool for testing, building, signing, and publishing go binaries from a clean workspace. 


## GO CODE LINTING
[go-tools]
https://github.com/dominikh/go-tools
go get honnef.co/go/tools/cmd/megacheck 
megacheck runs staticcheck, gosimple and unused at once. Because it is able to reuse work, it will be faster than running each tool separately.
gosimple 	Detects code that could be rewritten in a simpler way.
staticcheck 	Detects a myriad of bugs and inefficiencies in your code.
unused 	Reports unused identifiers (types, functions, ...) in your code.
unused 	Reports unused identifiers (types, functions, ...) in your code.
[other]
https://github.com/haya14busa/reviewdog
https://github.com/mvdan/interfacer A linter that suggests interface types. In other words, it warns about the usage of types that are more specific than necessary.
https://github.com/THE108/enumlinter
https://github.com/99designs/scopecheck
https://github.com/nearmap/goreportcardlite To generate reports on the quality of go project based on several measures, several measures, including gofmt, go vet, go lint and gocyclo
### Visualize Crawl
[Network]
**Phylotree like graph**
https://github.com/datastorm-open/visNetwork
http://datastorm-open.github.io/visNetwork/
## WebKit GO Bindings 
https://github.com/abcum/webkit **for a desktop app**
https://github.com/abcum/webdriver - bloody fucking aamzing, PURE go, NO DEPS, webdriver adhering to seleniums standard! for driving a browser!
## Pure Go Git
https://github.com/src-d/go-git A highly extensible Git implementation in pure Go.
https://github.com/speedata/gogit Pure Go read access of a Git repository

## Diff
https://github.com/sergi/go-diff

## Markdown Notes
[live edit, with graphs]
https://github.com/adriamb/runes

### Mass change imports
**I have wasted SO much time on this shit, this alone would be a massive help**
https://github.com/eginez/go-import-change
### Two separate command-lines
* Command-line tool with flags and sub-commands
  [Potential Libraries]
  https://github.com/tucnak/climax

[bloated framework that may ahve useful examples](https://github.com/ukautz/clif#real-life-example)

  https://github.com/alecthomas/kingpin 
  i like that this support combing flags POSX style like
		-a + -b = -ab

  https://github.com/ukautz/clif

  ufave/cli
			$ greet help
			NAME:
					greet - fight the loneliness!

			USAGE:
					greet [global options] command [command options] [arguments...]

			VERSION:
					0.0.0

			COMMANDS:
					help, h  Shows a list of commands or help for one command

			GLOBAL OPTIONS
					--version Shows version information


  ---
https://github.com/alexeyco/simpletable
  [CLI Interactions/Prompts]
  https://github.com/deiwin/interact



* Command-line tool for shell/console



#### Draft Design
Below is a draft/sketch of all the features needed to satisfy the original `gravity` build system functional requirements.

(1) Create a Go tool that will execute the follow command after issuing a `bundle init` like command:

		// For now use this shorthand method to
		// set GOPATH to project directory to 
		// segregate project dependencies from
		// other projects
		export GOPATH=$(pwd)

		// Once the `gravity` build toold is
		// more development, one will run:
		gravity init
		// which mirrors bundler functionality 
		// when initializing projects


(2) Track all the projects managed, this way we can get the .env like functionality of Ruby without overloading `cd` which is kinda insane thing to do. This way we just track the directories, then when we enter into them we can jsut change the context to that project :)

		

(3) Tool checks against operating system for default text editor, then users this text editor to provide both *scaffolding of documentation by providing code generation directly from `gravity` command-line tool*, and *editing README.md can occur from anywhere in the project structure*.

		gravity scaffold

(4) 



==========================================================================================================

==========================================================================================================
### Unorganized Notes / Resources 
*Below is gathered notes and resources that are being used to write the first draft of the documentation.*

### Command/Runners (like Make)
https://github.com/Nananas/ymake


github.com/LadyDascalie/vex - simple with toml, could be good starting point


https://github.com/wagoodman/bashful - Use a yaml file to stitch together commands and bash snippets and run them with a bit of style. Why? Because your bash script should be quiet and shy-like (...and not such a loud mouth).

## Autoreload Daemon
https://github.com/go-playground/justdoit

### Continuous Integration (CI)
https://github.com/drone/drone - Drone provides more than just continous integration, it provides command/runners and other features of a build system

### Debug


		github.com/d4l3k/go-pry/pry


https://github.com/cs01/gdbgui 

### Editor

[xi] =====
xi-core must be executable from your PATH.

		git clone https://github.com/google/xi-editor.git
		cd xi-editor/rust
		cargo install

Running gxi

		git clone https://github.com/bvinc/gxi.git
		cd gxi
		cargo run

__go ui for xi__

[neovim] =====
https://github.com/neovim/neovim


https://github.com/dzhou121/gonvim

__plugins__
https://github.com/neovim/go-client



__Setting up neovim with a shared vimrc__

Create a symbolic link into the nvim folder and this will create a shared configuration (I think the name is wrong)

		cd ~/.local/share/nvim/
		ln -s ~/.vimrc .


__plugin system in vim(and neovim)__
call plug#begin()

Plug 'philip-karlsson/bolt.nvim', { 'do': ':UpdateRemotePlugins' }

call plug#end()
