package compression

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

const BufSize = 4096

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
func FromBytes(algo string, encoded []byte, dst []byte, opts ...DecoderOption) ([]byte, error) {
	dec := NewDecoder(algo, bufio.NewReader(bytes.NewReader(encoded)), opts...)

	if dst != nil {
		if _, err := dec.Decode(dst); err != nil {
			return nil, err
		}
		return nil, nil
	}

	buf := bytes.Buffer{}
	b := make([]byte, BufSize)
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

	return buf.Bytes(), nil
}

// FromFile method.
func FromFile(algo string, file string, dst []byte, opts ...DecoderOption) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	dec := NewDecoder(algo, bufio.NewReader(fp), opts...)

	if dst != nil {
		if _, err := dec.Decode(dst); err != nil {
			return nil, err
		}
		return nil, fp.Close()
	}

	buf := bytes.Buffer{}
	b := make([]byte, BufSize)
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
