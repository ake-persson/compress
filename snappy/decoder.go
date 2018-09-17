package snappy

import "github.com/golang/snappy"

type decoder struct {
	decoder *snappy.Reader
}

func (d *decoder) Decode(v []byte) (int, error) {
	return d.decoder.Read(v)
}
