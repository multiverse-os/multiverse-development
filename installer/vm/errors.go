package vm

import (
	"fmt"
	"os"

	"github.com/libvirt/libvirt-go"
)

// A MissingPositionalArgErr occurs when a command is invoked without a required
// positional argument.
type MissingPositionalArgErr struct {
	name string
}

func (e MissingPositionalArgErr) Error() string {
	return fmt.Sprintf("error: %v is required", e.name)
}

// ErrDomainNameRequired represents a missing domain name argument.
var ErrDomainNameRequired = MissingPositionalArgErr{
	name: "domain name",
}

// ErrImageNameRequired represents a missing image name argument.
var ErrImageNameRequired = MissingPositionalArgErr{
	name: "image name",
}

// ErrTemplateNameRequired represents a missing template name argument.
var ErrTemplateNameRequired = MissingPositionalArgErr{
	name: "template name",
}

// ErrURLOrPathRequired represents a missing URL or path argument.
var ErrURLOrPathRequired = MissingPositionalArgErr{
	name: "URL or path",
}

// An InvalidArgumentErr occurs when a value passed to an argument is invalid.
type InvalidArgumentErr struct {
	name string
	err  error
}

func (e InvalidArgumentErr) Error() string {
	return "error: invalid argument: " + e.name + ": " + e.err.Error()
}

// An UnsupportedFormatErr occurs when a command is invoked that does not
// support the specified format.
type UnsupportedFormatErr struct {
	format string
}

func (e UnsupportedFormatErr) Error() string {
	return "error: unsupported format: " + e.format
}

// ErrUnsupportedJSONFormat represents a command requesting JSON when it is
// not supported.
var ErrUnsupportedJSONFormat = UnsupportedFormatErr{
	format: "JSON",
}

// ErrUnsupportedXMLFormat represents a command requesting XML when it is
// not supported.
var ErrUnsupportedXMLFormat = UnsupportedFormatErr{
	format: "XML",
}

// UnsupportedDomainCapabilityErr occurs when a domain capability is not
// supported on the host.
type UnsupportedDomainCapabilityErr struct {
	capability string
}

func (e UnsupportedDomainCapabilityErr) Error() string {
	return "error: domain capability not supported: " + e.capability
}

// LogErrorAndExit logs err and exits with a non-zero exit code.
func LogErrorAndExit(err error) {
	switch err.(type) {
	case libvirt.Error:
		e := err.(libvirt.Error)
		fmt.Fprintf(os.Stderr, "%v\n", e.Message)
		os.Exit(int(e.Code))
	default:
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
