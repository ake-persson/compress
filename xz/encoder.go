package xz

import "github.com/ulikunitz/xz"

type encoder struct {
	encoder *xz.Writer
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.encoder.Write(v)
}

func (e *encoder) Close() error {
	return e.encoder.Close()
}
