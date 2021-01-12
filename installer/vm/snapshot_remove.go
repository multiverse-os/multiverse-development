package vm

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

// SnapshotRemove deletes the given snapshot for the given domain.
func SnapshotRemove(uri, domainName, snapshotName string) error {
	var err error

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}
	defer dom.Free()

	snapshot, err := dom.SnapshotLookupByName(snapshotName, 0)
	if err != nil {
		return err
	}

	err = snapshot.Delete(0)
	if err != nil {
		return err
	}

	fmt.Println("Deleted snapshot " + snapshotName)

	return nil
}
