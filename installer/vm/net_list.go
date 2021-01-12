package vm

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/libvirt/libvirt-go"
)

// NetList prints a list of all defined networks known to libvirt.
func NetList(uri string) error {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}
	defer conn.Close()

	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tACTIVE\tAUTOSTART\tPERSISTENT\t")
	for _, network := range networks {
		defer network.Free()
		name, err := network.GetName()
		if err != nil {
			return err
		}
		active, err := network.IsActive()
		if err != nil {
			return err
		}
		autostart, err := network.GetAutostart()
		if err != nil {
			return err
		}
		persistent, err := network.IsPersistent()
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n", name, active, autostart, persistent)
	}
	w.Flush()

	return nil
}
