package compression

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Encoder interface.
type Encoder interface {
	Encode(v []byte) error
}

// EncoderOption variadic function.
type EncoderOption func(Encoder)

// NewEncoder variadic constructor.
func NewEncoder(algorithm string, writer io.Writer, options ...EncoderOption) Encoder {
	c, ok := algorithms[algorithm]
	if !ok {
		return nil
	}

	enc := c.NewEncoder(writer)
	for _, option := range options {
		option(enc)
	}

	return enc
}

// ToBytes method.
func ToBytes(algorithm string, value []byte, options ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer

	if err := NewEncoder(algorithm, &buf, options...).Encode(value); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToFile method.
func ToFile(algorithm string, file string, value []byte, options ...EncoderOption) error {
	fp, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(fp)

	if err := NewEncoder(algorithm, w, options...).Encode(value); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return fp.Close()
}
