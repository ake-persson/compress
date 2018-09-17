package compression

import (
	"bufio"
	"bytes"
	//	"fmt"
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
func NewDecoder(algo string, r io.Reader, opts ...DecoderOption) Decoder {
	c, ok := algorithms[algo]
	if !ok {
		return nil
	}

	dec := c.NewDecoder(r)
	for _, opt := range opts {
		opt(dec)
	}

	return dec
}

// FromBytes method.
func FromBytes(algo string, encoded []byte, v []byte, opts ...DecoderOption) error {
	r := bufio.NewReader(bytes.NewReader(encoded))
	return NewDecoder(algo, r, opts...).Decode(v)
}

// FromFile method.
func FromFile(algo string, file string, v []byte, opts ...DecoderOption) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}

	r := bufio.NewReader(fp)

	if err := NewDecoder(algo, r, opts...).Decode(v); err != nil {
		return err
	}

	return fp.Close()
}
