package gzip

import "compress/gzip"

type decoder struct {
	decoder *gzip.Reader
}

func (d *decoder) Decode(v []byte) (int, error) {
	return d.decoder.Read(v)
}
