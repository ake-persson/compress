package compression

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Decoder interface.
type Decoder interface {
	Decode(v []byte) (int, error)
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
/*
func FromBytes(algo string, encoded []byte, dst []byte, opts ...DecoderOption) ([]byte, error) {
	r := bufio.NewReader(bytes.NewReader(encoded))
	return nil, NewDecoder(algo, r, opts...).Decode(dst)
}
*/

// FromFile method.
func FromFile(algo string, file string, dst []byte, opts ...DecoderOption) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(fp)

	if dst != nil {
		if _, err := NewDecoder(algo, r, opts...).Decode(dst); err != nil {
			return nil, err
		}
		return nil, fp.Close()
	}

	dec := NewDecoder(algo, r, opts...)

	buf := bytes.Buffer{}
	b := make([]byte, 8)
	for {
		n, err := dec.Decode(b)
		if n > 0 {
			buf.Write(b[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), fp.Close()
}
