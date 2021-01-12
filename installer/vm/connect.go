package vm

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/libvirt/libvirt-go"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

const escapeSequence = byte(']') ^ 0x40

// Connect opens a connection to a domain by name. The mode argument determines
// the connection mode: either "ssh" or "console".
func Connect(uri string, name string, mode string, user, identity string) error {
	var err error

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return err
	}

	var dom *libvirt.Domain
	dom, err = conn.LookupDomainByName(name)
	if err != nil {
		return err
	}
	defer dom.Free()

	switch mode {
	case "ssh":
		return connectSSH(dom, user, identity)
	case "console":
		return connectConsole(dom)
	case "serial":
		return connectSerial(dom)
	default:
		return fmt.Errorf("error: unsupported connection mode: %v", mode)
	}
}

func connectSSH(dom *libvirt.Domain, user, identity string) error {
	if identity == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		identity = filepath.Join(homeDir, ".ssh", "id_rsa")
	}
	key, err := ioutil.ReadFile(identity)
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.RetryableAuthMethod(ssh.PasswordCallback(func() (secret string, err error) {
				fmt.Print("Password: ")
				data, err := terminal.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					return "", err
				}
				return string(data), nil
			}), 3),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	data, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return err
	}
	var domain domain
	if err := xml.Unmarshal([]byte(data), &domain); err != nil {
		return err
	}

	var addr string
	for _, iface := range domain.Devices.Interfaces {
		ip, err := findIP(iface.MAC.Address)
		if err != nil {
			return err
		}
		if ip != "" {
			addr = ip
			break
		}
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", addr), config)
	if err != nil {
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stderr, stderr)

	fd := int(os.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		oldState, err := terminal.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer terminal.Restore(fd, oldState)

		termWidth, termHeight, err := terminal.GetSize(fd)
		if err != nil {
			return err
		}

		if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
			return err
		}
	}

	if err := session.Shell(); err != nil {
		return err
	}

	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}

func connectSerial(dom *libvirt.Domain) error {
	return connectCharDev(dom, "serial0")
}

func connectConsole(dom *libvirt.Domain) error {
	return connectCharDev(dom, "console1")
}

func connectCharDev(dom *libvirt.Domain, devName string) error {
	var err error

	name, err := dom.GetName()
	if err != nil {
		return err
	}
	fmt.Println("Connected to " + name + " (" + devName + ")")
	fmt.Println("Escape character is ^]")

	conn, err := dom.DomainGetConnect()
	if err != nil {
		return err
	}

	oldstate, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), oldstate)

	signal.Ignore(syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGPIPE)
	defer signal.Reset()

	stream, err := conn.NewStream(0)
	if err != nil {
		return err
	}

	if err := dom.OpenConsole(devName, stream, libvirt.DOMAIN_CONSOLE_SAFE); err != nil {
		return err
	}

	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	defer cond.L.Unlock()
	var quit bool

	stdin := bufio.NewReader(os.Stdin)

	// read from stream and write to stdout
	go func() {
		var err error
		for !quit {
			var buf []byte
			var got, sent int

			buf = make([]byte, 1024)

			// read from the stream, continuing if no bytes are read
			got, err = stream.Recv(buf)
			if got == 0 {
				if err != nil {
					break
				}
				continue
			}

			// write to stdout
			sent, err = os.Stdout.Write(buf)
			if sent != len(buf) {
				if err != nil {
					break
				}
			}
		}
		if err != nil {
			fmt.Println(err)
		}
		quit = true
		cond.Broadcast()
	}()

	// read from stdin and write to stream
	go func() {
		var err error
		for !quit {
			var got byte

			got, err = stdin.ReadByte()
			if err != nil {
				break
			}

			if got == escapeSequence {
				break
			}

			_, err = stream.Send([]byte{got})
			if err != nil {
				break
			}
		}
		if err != nil {
			fmt.Println(err)
		}
		quit = true
		cond.Broadcast()
	}()

	for !quit {
		cond.Wait()
	}

	return nil
}
