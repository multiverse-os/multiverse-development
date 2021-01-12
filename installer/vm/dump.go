package vm

import (
	"fmt"
	"strings"

	"github.com/libvirt/libvirt-go"
)

// Dump prints the XML description of the given domain.
func Dump(uri, name string) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		return err
	}
	defer dom.Free()

	data, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return err
	}

	fmt.Println(strings.TrimSpace(data))

	return nil
}
