package vm

import (
	"fmt"
	"strings"

	"github.com/libvirt/libvirt-go"
)

// Down looks up an active domain by name and stops it. If force is true, the
// domain is stopped without prompting for confirmation. If graceful is true,
// the domain is shutdown gracefully.
func Down(uri, name string, force bool, graceful bool) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		return nil
	}
	defer dom.Free()

	if !force {
		fmt.Printf("Are you sure you wish to stop %v? (y/N) ", name)
		var response string
		fmt.Scan(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" {
			return nil
		}
	}

	if graceful {
		if err := dom.Shutdown(); err != nil {
			return err
		}
	} else {
		if err := dom.Destroy(); err != nil {
			return err
		}
	}

	return nil
}
