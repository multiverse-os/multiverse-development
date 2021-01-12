package vm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/libvirt/libvirt-go"
)

// Destroy stops and undefines a domain by name. If force is true, the
// domain is destroyed without prompting for confirmation.
func Destroy(uri, name string, force bool) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	instancesDir, err := getInstancesDir()
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		return err
	}
	defer dom.Free()

	name, err = dom.GetName()
	if err != nil {
		return err
	}

	if !force {
		fmt.Printf("Are you sure you wish to destroy %v? (y/N) ", name)
		var response string
		fmt.Scan(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" {
			return nil
		}
	}

	state, _, err := dom.GetState()
	if err != nil {
		return err
	}
	if state == libvirt.DOMAIN_RUNNING {
		err = dom.Destroy()
		if err != nil {
			return err
		}
	}

	UUID, err := dom.GetUUIDString()
	if err != nil {
		return err
	}

	os.Remove(filepath.Join(instancesDir, UUID+".qcow2"))
	os.RemoveAll(filepath.Join(instancesDir, UUID))

	err = dom.UndefineFlags(libvirt.DOMAIN_UNDEFINE_SNAPSHOTS_METADATA)
	if err != nil {
		return err
	}

	return nil
}
