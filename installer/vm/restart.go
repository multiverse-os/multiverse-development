package vm

import (
	"fmt"
	"strings"

	"github.com/libvirt/libvirt-go"
)

// Restart looks up a domain by name and restarts it. If force is true, the
// domain is restarted without prompting for user confirmation first. If
// graceful is false, the dom is reset forcibly rather than being sent a
// restart signal.
func Restart(uri, name string, force bool, graceful bool) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	var dom *libvirt.Domain
	dom, err = conn.LookupDomainByName(name)
	if err != nil {
		return nil
	}
	defer dom.Free()

	if !force {
		fmt.Printf("Are you sure you wish to restart %v? (y/N) ", name)
		var response string
		fmt.Scan(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" {
			return nil
		}
	}

	if graceful {
		if err := dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT); err != nil {
			return err
		}
	} else {
		if err := dom.Reset(0); err != nil {
			return err
		}
	}

	return nil
}
