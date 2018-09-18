package compress

import (
	"bytes"
	"io"
)

// BufSize in bytes of read buffer.
const BufSize = 4096

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

	dec, err := a.NewDecoder(r)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(dec); err != nil {
			return nil, err
		}
	}

	return dec, nil
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
