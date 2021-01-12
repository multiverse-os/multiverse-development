package vm

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

// Up looks up a defined domain by name and starts it.
func Up(uri, name string, connectAfterUp bool) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		return err
	}
	defer dom.Free()

	state, _, err := dom.GetState()
	if err != nil {
		return err
	}

	switch state {
	case libvirt.DOMAIN_SHUTOFF:
		if err := dom.Create(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("error: cannot start VM in state: %v", state)
	}

	if connectAfterUp {
		return connectSerial(dom)
	}

	return nil
}
