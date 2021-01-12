package vm

import "github.com/libvirt/libvirt-go"

func findIP(mac string) (string, error) {
	var err error
	var ip string
	for _, c := range []string{"qemu:///system", "qemu:///session"} {
		conn, err := libvirt.NewConnect(c)
		if err != nil {
			continue
		}
		defer conn.Close()

		networks, err := conn.ListAllNetworks(0)
		if err != nil {
			continue
		}

		for _, network := range networks {
			defer network.Free()
			leases, err := network.GetDHCPLeases()
			if err != nil {
				continue
			}
			for _, lease := range leases {
				if lease.Mac == mac {
					ip = lease.IPaddr
					break
				}
			}
		}
	}
	if err != nil {
		return "", err
	}
	return ip, nil
}
