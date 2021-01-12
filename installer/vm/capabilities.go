package vm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/libvirt/libvirt-go"
)

// Capabilities prints detailed information about the host hypervisor capabilities
func Capabilities(uri string, outputformat string) error {
	cap, err := getCapabilities(uri)
	if err != nil {
		return err
	}

	var output []byte
	switch outputformat {
	case "json":
		output, err = json.MarshalIndent(cap, "", "\t")
		if err != nil {
			return err
		}
	case "xml":
		output, err = xml.MarshalIndent(cap, "", "\t")
		if err != nil {
			return err
		}
	default:
		return UnsupportedFormatErr{outputformat}
	}
	fmt.Println(strings.TrimSpace(string(output)))

	return nil
}

func getCapabilities(uri string) (*capabilities, error) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return nil, err
	}

	data, err := conn.GetCapabilities()
	if err != nil {
		return nil, err
	}

	var cap capabilities
	err = xml.Unmarshal([]byte(data), &cap)
	if err != nil {
		return nil, err
	}

	return &cap, nil
}
