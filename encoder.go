package compress

import (
	"bytes"
	"fmt"
	"io"
)

// Encoder interface.
type Encoder interface {
	Write(v []byte) (int, error)
	Close() error
	SetLevel(l int) error
	SetOrder(o int) error
	SetLitWidth(w int) error
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

// WithLevel compression level.
func WithLevel(level int) EncoderOption {
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
