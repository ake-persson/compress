package compress

import (
	"bytes"
	"fmt"
	"io"
)

var methods = make(map[string]Method)

// Method interface.
type Method interface {
	NewEncoder(w io.Writer) (Encoder, error)
	NewDecoder(r io.Reader) (Decoder, error)
	Encode(v []byte) ([]byte, error)
	Decode(v []byte) ([]byte, error)
	SetLevel(level Level) error
	SetLitWidth(width int) error
	SetEndian(endian Endian) error
}

// Encoder interface.
type Encoder interface {
	Write(v []byte) (int, error)
	Close() error
}

// Decoder interface.
type Decoder interface {
	Read(v []byte) (int, error)
	Close() error
}

// Option variadic function.
type Option func(Method) error

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

// Endian the order in which bytes are arranged into larger values.
type Endian int

const (
	// Little endian LSB (Least Significant Bit).
	Little Endian = 0

	// Big endian MSB (Most Significant Bit).
	Big Endian = 1
)

// Register method.
func Register(name string, method Method) {
	methods[name] = method
}

// Methods registered.
func Methods() []string {
	l := []string{}
	for a := range methods {
		l = append(l, a)
	}
	return l
}

// NewMethod constructor.
func NewMethod(name string, opts ...Option) (Method, error) {
	m, ok := methods[name]
	if !ok {
		return nil, fmt.Errorf("method not registered: %s", name)
	}
	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}
	return m, nil
}

// WithLevel compression level.
// Supported by gzip, zlib.
func WithLevel(level Level) Option {
	return func(m Method) error {
		return m.SetLevel(level)
	}
}

// WithLitWidth the number of bit's to use for literal codes.
// Supported by lzw.
func WithLitWidth(width int) Option {
	return func(m Method) error {
		return m.SetLitWidth(width)
	}
}

// WithEndian either MSB (most significant byte) or LSB (least significant byte).
// Supported by lzw.
func WithEndian(endian Endian) Option {
	return func(m Method) error {
		return m.SetEndian(endian)
	}
}

// Encode method.
func Encode(m Method, v []byte) ([]byte, error) {
	var buf bytes.Buffer
	e, err := m.NewEncoder(&buf)
	if err != nil {
		return nil, err
	}

	if _, err := e.Write(v); err != nil {
		return nil, err
	}

	if err := e.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode method.
func Decode(m Method, v []byte) ([]byte, error) {
	d, err := m.NewDecoder(bytes.NewBuffer(v))
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
