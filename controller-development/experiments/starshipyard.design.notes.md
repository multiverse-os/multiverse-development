Starship Yard Design 
______

Multiverse OS webframework starshipyard needs to migrate to utilizing the same
method being developed for Portfile declaration via Ruby.

**Extracting Ruby MVC from a Rails project**
Currently we are successfully building our own wrapper around Rails through the
use of a Rack Server written in Go. And while this is a fantasic idea, we could
get further preformance increases

Using this technique we should parse all the files in the app folder into the Go
web application. This would enable us to avoid determining the best way to
layout the MVC logic which has been an issue with continued development

This would make web development in Go frictionless. One could even entirely
avoid using Javascript entirely by using Gopherscript, which also improves the
preformance. 

All that would be required is using a Ruby script to parse the Ruby file.
Originally the idea was to parse the Ruby files with our Ruby executable. But
using Go Grubby, we may be able to avoid that requirement, and the complexity of
such and endeavor should be considered. 

#### Rack too 
We should continue using our Go Rack server, infact we should ensure we offer C,
Go and Rack since it would be very easy to accomplish. 

#### Model 
[Active RecordReference](https://github.com/rails/rails/tree/master/activerecord) 




```
class Account < ActiveRecord::Base
  validates :subdomain, :name, :email_address, :password, presence: true
  validates :subdomain, uniqueness: true
  validates :terms_of_service, acceptance: true, on: :create
  validates :password, :email_address, confirmation: true, on: :create
end
```


#### Controller 

```
```

#### View


```
```

