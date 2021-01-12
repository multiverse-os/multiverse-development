package vm

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/libvirt/libvirt-go"
)

// SnapshotList prints a table of snapshots for the given domain.
func SnapshotList(uri, name string) error {
	var err error

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		return err
	}
	defer dom.Free()

	snapshots, err := dom.ListAllSnapshots(0)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tCREATED\tSTATE\t")
	for _, snapshot := range snapshots {
		var ds domainSnapshot
		xmlDesc, err := snapshot.GetXMLDesc(0)
		if err != nil {
			return err
		}
		if err := xml.Unmarshal([]byte(xmlDesc), &ds); err != nil {
			return nil
		}
		sec, err := strconv.ParseInt(ds.CreationTime, 10, 64)
		if err != nil {
			return err
		}
		creationTime := time.Unix(sec, 0)
		fmt.Fprintf(w, "%v\t%v\t%v\n", ds.Name, creationTime, ds.State)
	}
	w.Flush()

	return nil
}
