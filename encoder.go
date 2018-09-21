package compress

import (
	"bytes"
	"fmt"
	"io"
)

// Level compression level.
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

// Endian little (Least Significant Bit) or big endian (Most Significant Bit).
type Endian int

const (
	// Little endian LSB (Least Significant Bit).
	Little Endian = 0

	// Big endian MSB (Most Significant Bit).
	Big Endian = 1
)

// Encoder interface.
type Encoder interface {
	Write(v []byte) (int, error)
	Close() error
	SetEndian(endian Endian) error
	SetLitWidth(width int) error
	SetLevel(level Level) error
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
func WithLitWidth(width int) EncoderOption {
	return func(e Encoder) error {
		return e.SetLitWidth(width)
	}
}

// WithEndian either MSB (most significant byte) or LSB (least significant byte).
// Supported by lzw.
func WithEndian(endian Endian) EncoderOption {
	return func(e Encoder) error {
		return e.SetEndian(endian)
	}
}

// WithLevel compression level.
// Supported by gzip, zlib.
func WithLevel(level Level) EncoderOption {
	return func(e Encoder) error {
		return e.SetLevel(level)
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
