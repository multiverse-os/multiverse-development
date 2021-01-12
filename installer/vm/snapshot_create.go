package vm

import (
	"encoding/xml"
	"fmt"

	"github.com/libvirt/libvirt-go"
)

// SnapshotCreate saves a new snapshot for the given domain.
func SnapshotCreate(uri, domainName, snapshotName string) error {
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

	var snapshot domainSnapshot
	if err := xml.Unmarshal([]byte(domainSnapshotXML), &snapshot); err != nil {
		return err
	}

	if snapshotName != "" {
		snapshot.Name = snapshotName
	}

	data, err := xml.Marshal(snapshot)

	snap, err := dom.CreateSnapshotXML(string(data), 0)
	if err != nil {
		return err
	}

	snapshotName, err = snap.GetName()
	if err != nil {
		return err
	}

	fmt.Println("Created snapshot " + snapshotName)

	return nil
}
