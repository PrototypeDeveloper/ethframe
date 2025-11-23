package ethframe

import (
	"os"
)

type darwinDriver struct {
	f *os.File
}

func (d *darwinDriver) Open(iface string) error {
	f, err := os.OpenFile("/dev/bpf", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	d.f = f
	return nil
}

func (d *darwinDriver) Send(frame []byte) error {
	_, err := d.f.Write(frame)
	return err
}

func (d *darwinDriver) Receive() ([]byte, error) {
	buf := make([]byte, 2048)
	n, err := d.f.Read(buf)
	return buf[:n], err
}

func (d *darwinDriver) Close() error {
	return d.f.Close()
}
