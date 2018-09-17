package snappy

import "github.com/golang/snappy"

type encoder struct {
	encoder *snappy.Writer
}

func (e *encoder) Encode(value []byte) error {
	_, err := e.encoder.Write(value)
	return err
}
