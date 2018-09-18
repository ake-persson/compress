package zlib

import "io"

type encoder struct {
	encoder io.WriteCloser
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.encoder.Write(v)
}

func (e *encoder) Close() error {
	return e.encoder.Close()
}
