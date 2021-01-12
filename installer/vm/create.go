package vm

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/google/uuid"
	"github.com/libvirt/libvirt-go"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// CreateOptions store various options for modifying the domain creation behavior.
type CreateOptions struct {
	// Automatically connect to the domain's first serial port after creation.
	ConnectAfterCreate bool

	// Remove the create domain after it is shutdown.
	IsTransient bool

	// Create a snapshot immediately after creating the domain, before it is
	// started.
	CreateInitialSnapshot bool
}

// CreateConfig store customization options for modifying the domain parameters
// during creation.
type CreateConfig struct {
	// Use UEFI boot instead of BIOS
	UEFI bool

	// Use the specified video device
	Video string

	// Use the specified network bridge device
	Network string

	// Use the specified amount of RAM
	Memory string
}

// Create defines a new domain using name and creating a disk image backed by
// image. If connect is true, a console is attached to the newly created domain.
// If transient is true, the domain is destroy upon shutdown.
func Create(uri, name, image string, disks []string, options CreateOptions, config CreateConfig) error {
	if name == "" {
		name = petname.Generate(2, "-")
	}

	imagesDir, err := getImagesDir()
	if err != nil {
		return err
	}

	instancesDir, err := getInstancesDir()
	if err != nil {
		return err
	}
	if options.IsTransient {
		instancesDir, err = ioutil.TempDir("", "vm-")
		if err != nil {
			return err
		}
	}

	baseImagePath := filepath.Join(imagesDir, image+".qcow2")
	if _, err := os.Stat(baseImagePath); os.IsNotExist(err) {
		return err
	}

	var domain domain
	if err := xml.Unmarshal([]byte(domainXML), &domain); err != nil {
		return err
	}
	domain.UUID = uuid.New().String()
	domain.Name = name

	overlayImagePath := filepath.Join(instancesDir, domain.UUID+".qcow2")
	cmd := exec.Command("qemu-img",
		"create",
		"-f",
		"qcow2",
		"-o",
		fmt.Sprintf("backing_file=%v,backing_fmt=qcow2", baseImagePath),
		overlayImagePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	domain.Devices.Disks[0].Source.File = overlayImagePath

	instanceDataDir := filepath.Join(instancesDir, domain.UUID)
	if _, err := os.Stat(instanceDataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(instanceDataDir, 0755); err != nil {
			return err
		}
	}

	for _, d := range disks {
		src, err := os.Open(d)
		if err != nil {
			return err
		}
		defer src.Close()

		dest, err := os.Create(filepath.Join(instanceDataDir, filepath.Base(d)))
		if err != nil {
			return err
		}
		defer dest.Close()

		if _, err := io.Copy(dest, src); err != nil {
			return err
		}

		var device string
		switch filepath.Ext(d) {
		case ".iso":
			device = "cdrom"
		case ".img":
			device = "floppy"
		default:
			device = "disk"
		}

		var drv struct {
			Name string `xml:"name,attr"`
			Type string `xml:"type,attr"`
		}
		switch filepath.Ext(d) {
		case ".qcow2":
			drv = struct {
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			}{
				Name: "qemu",
				Type: "qcow2",
			}
		default:
			drv = struct {
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			}{
				Name: "qemu",
				Type: "raw",
			}
		}

		disk := struct {
			Type   string `xml:"type,attr"`
			Device string `xml:"device,attr"`
			Driver struct {
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"driver"`
			Source struct {
				File string `xml:"file,attr"`
			} `xml:"source"`
			Target struct {
				Dev string `xml:"dev,attr"`
				Bus string `xml:"bus,attr"`
			} `xml:"target"`
			ReadOnly string `xml:"readonly,omitempty"`
		}{
			Type:   "file",
			Device: device,
			Driver: drv,
			Source: struct {
				File string `xml:"file,attr"`
			}{
				File: dest.Name(),
			},
			Target: struct {
				Dev string `xml:"dev,attr"`
				Bus string `xml:"bus,attr"`
			}{
				Dev: "hdb",
				Bus: "ide",
			},
		}
		domain.Devices.Disks = append(domain.Devices.Disks, disk)
	}

	if config.Memory != "" {
		m, err := humanize.ParseBytes(config.Memory)
		if err != nil {
			return InvalidArgumentErr{"memory", err}
		}
		// libvirt domains expect memory in KiB
		domain.Memory = m / 1024
		domain.CurrentMemory = m / 1024
	}

	if config.UEFI {
		options.CreateInitialSnapshot = false
		domainCapabilities, err := getDomainCapabilities(uri)
		if err != nil {
			return err
		}

		domain.OS.Loader = &struct {
			ReadOnly string `xml:"readonly,attr"`
			Type     string `xml:"type,attr"`
			CharData string `xml:",chardata"`
		}{
			ReadOnly: "yes",
			Type:     "pflash",
			CharData: domainCapabilities.OS.Loader.Value,
		}
	}

	if config.Video != "" {
		domain.Devices.Videos = append(domain.Devices.Videos, struct {
			Model struct {
				Type         string "xml:\"type,attr,omitempty\""
				VRAM         string "xml:\"vram,attr,omitempty\""
				Heads        string "xml:\"heads,attr,omitempty\""
				Acceleration *struct {
					Accel3d string "xml:\"accel3d,attr\""
				} "xml:\"acceleration,omitempty\""
			} "xml:\"model,omitempty\""
		}{
			Model: struct {
				Type         string "xml:\"type,attr,omitempty\""
				VRAM         string "xml:\"vram,attr,omitempty\""
				Heads        string "xml:\"heads,attr,omitempty\""
				Acceleration *struct {
					Accel3d string "xml:\"accel3d,attr\""
				} "xml:\"acceleration,omitempty\""
			}{
				Type: config.Video,
			},
		})
	}

	if config.Network != "" {
		domain.Devices.Interfaces = []struct {
			Type   string "xml:\"type,attr\""
			Source struct {
				Network string "xml:\"network,attr,omitempty\""
				Bridge  string "xml:\"bridge,attr,omitempty\""
			} "xml:\"source\""
			MAC *struct {
				Address string "xml:\"address,attr\""
			} "xml:\"mac,omitempty\" json:\",omitempty\""
			Model struct {
				Type string "xml:\"type,attr\""
			} "xml:\"model\""
		}{
			{
				Type: "bridge",
				Source: struct {
					Network string "xml:\"network,attr,omitempty\""
					Bridge  string "xml:\"bridge,attr,omitempty\""
				}{
					Bridge: config.Network,
				},
				Model: struct {
					Type string "xml:\"type,attr\""
				}{
					Type: "virtio",
				},
			},
		}
	}

	data, err := xml.Marshal(domain)
	if err != nil {
		return err
	}

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}
	defer conn.Close()

	var dom *libvirt.Domain
	if options.IsTransient {
		dom, err = conn.DomainCreateXML(string(data), libvirt.DOMAIN_START_AUTODESTROY)
		if err != nil {
			return err
		}
	} else {
		dom, err = conn.DomainDefineXML(string(data))
		if err != nil {
			return err
		}
		if options.CreateInitialSnapshot {
			err = SnapshotCreate(uri, name, "Initial state")
			if err != nil {
				return err
			}
		}
		if err := dom.Create(); err != nil {
			return err
		}
	}
	defer dom.Free()

	fmt.Println("Created " + domain.Name)

	if options.ConnectAfterCreate {
		return connectSerial(dom)
	}

	return nil
}
