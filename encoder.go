package compress

import (
	"bytes"
	"fmt"
	"io"
)

type Level int

const (
	// NoCompression no compression.
	NoCompression Level = 0

	// BestSpeed best speed.
	BestSpeed Level = 1

	// BestCompression best compression.
	BestCompression Level = 9

	// DefaultCompression default compression.
	DefaultCompression Level = -1

	// HuffmanOnly huffman only.
	HuffmanOnly Level = -2
)

type Endian int

const (
	// LSB (Least Significant Bit) big endian format.
	LSB Endian = 0

	// MSB (Most Significant Bit) little endian format.
	MSB Endian = 1
)

// Encoder interface.
type Encoder interface {
	Write(v []byte) (int, error)
	Close() error
	SetOrder(o int) error
	SetLitWidth(w int) error
	SetLevel(l int) error
}

// EncoderOption variadic function.
type EncoderOption func(Encoder) error

// NewEncoder variadic constructor.
func NewEncoder(algo string, w io.Writer, opts ...EncoderOption) (Encoder, error) {
	a, ok := algorithms[algo]
	if !ok {
		return nil, fmt.Errorf("algorithm is not registered: %s", algo)
	}

	return a.NewEncoder(w, opts...)
}

// WithLitWidth the number of bit's to use for literal codes.
// Supported by lzw.
func WithLitWidth(w int) EncoderOption {
	return func(e Encoder) error {
		return e.SetLitWidth(w)
	}
}

// WithOrder either MSB (most significant byte) or LSB (least significant byte).
// Supported by lzw.
func WithOrder(o Endian) EncoderOption {
	return func(e Encoder) error {
		return e.SetOrder(int(o))
	}
}

// WithLevel compression level.
// Supported by gzip, zlib.
func WithLevel(l Level) EncoderOption {
	return func(e Encoder) error {
		return e.SetLevel(int(l))
	}
}

// Encode method.
func Encode(algo string, v []byte, opts ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer

	enc, err := NewEncoder(algo, &buf, opts...)
	if err != nil {
		return nil, err
	}

	if _, err := enc.Write(v); err != nil {
		return nil, err
	}

	if err := enc.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
