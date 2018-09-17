package compression

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Encoder interface.
type Encoder interface {
	Encode(v []byte) (int, error)
}

// EncoderOption variadic function.
type EncoderOption func(Encoder)

// NewEncoder variadic constructor.
func NewEncoder(algorithm string, w io.Writer, opts ...EncoderOption) Encoder {
	c, ok := algorithms[algorithm]
	if !ok {
		return nil
	}

	enc := c.NewEncoder(w)
	for _, opt := range opts {
		opt(enc)
	}

	return enc
}

// ToBytes method.
func ToBytes(algo string, v []byte, opts ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer

	if _, err := NewEncoder(algo, &buf, opts...).Encode(v); err != nil {
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

	if _, err := NewEncoder(algo, w, opts...).Encode(v); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return fp.Close()
}
