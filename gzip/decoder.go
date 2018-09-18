package gzip

import "compress/gzip"

type decoder struct {
	decoder *gzip.Reader
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
}

func (d *decoder) Close() error {
	return d.decoder.Close()
}
