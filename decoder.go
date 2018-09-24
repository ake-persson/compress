package compress

import (
	"bytes"
	"io"
)

// Decoder interface.
type Decoder interface {
	Read(v []byte) (int, error)
	Close() error
	SetOrder(o int) error // Should be SetEndian
	SetLitWidth(w int) error
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

// WithLitWidth the number of bit's to use for literal codes.
// Supported by lzw.
/*
func WithLitWidth(width int) DecoderOption {
	return func(d Decoder) error {
		return d.SetLitWidth(width)
	}
}
*/

// WithEndian either MSB (most significant byte) or LSB (least significant byte).
// Supported by lzw.
/*
func WithEndian(endian Endian) DecoderOption {
	return func(d Decoder) error {
		return d.SetEndian(endian)
	}
}
*/

// Decode method.
func Decode(name string, encoded []byte, opts ...DecoderOption) ([]byte, error) {
	d, err := NewDecoder(name, bytes.NewBuffer(encoded), opts...)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, d); err != nil {
		return nil, err
	}

	if err := d.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
