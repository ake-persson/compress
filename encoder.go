package compression

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// Encoder interface.
type Encoder interface {
	Encode(v []byte) (int, error)
}

// EncoderOption variadic function.
type EncoderOption func(Encoder) error

// NewEncoder variadic constructor.
func NewEncoder(algo string, w io.Writer, opts ...EncoderOption) (Encoder, error) {
	c, ok := algorithms[algo]
	if !ok {
		return nil, fmt.Errorf("algorithm is not registered: %s", algo)
	}

	enc, err := c.NewEncoder(w)
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

// ToBytes method.
func ToBytes(algo string, v []byte, opts ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer

	enc, err := NewEncoder(algo, &buf, opts...)
	if err != nil {
		return nil, err
	}

	if _, err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToFile method.
func ToFile(algo string, file string, v []byte, opts ...EncoderOption) error {
	fp, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(fp)

	enc, err := NewEncoder(algo, w, opts...)
	if err != nil {
		return err
	}

	if _, err := enc.Encode(v); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return fp.Close()
}
