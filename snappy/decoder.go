package snappy

import "github.com/golang/snappy"

type decoder struct {
	decoder *snappy.Reader
}

func (d *decoder) Decode(v []byte) error {
	_, err := d.decoder.Read(v)
	return err
}
