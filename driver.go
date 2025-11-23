package ethframe

import "runtime"

type Driver interface {
	Open(iface string) error
	Close() error
	Send(frame []byte) error
	Receive() ([]byte, error)
}

func NewDriver() Driver {
	switch runtime.GOOS {
	case "linux":
		return &linuxDriver{}
	case "darwin":
		return &darwinDriver{}
	case "freebsd", "openbsd", "netbsd":
		return &bsdDriver{}
	default:
		panic("unsupported OS: " + runtime.GOOS)
	}
}
