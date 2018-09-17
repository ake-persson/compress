package compression

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Decoder interface.
type Decoder interface {
	Decode(v []byte) error
}

// DecoderOption function.
type DecoderOption func(Decoder)

// NewDecoder variadic constructor.
func NewDecoder(algorithm string, reader io.Reader, options ...DecoderOption) Decoder {
	c, ok := algorithms[algorithm]
	if !ok {
		return nil
	}

	dec := c.NewDecoder(reader)
	for _, option := range options {
		option(dec)
	}

	return dec
}

// FromBytes method.
func FromBytes(algorithm string, encoded []byte, value []byte, options ...DecoderOption) error {
	r := bufio.NewReader(bytes.NewReader(encoded))
	return NewDecoder(algorithm, r, options...).Decode(value)
}

// FromFile method.
func FromFile(algorithm string, file string, value []byte, options ...DecoderOption) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}

	r := bufio.NewReader(fp)

	if err := NewDecoder(algorithm, r, options...).Decode(value); err != nil {
		return err
	}

	return fp.Close()
}
