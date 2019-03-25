##
##  Multiverse OS Scramble Suit
=================================================================

https://github.com/corpix/uarand - random user agent generator

=================================================================
## SSH (Key & Remote Host) Management 
-----------------------------------------------------------------
A CLI tool that manages SSH keys and hosts segregated by identity
supporting Hostnames for remote IP addresses that do not require 
changes to the Host file, SSH Key generation, rotation, preference
migration, advanced connection routing (supporting SOCKS proxies 
and multi-hop connections).
##### SSH Client Features
    * **Per scramble suit identity SSH key management** with *auto-*
      *matic key rotation*, per-host key generation
    * **Host aliasing** that does not require changes to the hostfile
      (so it does not pollute tab autocomplete)
    * **Preference migration** to new hosts and expiry scripts to 
      delete files even in the event of unexpected disconnection
    * **Advanced connection routing** supporting proxies (SOCKS, Tor,
      reverse SSH, and so on...)
      and multi-hop connections 
    * **Simplified scripting system** supporting hooks on
				(1) on first login, (2) on login, (3) on log off,... etc
       [Use cases] delete history on logoff, check for other users
       on login, prepare development environment and update source
       code against repositories on login, setup TAILF log watching
    * Clientside emulation to provide lag free UI experience when
      connected to high latency servers
    * Simple reverse SSH, port forwarding, and reverse proxy support
      with a simplified API and use patterns in the new SSH client 
      CLI
    * **Stealth mode** that provides a simplified use that automatically
      conceals usage, hides presence, prevents history generation
    * **Command multiplexing** Server categories, command multiplexing,
      execution across computers by class, and other intuitive filters
    * **Simplified key installation** supporting auto-logging in, 
      disabling key authentication, installing key and testing
      connection in a single command. Transfer keys between servers,
      including two remote servers and authorized keys management [auto ensuring key files, configuration, have the correct permissions]
        [Similar to: ssh-copy-id]
    * **See key status** quickly display last time a key was used, when it was generated, who is logged in using what keys. 
    * **Track bandwidth stats** usage, in real-time and over time
    * **Simplified between server copying, overloading `cp` to render scp
      obsolete
    * **Simplified configuration**, secure defaults that are completely
      backwards compatible, keeping the same location and files in
      addition to new configuration files and locations

##### SSHd Server Features
    * **Simplified jailing** command restriction, Virtual FS jailing,
      encrypted output ontop of typical SSH encryption and requirement
      of signed input
    * **Simplified application framework** to support building applications
      ontop of the secured protocol; [Examples] SSH MUChat
    * **Improved encryption and compression** options
    * **Simplified configuration**, secure defaults that are completely
      backwards compatible, keeping the same location and files in
      addition to new configuration files and locations
    * **Stealth mode** that provides hidden connection
    * **Simplified reverse proxy** supporting WebUI and on-demand
      mapping to URL routes 
    * Multi hop and rendevous double reverse proxy (like tor) connections
    * **Coop mode** allowing multi-user shared connections, like tmux 
      but more than just multiplexing, built for cooperative use supports
      on-demand split-screening
    * MITM proxy
    * **WebUI** to share screen, or access via Websockets/JS
    * **Simplified decentralized/centralized authenitcation** Oauth2
      support, (H|T)OTP (sometimes called 2-factor authentication)
    * **Honeypot provider** jailing the user in a stealth container 
      andtracking the use.
    * **Simplified daemon to accept reverse proxy connections**
    * **Daemon to accept and install key into authorized keys** file upon
      successful authorization
    * **Cluster support** to simplify management and interconnects, complex
      routing and more
    * **simple git hosting support** with HTTP serving too
    * **Simplified scripting system with convienent hooks** on real-time
      evens such as:
				(1) on-client-login, (2) on-client-logoff, (3) on-first-start, (4) on start, (3) before daemon shutdown, (4) on-client-sudo, (4) on client executing command that matches x regex, (6) on client is logged in x amount of time, (7) on client is idle x amount of time
    * **Command roll-back, state versioning, branching**


=================================================================
## Rust Version
https://github.com/briansmith/ring
https://github.com/RustCrypto/hashes - all hashes (sha3, ..)
https://github.com/RustCrypto/block-ciphers - all block cyphers (aes-sof, blowfish, des, ...)
https://github.com/RustCrypto/password-hashing - PBKDF2
https://github.com/RustCrypto/MACs - (cmac, pmac, hmac)
https://github.com/RustCrypto/rust-crypto-decoupled - symetric (chacha20, chacha20poly1305, hc128, rc4, salsa20, sosemanuk)

