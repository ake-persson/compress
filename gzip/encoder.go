package gzip

import "compress/gzip"

type encoder struct {
	encoder *gzip.Writer
}

func (e *encoder) Encode(v []byte) (int, error) {
	return e.encoder.Write(v)
}
