package compress

import (
	"bytes"
	"io"
)

// Decoder interface.
type Decoder interface {
	Read(v []byte) (int, error)
	Close() error
}

// DecoderOption function.
type DecoderOption func(Decoder) error

// NewDecoder variadic constructor.
func NewDecoder(algo string, r io.Reader, opts ...DecoderOption) (Decoder, error) {
	a, err := Registered(algo)
	if err != nil {
		return nil, err
	}

	return a.NewDecoder(r, opts...)
}

// Decode method.
func Decode(algo string, encoded []byte, opts ...DecoderOption) ([]byte, error) {
	dec, err := NewDecoder(algo, bytes.NewBuffer(encoded), opts...)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, dec); err != nil {
		return nil, err
	}

	if err := dec.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
