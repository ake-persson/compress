package xz

import "github.com/ulikunitz/xz"

type encoder struct {
	encoder *xz.Writer
}

func (e *encoder) Encode(v []byte) (int, error) {
	return e.encoder.Write(v)
}
