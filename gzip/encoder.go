package gzip

import "compress/gzip"

type encoder struct {
	encoder *gzip.Writer
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.encoder.Write(v)
}

func (e *encoder) Close() error {
	return e.encoder.Close()
}
