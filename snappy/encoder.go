package snappy

import "github.com/golang/snappy"

type encoder struct {
	encoder *snappy.Writer
}

func (e *encoder) Encode(v []byte) error {
	_, err := e.encoder.Write(v)
	return err
}