https://github.com/RustCrypto/rust-crypto-decoupled/blob/master/etc/curve25519
https://github.com/RustCrypto/rust-crypto-decoupled/tree/master/etc/ed25519
https://github.com/ryantm/rust_ecdsa


## Scramble Suit Key System
https://github.com/rjsberry/zzzpass = RESTFUL passstore


<IMPORTANT>https://seattle.github.com/awnumar/memguard


https://github.com/chain/chain/blob/chainkd-dh/docs/protocol/specifications/chainkd.md
----------------------------------------------------------------
https://github.com/dedis/kyber/blob/master/sign/eddsa/eddsa.go - users supercop
https://github.com/dedis/kyber/blob/master/sign/eddsa/eddsa.go

https://github.com/substack/ed25519-supercop (JS)


--- https://github.com/paragonie/paseto [Paseto is everything you love about JOSE (JWT, JWE, JWS) without any of the many design deficits that plague the JOSE standards.]


__Nearly identical tow hat im building but different__ Could be a great refences
https://github.com/dedis/cothority


[very clean crypto lib](https://github.com/dedis/kyber)
**This should become the basis, it is VERY powerful, supports multiple ecdh, even better than below**
[generic ecdh implementation](https://github.com/aead/ecdh)
This is cool because it provides a very generic implementation that
I could build around, that way each crypto follows this interface spec
then it can easily switch ou the crypto being used but use the same
consistent API 
**interesting bits**
In this library I found a way to use scalarmult function a low level function to combine the public and private key to generate a secret. This may be useful in the future. **You use YOUR privatekey and the OTHER persons public key to compute a SHARED SECRET!**
		func (g genericCurve) ComputeSecret(private crypto.PrivateKey, peersPublic crypto.PublicKey) (secret []byte) {
			priKey, ok := checkPrivateKey(private)
			if !ok {
				panic("ecdh: unexpected type of private key")
			}
			pubKey, ok := checkPublicKey(peersPublic)
			if !ok {
				panic("ecdh: unexpected type of peers public key")
			}

			sX, _ := g.curve.ScalarMult(pubKey.X, pubKey.Y, priKey)

			secret = sX.Bytes()
			return
		}


## Multiverse Bitcoin
----------------------------------------------------------------
# Bitcoin
Cryptocurrency is deeply woven into the Mutliverse operating system, it both provides the currency to pay for decentralized services and Multiverse OS provides the foundation for Bitcoin to escape centralized services; because like AWS to the Internet, centralized services have crippled the functionality of Bitcoin, in addition to diverting development towards centralized services that are counter to the original design premise.
## WebKit GO Bindings 
https://github.com/abcum/webkit
**THIS HAS EASY ABILITY TO DISABLE JAVASCRIPT!**
https://github.com/klauspost/pgzip - parallel gzip, drop in replacement for regular gzip, improves speed signfiicantly by brekaing up the files doing parallel compression and combining
## Cahce
https://github.com/abcum/cachr - **best caching library**
#### Examples / Samples & Useful Source Code
https://github.com/tyler-smith/go-bip32 - KEY SYSTEM
https://github.com/wenweih/bitcoin_address_protocol
	[ergo](https://github.com/ergoplatform/ergo)
	This repository contains the reference implementation of the Ergo Platform protocol, which is an alternative to the
	Bitcoin protocol. Differences from Bitcoin
    * [Memory-hard Proof-of-Work function Equihash]
    * [New modes of operation: light-fullnode, light-SPV, hybrid modes]
    * [Alternative transactional language, which is more powerful that Bitcoin Script but also safe against heavy validation attacks]
    * [Alternative fee model with mandatory storage-rent component]

#### Cryptocurrency Wallet / Key Manager
A Bitcoin wallet (and other cryptocurrencies) have been sabotagoed by developers trying to cash in on issuing securities, tokens and other things to transfer wealth that could have potentially went to honest open source development into the hands of greedy, myopic shills who rely on hype over substance.

This entire time many of the features promised by other cryptocurrencies could have simply been abstracted onto Bitcoin using an overlay protocol, through the use of a customized wallet.

For this reason Multiverse will provide a multisig wallet with advanced features not yet seen because no one has bothered to advance beyond the cloning the spec wallet. Combining wallet functionality with OHT will enable much richer multisig support, shamirs secret sharing backups, and using OP_RETURN and OHT rings to execute code without the limitations of Ethereum and other projects that have failed while poisoning concepts like distributed autonomous organizations and doing damage by convincing so many that blockchains or ethereum is required to create decentralized applciations; as if they did not exist before Ethereum. 

_Wallet development will occur utilizing both Rust & Go._
  **Gravity Wallet**
  	[Keys: HD Keys, Bec32, Mneomics & Ephemeral Hierarchal Key Trees]
			[BIP44](https://github.com/wuminzhe/bip44)

	  [Bitcoin Smart Contracts & Payment Channels]
		  Utilizing payment channels, lightning network upgrades to the multisignature system will enable smart contracts and most importantly decentralized autonomous corporations built ontop of distributed permissioned ledgers, using overlay tools, Bitcoin blockchain and other blockchains.
		  The system will be much more powerful, cost much less to use, be incredibly flexible and not confined to one specific VM type.

  **Blockchain Explorers**
  	[JS:bitcore-node](https://github.com/bitpay/bitcore-node)
    <IMPORTANT>_bitcore-node should be converted to pure Go using gopherjs_

==================================================================
## Multiverse Media
==================================================================


https://github.com/laher/kv
### Media Collection

https://github.com/barsanuphe/endive  < interesting diea for media collection management>

* Special FS that automatically organizes files by type/length/category while leaving all thea actual structure the same so torrents work without issue.

## Shell options
https://github.com/NeowayLabs/nash
### Transcoding
* Transcoding using FFMPEG
https://github.com/fzakaria/transcoding

## DB
https://github.com/abcum/blist - binary time series list 
## Codecs
[!][FLAC]

* https://github.com/mewkiz/flac
  Package flac provides access to FLAC (Free Lossless Audio Codec) streams.



# Multiverse OS Scramble Suit

### CONSOLE
https://github.com/Armored-Dragon/ensh/blob/master/main.go
# Stream Control Transmission Protocol
On the level of TCP and UDP, SCTP offers functionality from both TCP and UDP.




https://github.com/ishidawataru/sctp 



## Git Repository Analysis
https://github.com/src-d/hercules



## Bablefish
Use babel fish to do machine learning analysis on your source code. This will impove over time, also allow deep analysis and debugging live. 


https://github.com/bblfsh/client-go
Babelfish Go client library provides functionality to both connecting to the Babelfish server for parsing code (obtaining an UAST as a result) and for analysing UASTs with the functionality provided by libuast.

https://github.com/bblfsh/go-

## ENv management
https://github.com/antham/envh
## net tools
https://github.com/zephyrproject-rtos/net-tools
-------------------------------------------------
## RESULTS OF PASSIVE SCANNING
[Grafeas](https://github.com/grafeas/grafeas)
Grafeas defines metadata API spec for computing components (e.g., VM images, container images, jar files, scripts) that can assist with aggregations over your metadata. Grafeas uses two API concepts, a note and an occurrence. This division allows 3rd party metadata providers to create and manage metadata on behalf of many customers. Additionally, the division also allows implementation of access control settings that allow fine grain access control.
## web ui
https://github.com/go-http-utils/cookie
## WebKit GO Bindings 
https://github.com/abcum/webkit
**THIS HAS EASY ABILITY TO DISABLE JAVASCRIPT!**
## IDENTITY MANAGEMENT
Identity management MUST include management of:
  * SSH Keys
  * Cryptographic keys
  * Remote servers
  * Environmental Variables
  * Projects
  * Aliases
  * Snippets
  * Notes



*Examples:*
[vaulted](https://github.com/miquella/vaulted)
*Spawn environments from securely stored secrets.* With so many secrets floating around in our modern lives, it's a wonder we're able to keep track of any of them! vaulted allows you to create vaults of related secrets and then spawn sessions with these secrets. Vaults can contain secure environment variables, AWS credentials, or SSH keys (RSA, DSA, & ECDSA). vaulted also attempts to insulate spawned environments from other environments on the system.



-=-----------------------
===================================================
## Tables
https://github.com/gholt/brimtext

## Homograph Attack
https://github.com/dutchcoders/homographs - Brute force fonts to reveal homographs
https://github.com/e-XpertSolutions/punycode-attack/ - attack domains with new UTF8 support-
https://github.com/jsidrach/idn-homograph-attack

## Index Search 
[bolt db indexing]
https://github.com/blevesearch/segment

## Fuzzy
https://github.com/sahilm/fuzzy

## Search Engine
[in memory]
https://github.com/tyleregeto/memsearch

## Color Text With Regex
https://github.com/augustoroman/highligt
https://github.com/ktat/kolorit
[active rainbow]
https://github.com/tidwall/pony

## Pure Go Git
https://github.com/src-d/go-git A highly extensible Git implementation in pure Go.
https://github.com/speedata/gogit Pure Go read access of a Git repository

## Homograph Attack
https://github.com/dutchcoders/homographs - Brute force fonts to reveal homographs
https://github.com/e-XpertSolutions/punycode-attack/ - attack domains with new UTF8 support-
https://github.com/jsidrach/idn-homograph-attack
==================================================
# Multiverse OS Scramble Suit
### Key System ###############################
https://github.com/tyler-smith/go-bip32 - KEY SYSTEM
### Password Managers ###############################
[SSH]
[cryptorious](https://github.com/malnick/cryptorious)
CLI-based encryption for passwords and random data
		rename	 Rename an entry in the vault
		rotate	 Rotate your cryptorious SSH keys and vault automatically
		delete	 Remove an entry from the cryptorious vault
		decrypt	 Decrypt a value in the vault `VALUE`
		encrypt	 Encrypt a value for the vault `VALUE`
		generate Generate a RSA keys or a secure password.	



### Modifications to Debian commands ###############################
* APT
  * If a package is mispelled, perhaps try to fuzzy find and replace automatically or just remove and install 
### Bash / Shell Preferences Install ###############################
* Install a color alias/variable scheme like this but using Go directly instead of bash script
https://github.com/cep21/jackbash/blob/master/term_colors 

Additionally, then use these colors to assign to different file types, or better file categories:
		export LS_COLORS="no=00:\
		fi=00:\
		di=01;36:\
		ln=01;36:\
		pi=40;33:\
		so=01;35:\
		do=01;35:\
		bd=40;33;01:\
		cd=40;33;01:\
		or=40;31;01:\
		ex=01;32:\
		*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.zip=01;31:*.gz=01;31:*.bz2=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:\
		*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:\
		*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.avi=01;35:\
		*.ogg=01;35:*.mp3=01;35:*.wav=01;35:\
		";

I really love the idea of exporting --color=auto on grep, similarly for ls would be great. Setting firefox. A go tool to manage user env variables is clearly desirable.

		#export GREP_OPTIONS='--color=auto'
		export GIT_CEILING_DIRECTORIES
		GIT_CEILING_DIRECTORIES=$(echo $HOME | sed 's#/[^/]*$##')  # Either /home(linux) or /Users(mac)
		export HISTFILESIZE=1000000000
		export HISTSIZE=1000000
		export PROMPT_COMMAND='history -a'
		export BROWSER='firefox'
		#export LANG='en_US.utf8'
		export LANG='C' # Testing: Try out the C locale
		if [ -f "$HOME/.inputrc" ]; then
			export INPUTRC="$HOME/.inputrc"
		fi;
		export MAN_AUTOCOMP_FILE
		MAN_AUTOCOMP_FILE="/tmp/man_completes_$(whoami)"

		
		# GNU vs BSD ls for color
		(ls --color=tty &> /dev/null)
		if [ $? -eq 0 ]; then
			export LS_COLOR='--color=tty'
		else
			export LS_COLOR='-G'
		fi;


### Cryptography ###############################
**hdkeys**
https://github.com/NebulousLabs/hdkey
### Shell / Console ###############################
https://github.com/sajari/fuzzy
### Ruby 
https://github.com/goby-lang/goby
Not ruby but ruby inspried VM 
### Queue 
https://github.com/benmanns/goworker
[Terminal Emulator]
https://github.com/gnunn1/tilix
https://github.com/kovidgoyal/kitty
C GPU terminal emualtor (1 of 2 that use gpu)
[search files]
https://github.com/donomii/tagdb
[coreutils]
https://github.com/as/torgo
[UI]
dialog boxes that match gnome, look very nice
https://github.com/gen2brain/dlgs
[GO VERSION OF PASS]
https://github.com/aviau/gopass
[encoding](https://github.com/gostores/encoding)
    ASN.1
    hcl
    ini
    json
    markdown
    properties
    toml
    xmltree
    xmlsign
    yaml
### Example / Sample Code ###############################
https://github.com/miracl/sample-bank
_sample-bank is a relying party application for demonstrating and testing DVS (Designated Verifier Signature) workflow. This server allows testing DVS flow with minimal setup and is aimed only for testing purposes. The authentication mechanism is omitted to allow minimal effort from the client application using this server._
