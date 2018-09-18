package xz

import "github.com/ulikunitz/xz"

type decoder struct {
	decoder *xz.Reader
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
}

func (d *decoder) Close() error {
	return nil
}
