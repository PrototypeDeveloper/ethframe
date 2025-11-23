package ethframe

import (
	"os"
)

type bsdDriver struct {
	f *os.File
}

func (d *bsdDriver) Open(iface string) error {
	f, err := os.OpenFile("/dev/bpf", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	d.f = f
	return nil
}

func (d *bsdDriver) Send(frame []byte) error {
	_, err := d.f.Write(frame)
	return err
}

func (d *bsdDriver) Receive() ([]byte, error) {
	buf := make([]byte, 2048)
	n, err := d.f.Read(buf)
	return buf[:n], err
}

func (d *bsdDriver) Close() error {
	return d.f.Close()
}
