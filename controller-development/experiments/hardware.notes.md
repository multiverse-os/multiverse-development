


## BusPirate 
https://github.com/jeffbuttars/gobuspirate



## LCDs 
Needed for laptop, phone and tablet designs 
https://github.com/PaulB2Code/lcd-1602-uart-i2c/blob/master/display.go
## Connections 


https://github.com/google/periph/tree/master/conn 
**This is one of the most important packages** it has all the most important
connection drivers. Which could be implemented pretty eaisly but these are
implemented in a beautiful way. 
 	conntest 	Reduce number of errcheck lint messages from 117 to 2. 	2 years ago
	display 	display: add missing package doc 	last year
	gpio 	gpiotest: futureproof for v4 	3 months ago
	i2c 	Reduce the number of golint comments. 	5 months ago
	i2s 	Reduce the number of golint comments. 	5 months ago
	ir 	Rename pio to periph. (#27) 	3 years ago
	jtag 	Reduce the number of golint comments. 	5 months ago
	mmr 	Stop using t.Fail() and FailNow 	11 months ago
	onewire 	Reduce the number of golint comments. 	5 months ago
	physic 	travis: enable more of go vet shadow check 	4 months ago
	pin 	Stop using t.Fail() and FailNow 	11 months ago
	spi 	Reduce the number of golint comments. 	5 months ago
	uart 	Reduce the number of golint comments. 	5 months ago
	conn.go 	conn: Resource and Conn now requires String() 	last year
	conn_test.go 	I was so happy with recent Example refactors, did it for all packages 	2 years ago
	doc.go 	conn: fix incorrect reference to smoketest in package doc 	last year
	duplex_string.go 	Add conn.Conn.Duplex() (#118) 	3 years ago
	example_test.go

https://github.com/google/periph
    https://github.com/google/periph/blob/master/experimental/conn/uart/uart.go
    **These are written so fucking well** it almost makes everything else just
    seem bad. The structure is fantstic. It outlines exactly how a connection
    driver should be written. and enables basically for example supporting all
    wire typeps.  
      Could still use HDMI CEC, Display cable, And possible lazer/light for FIOS 



https://github.com/djthorpe/gopi 
This repository contains an application framework for the Go language, which will allow you to develop applications which utilize a number of features of your computer. It's targetted at the Raspberry Pi presently. The following features are intended to be supported:
    The GPIO, I2C and SPI interfaces
    Display and display surfaces, bitmaps and vector graphics
    GPU acceleration for 2D graphics
    Font loading and rendering in bitmap and vector forms
    Input devices like the mouse, keyboard and touchscreen
    Infrared transmission and receiving, for example for remote controls
    Network microservices, announcement and discovery



const isLinux = true

// connSocket is a simple wrapper around a Linux netlink connector socket.
type connSocket struct {
	fd int
}

// newConnSocket returns a socket instance.
func newConnSocket() (*connSocket, error) {
	// Open netlink socket.
	fd, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, syscall.NETLINK_CONNECTOR)
	if err != nil {
		return nil, fmt.Errorf("failed to open netlink socket: %v", err)
	}

	if err := syscall.Bind(fd, &syscall.SockaddrNetlink{Family: syscall.AF_NETLINK}); err != nil {
		return nil, fmt.Errorf("failed to bind netlink socket: %v", err)
	}

	return &connSocket{fd: fd}, nil
}

// send writes w to the socket.
func (s *connSocket) send(w []byte) error {
	return syscall.Sendto(s.fd, w, 0, &syscall.SockaddrNetlink{Family: syscall.AF_NETLINK})
}

// recv reads at most len(r) bytes from the socket into r. Returns the actually
// read number of bytes.
func (s *connSocket) recv(r []byte) (int, error) {
	n, _, err := syscall.Recvfrom(s.fd, r, 0)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// close closes the socket.
func (s *connSocket) close() error {
	fd := s.fd
	s.fd = 0
	return syscall.Close(fd)
}
