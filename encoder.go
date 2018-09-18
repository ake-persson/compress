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
}

// EncoderOption variadic function.
type EncoderOption func(Encoder) error

// NewEncoder variadic constructor.
func NewEncoder(algo string, w io.Writer, opts ...EncoderOption) (Encoder, error) {
	a, ok := algorithms[algo]
	if !ok {
		return nil, fmt.Errorf("algorithm is not registered: %s", algo)
	}

	enc, err := a.NewEncoder(w)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(enc); err != nil {
			return nil, err
		}
	}

	return enc, nil
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
