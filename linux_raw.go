package ethframe

import (
	"net"
	"syscall"
)

type linuxDriver struct {
	fd int
}

func (d *linuxDriver) Open(iface string) error {
	netIface, err := net.InterfaceByName(iface)
	if err != nil {
		return err
	}

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(0x0003)))
	if err != nil {
		return err
	}
	d.fd = fd

	sll := syscall.SockaddrLinklayer{
		Protocol: htons(0x0003),
		Ifindex:  netIface.Index,
	}
	return syscall.Bind(fd, &sll)
}

func (d *linuxDriver) Send(frame []byte) error {
	return syscall.Sendto(d.fd, frame, 0, nil)
}

func (d *linuxDriver) Receive() ([]byte, error) {
	buf := make([]byte, 2048)
	n, _, err := syscall.Recvfrom(d.fd, buf, 0)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (d *linuxDriver) Close() error {
	return syscall.Close(d.fd)
}

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}
