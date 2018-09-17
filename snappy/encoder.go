package snappy

import "github.com/golang/snappy"

type encoder struct {
	encoder *snappy.Writer
}

func (e *encoder) Encode(v []byte) (int, error) {
	return e.encoder.Write(v)
}
