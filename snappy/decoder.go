package snappy

import "github.com/golang/snappy"

type decoder struct {
	decoder *snappy.Reader
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
}

func (d *decoder) Close() error {
	return nil
}
