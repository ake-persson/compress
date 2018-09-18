package snappy

import "github.com/golang/snappy"

type encoder struct {
	encoder *snappy.Writer
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.encoder.Write(v)
}

func (e *encoder) Close() error {
	return e.encoder.Close()
}
