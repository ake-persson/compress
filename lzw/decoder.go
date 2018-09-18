package lzw

import "io"

type decoder struct {
	decoder io.ReadCloser
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
}

func (d *decoder) Close() error {
	return d.decoder.Close()
}
