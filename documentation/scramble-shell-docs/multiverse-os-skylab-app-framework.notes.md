##
##  Multiverse OS - Skylab Application Framework
=================================================================
**Go Development Resources/Notes**

_I should be testing_, and I found a project that may help start the process:
https://github.com/stretchr/tdd-present/

[diagrams](https://github.com/cascades-fbp/flowmaker)

**cork**[a binary serialization format for go](https://github.com/abcum/cork)
<IMPORTANT>
Features
  * Simple and efficient encoding
  * Based on MsgPack serialization algorithm
  * **Stores go type information inside encoded data**
  * **Faster serialization than gob binary encoding**
  * **More efficient output data size than gob binary encoding**
  * Serializes native go types, and arbritary structs or interfaces
  * Enables predetermined encoding for structs without run-time reflection
  * Allows serialization to and from maps, structs, slices, and nil interfaces

**orbit**[wrapper for otto: A Node.js lambda environment library for Go (Golang).](A Node.js lambda environment library for Go (Golang).)
<LOOK INTO THIS TO BUILD OUT SURF MROE>

**fibre**[a simple and fast http framework for go](https://github.com/abcum/fibre)
<I like how this already supports CBOR>
**it also ahs a very nice set of middleware already, including security, gzip
Features
    Simple and efficient router
    Extensible middleware framework
    Customise when middleware should run
    Built to run with REST or Websockets
    Build APIs with RESTful methodologies
    Build APIs with Websocket methodologies
    Build APIs with JSONRpc methodologies
    Centralized and customisable error logging
    Works seamlessly with Golang's standard HTTP server
    Automatic data binding for Form, XML, JSON, CBOR, BINC, MsgPack
    Automatic response type detection for XML, JSON, CBOR, BINC, MsgPack

## WebUI 
-----------------------------------------------------------------
https://github.com/abcum/webkit/ - webkit that supports easy ability to disable javascript execution!!!

## Caching
-----------------------------------------------------------------
https://github.com/abcum/cachr **BEST CACHING LIBRARY**


=================================================================
## Example Frameworks
-----------------------------------------------------------------
**Web Frameworks**

[gear](https://github.com/teambition/gear)
Very nice framework
<Serves gRPC with HTTP on same port>


[Heavy]
[gondola](https://github.com/rainycape/gondola)
The web framework for writing faster sites, faster http://gondolaweb.com
signals, social, cache, kvstore, blobstore, httpclient, crypto, encoding, log, org, template, and more. Very robust maybe bloated

[Medium]
https://github.com/stretchr/goweb

[Light]
https://github.com/VividCortex/siesta

https://github.com/go-playground/pure - pure go radix router and gin like functionality
**System Frameworks**
[Heavy]

[Medium]
[augustine](https://github.com/plainprogrammer/augustine)
is a standard application MVC framework for Go
This is an interesting initiation because it follows a typically webui middleware scheme. I like this
		app := augustine.New()
			log.Println("Setting up middleware stack...")
			app.MiddlewareStack = []augustine.Middleware{middlewares.HelloWorld}
			log.Println("Finished setting up middleware stack...")
			log.Println("Starting augustine!")
		app.Run()




[Light]
**db**
https://github.com/abcum/rixxdb
**Special Frameworks**
[Heavy]

[Medium]

[Light]
[gmf](https://github.com/3d0c/gmf)
Go media framework It covers very basic avformat, avcodec and swscale features.
More bindings and cool features are coming soon. FFMpeg bidnings



=================================================================
## Framework Components
-----------------------------------------------------------------


=================================================================
## [USER INTERFACES] 
-----------------------------------------------------------------

=================================================================
## JS/GopherJS (WebUI & Desktop)
-----------------------------------------------------------------
[pinhole-js](https://github.com/tidwall/pinhole-js)
3D Wireframe Drawing Library for Javavscript. This is a port of pinhole for Go. 
_Very cool library, could supply very nice UI elements for Multiverse OS_

[mapbox](https://github.com/tidwall/mbdraw)


=================================================================
## GUI
-----------------------------------------------------------------


=================================================================
## Desktop WebUI
-----------------------------------------------------------------
[sciter](https://github.com/sciter-sdk/go-sciter)


**systray**
https://github.com/getlantern/systray
**notifications**

=================================================================
## CLI
-----------------------------------------------------------------
**General Utils**
[ASCII Image]
https://github.com/ichinaski/pxl

[progress bar]
https://github.com/cheggaaa/pb

**Command-line**
[cli](https://github.com/mkideal/cli)
[live color changing in terminal](https://github.com/tidwall/pony)

**Console**

**TUI**
[gocui](https://github.com/jroimartin/gocui)


=================================================================
## [APPLICATION UTILITIES] 
-----------------------------------------------------------------

=================================================================
## Linux
-----------------------------------------------------------------


**Modify Configs**
[put](https://github.com/n3phtys/put)
Go tool to ensure a given line is inside a text file. If it is not, the tool will insert the line (configurable in which line number). 
**Signals**
[go-exit](https://github.com/cheggaaa/go-exit)
This package provide handling exit signals (SIGKILL, SIGTERM, SIGQUIT and Interrupt). It's not difficult to write, just dead simple to use)

=================================================================
## Configuration
-----------------------------------------------------------------
[configer](https://github.com/go-configer/configer)
_Configuration loader that support INI, XML, YAML, JSON, HCL, TOML, Shell Environment_

[yamlconf](https://github.com/getlantern/yamlconf)
Provides mechanism for managing yaml-based configuration. _Mostly just sample code to ensure we are including, Getting, Setting and preferably live update if configuration gets changed 2-way.

[config](https://github.com/daewood/config)
_go config. support ini json yaml toml_

[frep](https://github.com/subchen/frep)
_Generate file using template from environment, arguments, json/yaml/toml config files_

[conflate](https://github.com/miracl/conflate)
_Conflate is a library that helps to merge and validate data from multiple formats (JSON/YAML/TOML), and multiple locations (filesystem paths and urls)._


=================================================================
### Design Patterns
--------------------------------------------------
**Circuit/Hystrix**
[hystrix-go](https://github.com/afex/hystrix-go)
Hystrix is a latency and fault tolerance library designed to isolate points of access to remote systems, services and 3rd party libraries, stop cascading failure and enable resilience in complex distributed systems where failure is inevitable. I think the Hystrix patterns of programmer-defined fallbacks and adaptive health monitoring are good for any distributed system. Go routines and channels are great concurrency primitives, but don't directly help our application stay available during failures.

[circuit](https://github.com/cep21/circuit)
Circuit is an efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. Learn more about the problems Hystrix and other circuit breakers solve on the Hystrix Wiki. A short summary of advantages are:
		A downstream service failed and all requests hang forever. Without a circuit, your service would also hang forever. Because you have a circuit, you detect this failure quickly and can return errors quickly while waiting for the downstream service to recover.
		Circuits make great monitoring and metrics boundaries, creating common metric names for the common downstream failure types. This package goes further to formalize this in a SLO tracking pattern.
		Circuits create a common place for downstream failure fallback logic.
		Downstream services sometimes fail entirely when overloaded. While in a degraded state, circuits allow you to push downstream services to the edge between absolute failure and mostly working.
		Open/Close state of a circuit is a clear early warning sign of downstream failures.
		Circuits allow you to protect your dependencies from abnormal rushes of traffic.

**Retry/Backoff**
https://github.com/go-playground/backoff - Backoff uses an exponential backoff algorithm to backoff between retries with optional auto-tuning functionality. 

**Events**
[events-pipeline](https://github.com/getlantern/events-pipeline)
Modular processing library for events 

**Emit**

**Pool**
[Pool](https://github.com/go-playground/pool) Package pool implements a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. <Solid Library>

**Concurrency**
[fastlane](https://github.com/tidwall/fastlane)
Fast single-producer / single-consumer channels for Go.

=================================================================
### Scheduler / Job Runner / Background Jobs
--------------------------------------------------
**Scheduler**
[jobrunner](https://github.com/bamzi/jobrunner)
JobRunner is framework for performing work asynchronously, outside of the request flow. It comes with cron to schedule and queue job functions for processing at specified time.

It includes a live monitoring of current schedule and state of active jobs that can be outputed as JSON or Html template.


**Queue**
[priority queue](https://github.com/cheggaaa/pq)
Priority queue for golang 

=================================================================
### Crypto
--------------------------------------------------
**Encryption**
https://github.com/miracl/gomiracl
_Golang wrapper for Milagro Crypto -aes,curve,ecdsa,pbkdf2,rsa,mpin,rand_

**Hashing**
[xhash](https://github.com/ricardobranco777/xhash)
_hashing like new xhash, useful for calculating differences in files, XML, and other files_

[murmur3](https://github.com/tidwall/murmur3)
This is a port of the Murmur3 hash function. Murmur3 is a non-cryptographic hash, designed to be fast and excellent-quality for making things like hash tables or bloom filters

**Signing**
[signature](https://github.com/stretchr/signature)
Signature secures web calls by generating a security hash on the client (using a private key shared with the server), to ensure that the request is geniune.

**Entropy Generation**

**Random Number Generation**
[weyl](https://github.com/tidwall/weyl)

**Steganography**
[JPEG]
[jsteg](https://github.com/lukechampine/jsteg)



=================================================================
### Internationalization / Localizaiton
--------------------------------------------------
Internationalization is made up of two core components: CLDR and i18n. 

**CLDR**
CLDR - as in day of week, `,` or `.` use in money, etc
https://github.com/go-playground/locales

**i18n**
https://github.com/go-playground/universal-translator - 18

**TZ (Time Zones**
[tz](https://github.com/go-playground/tz)
Timezone Country and Zone data generated from timezonedb.com 

**Geolocation**
[cities](https://github.com/tidwall/cities)
10,000 Cities with Latitude, Longitude, and Elevation in Go 
[geobin](https://github.com/tidwall/geobin)
The Geobin Object represents tightly packed geometry that is compatible with GeoJSON RFC 7946.


=================================================================
## [DATA]
=================================================================
## Users & User Data
-------------------------------------------------

=================================================================
## Authentication
-------------------------------------------------
**Authorization (Logging In Process**
[grpcauth](https://github.com/bamzi/grpcauth)
using certificate and token authentication in GRPC and golang



**User Sessions**
[mongostore](https://github.com/go-playground/mongostore)
Uses MongoDB to store User session data


=================================================================
## Data Types
-------------------------------------------------

=================================================================
## Validation
-------------------------------------------------

[govalidator](https://github.com/thedevsaddam/govalidator)

[validator](https://github.com/go-playground/validator)

[__uput__]:Multiverse OS Validation and user input library

**Matching / Regex**
[match](https://github.com/tidwall/match)
Match is a very simple pattern matcher where '*' matches on any number characters and '?' matches on any one character.

**Transform: Pre or Post Validation Transform**
https://github.com/go-playground/mold - provides conform, or transform type functionlaity I wanted in uput



=================================================================
## Data Formatting - Input & Output
---------------------------------------------------


**Input/Output: Encoding/Decoding**
[codecs](https://github.com/stretchr/codecs)
Provides interfaces, functions and codecs that can be used to encode/decode data to/from various formats.
		json, csv, jsonp, msgpack, services, test, xml, bson, constants


  [TOML](https://github.com/BurntSushi/toml)
  TOML parser and encoder for Go with reflection



[few](https://github.com/lukechampine/few)
Fastest Encoder in the West 

[rsraid](https://github.com/utamaro/rsraid)
RS-RAID is a library for encoding and decoding erasure codes for storage in golang. You can divide a file into k files and create redundant m files, and can recover the original file from arbitrary k files from these (k+m) files.

[!][__XOR Libraries__]
[xor](https://github.com/templexxx/xor)
Xor engine in Go

[!][__Reedsolomon Libraries__]
[Native-Go][reedsolomon](https://github.com/klauspost/reedsolomon)
Reed-Solomon Erasure Coding in Go, with speeds exceeding 1GB/s/cpu core implemented in pure Go.

[reedsolomon](https://github.com/templexxx/reedsolomon)
Erasure-engine




**Input: Parsing (Encoding)**
__We should be using CJON


[arg.js](https://github.com/stretchr/arg.js)
Lightweight URL argument and parameter parser

[objx](https://github.com/stretchr/objx)
Objx - Go package for dealing with maps, slices, JSON and other data.
		m, err := objx.FromJSON(json)

		m.Get("places[0].latlng")

[gjson](https://github.com/tidwall/gjson)
Get JSON values quickly - JSON Parser for Go 

[jjson](https://github.com/tidwall/jj)
JJ is a command line utility that provides a fast and simple way to retrieve or update values from JSON documents. It's powered by GJSON and SJSON under the hood.

It's fast because it avoids parsing irrelevant sections of json, skipping over values that do not apply, and aborts as soon as the target value has been found or updated.

[sjson](https://github.com/tidwall/sjson)
SJSON is a Go package that provides a very fast and simple way to set a value in a json document. The purpose for this library is to provide efficient json updating for the SummitDB project. For quickly retrieving json values check out GJSON.

For a command line interface check out JSONed.





**Output: Rendering (Decoding)**
https://github.com/gostores/encoding
_ASN.1,hcl,ini,json,markdown,properties,toml,xmltree,xmlsign,yaml_
_
https://github.com/night-codes/tokay-render
_Go fasthttp package for easily rendering JSON, XML, binary data, and HTML templates responses._

https://github.com/thedevsaddam/renderer
_Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go_ <Very nice templating >

[pretty print json](https://github.com/tidwall/pretty)


**Input: Form** 
https://github.com/go-playground/form - Package form Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values.



=================================================================
## Middleware
--------------------------------------------------
https://github.com/klauspost/compress
[pgzip](https://github.com/klauspost/pgzip)
Parallel gzip that breaks up the process and combines the result for faster gzip. drop in replacement for regular gzip <!>

[smallz](https://github.com/sean-/smallz)
same
[compression](https://github.com/pointlander/compress)
paralleized compression like above


[snappy compression](https://github.com/golang/snappy)
best for streaming

==================================================
## [NETWORKING]
--------------------------------------------------


**Lowest Level: Raw, Device**


**Medium Level: TCP / UDP / KCP...**


[wave](https://github.com/go-playground/wave)
The intention of this library is to provide a thin wrapper around the std net/rpc package allowing the user to add functionality via hooks instead of creating a whole new framework.



**Highest Level: HTTP, gRPC,...**

=================================================================
## Event Loop Networking
--------------------------------------------------

[evio](https://github.com/tidwall/evio)
evio is an event loop networking framework that is fast and small. It makes direct epoll and kqueue syscalls rather than using the standard Go net package, and works in a similar manner as libuv and libevent. _Supports multinetwork binding_

=================================================================
## Raft
--------------------------------------------------

[finn](https://github.com/tidwall/finn)
Fast Raft framework using the Redis protocol for Go

=================================================================
## gRPC
--------------------------------------------------
[lile](https://github.com/lileio/lile)
Lile is a generator and set of tools/libraries to help you quickly create services that communicate via gRPC (REST via a gateway) and publish subscribe.

=================================================================
## Game Networking
--------------------------------------------------

[nano](https://github.com/lonnng/nano)
Nano is an easy to use, fast, lightweight game server networking library for Go. It provides a core network architecture and a series of tools and libraries that can help developers eliminate boring duplicate work for common underlying logic. The goal of nano is to improve development efficiency by eliminating the need to spend time on repetitious network related programming._Nano was designed for server-side applications like real-time games, social games, mobile games, etc of all sizes._

=================================================================
#### HTTP
--------------------------------------------------
**Rest API**
https://github.com/gin-gonic/gin

[Rest API Authentication]
[hmac-signature](https://github.com/dcu/hmac-signature)
because JWT is not secure. this is the proper way to validate for API calls


[Rest API](https://github.com/rs/rest-layer)
REST Layer is an opinionated framework. Unlike many API frameworks, you don't directly control the routing and you don't have to write handlers. You just define resources and sub-resources with a schema, the framework automatically figures out what routes need to be generated behind the scene. You don't have to take care of the HTTP headers and response, JSON encoding, etc. either. REST layer handles HTTP conditional requests, caching, integrity checking for you.

=================================================================


**Websockets**
[claws](https://github.com/thehowl/claws)
an Awesome WebSocket CLient.

[ws](https://github.com/go-playground/ws)
ws creates a hub for WebSocket connections and abstracts away allot of the boilerplate code for managing connections using Gorilla WebSocket 

**Middleware**
https://github.com/tdewolff/minify

**Gin WebFramework**
[AutoTLS for Gin](https://github.com/gin-gonic/autotls)

[Gin Middleware](https://github.com/dogenzaka/gin-tools)
gzip, logging, validation

[Gin Boilerplate with Mgo](https://github.com/Luncher/go-rest)

[Gin Scaffolding](https://github.com/dcu/gin-scaffold)
<IMPORTANT I want to build something similar to rails and code generation is very important to a solid framework as shown by rails. Code generation is massively useful in speeding up development. Using templates for Go files.>

[Gin Control Panel Example](https://github.com/night-codes/summer)
Very very nice control panel using MGO and Gin

**Alternative Web Frameworks**
https://github.com/night-codes/tokay - The way it stacks middleware is very nice, this feature shouldb e built into gin

=================================================================
## [MATH]
-----------------------------------------------
**Math**
    [rollingavg](https://github.com/tidwall/rollingavg)
    A rolling averager for Go.



=================================================================
## [DATABASE]

https://github.com/abcum/emitr - basic emitter, would be good to add to db hooks
-----------------------------------------------
**Redis Protocol**


[redcon](https://github.com/tidwall/redcon)
Redis compatible server framework for Go
			Create a Fast custom Redis compatible server in Go
			Simple interface. One function ListenAndServe and two types Conn & Command
			Support for pipelining and telnet commands
			Works with Redis clients such as redigo, redis-py, node_redis, and jedis
			TLS Support



**Merkle Trees**
[merkletree](https://github.com/cbergoon/merkletree)


**Berkley Tree DB**
[btrdb-server](https://github.com/BTrDB/btrdb-server)

**Red-Black Tree**
[rbtree](https://github.com/sakeven/RbTree)

**Embedded**
[Trees]
    [+][go-radix](https://github.com/armon/go-radix)
    Provides the radix package that implements a radix tree. The package only provides a single Tree implementation, optimized for sparse nodes.
    As a radix tree, it provides the following:
				O(k) operations. In many cases, this can be faster than a hash table since the hash function is an O(k) operation, and hash tables have very poor cache locality.
				Minimum / Maximum value lookups
				Ordered iteration
    [+][go-radix-immutable](https://github.com/hashicorp/go-immutable-radix)
    Same as above but immutable

    [+][rtree](https://github.com/tidwall/rtree)
    This package provides an in-memory R-Tree implementation for Go, useful as a spatial data structure. It has support for 1-20 dimensions, and can store and search multidimensions interchangably in the same tree.

    [+][bbtree](https://github.com/tidwall/bbtree)
    

    [+][btree](https://github.com/tidwall/btree)
    Added features over the google implementation it was based on

    [+][paired btree](https://github.com/tidwall/pairtree)
    This package provides an in-memory B-Tree implementation for Go, useful as an ordered, mutable data structure.

    [+][pair-rtree](https://github.com/tidwall/pair-rtree)
    A specialized hybrid 2D/3D R-Tree library for Go. It uses pair objects with geobin values. Supports KNN searching and rectangle transformations.

    [+][rbush](https://github.com/tidwall/rbush)
    __Spatial index is a special data structure for points and rectangles that allows you to perform queries like "all items within this bounding box" very efficiently__ (e.g. hundreds of times faster than looping over all items). It's most commonly used in maps and data visualizations.

[Heap]
  [+][tinyqueue](https://github.com/tidwall/tinyqueue)
  tinyqueue is a Go package for binary heap priority queues. Ported from the tinyqueue Javascript library.


[Key/Value] 
  [+][pair:low memory key/value store](https://github.com/tidwall/pair)
  create low memory key/value objects in Go 



=================================================================
## [BUILDING / TESTING/ RUNTIME]
=================================================================
### Debugging
-----------------------------------------------
[Pry](https://github.com/d4l3k/go-pry)
Based on a Ruby tool of the same name, that allows you to drop in breakpoints that stop the process and let you interact with the application state from an interactive shell

=================================================================
### Testing
-----------------------------------------------

**Mocking**


**TDD**
[assert](https://github.com/go-playground/assert)

**BDD**
[ginkgo](https://github.com/onsi/ginkgo)
Most popular BDD framework for Go

=================================================================
### Profiling / Preformance / Benchmarking
---------------------------------------------------
**Benchmarking**

**Metrics**
[go-metrics](https://github.com/bamzi/go-metrics)
This library provides a metrics package which can be used to instrument code, expose application metrics, and profile runtime performance in a flexible manner.

**Stats**
[stats](https://github.com/go-playground/stats)


[Sample Code]
https://github.com/appleboy/gorush - nice graphs and rest api stats to sample fromk

=================================================================
## Version Control
[go-github](https://github.com/google/go-github)
Go library for accessing the GitHub API http://godoc.org/github.com/google/go-github/github

=================================================================
## Multiverse + Hackwave UI

=================================================================
## NTP

https://github.com/JeffBelgum/ntp

=================================================================
## Image Modification Server (Resize typcially, thumbnial)


https://github.com/thoas/picfit


=================================================================
## Ruby Config


https://github.com/k0kubun/itamae-go

=================================================================
## Web Renderer 


**rust**
[webrenderer](https://github.com/servo/webrender)
[toy_web_rendering_engine](https://github.com/nham/toy_web_rendering_engine)
**go**
https://github.com/vecty/blink-idl

=================================================================
## chord


https://github.com/armon/go-chord

=================================================================
## Rust to Go


Using the ability to call Rust from Go I would love to get webrender to run in Go and function as a UI based on CSS and crippled JS or no JS.


=================================================================
## Web Server In Go


https://github.com/puma/puma-dev = MOST FEATURE RICH WEB SERVER I HAVE SEEN
^ Most explicit
=================================================================
https://github.com/kurin/blazer
https://github.com/mosuka/indigo full text serach

=================================================================
## Database JSON


[jfc](https://github.com/n3phtys/jfc)
Reads JSON file and maps it all to key/value map

=================================================================
## Custom net userspace


https://github.com/orivej/tcpassembly
=================================================================
## Workers Deterministic ticking



https://github.com/VividCortex/multitick

=================================================================
## Mux/Router


https://github.com/claygod/Bxog
https://github.com/go-zoo/bone
https://github.com/ranbochen/h2stream
https://github.com/Kiricon/Rapid

=================================================================
## Packet multiplexer


https://github.com/midbel/pex

=================================================================
## RSS

https://github.com/SlyMarbo/rss

=================================================================
## Exchange

https://github.com/cyanly/gotrade

=================================================================
## Market


https://github.com/foomo/shop
https://github.com/johnsiilver/boutique
https://github.com/streadway/quantile/
**Tree comments/categories**
https://github.com/hit9/htree
https://github.com/spouk/tree_comments_go
[categories]
https://github.com/mrsinham/catego


=================================================================
## Fourth Like VM


https://github.com/acook/blacklight
https://github.com/unixdj/forego

