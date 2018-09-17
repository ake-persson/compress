package snappy

import "github.com/golang/snappy"

type decoder struct {
	decoder *snappy.Reader
}

func (d *decoder) Decode(v []byte) (int, error) {
	n, err := d.decoder.Read(v)
	return n, err
}
