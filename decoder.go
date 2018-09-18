package compress

import (
	//	"bufio"
	"bytes"
	"io"
	"os"
)

// BufSize in bytes of read buffer.
const BufSize = 4096

// Decoder interface.
type Decoder interface {
	Read(v []byte) (int, error)
	Close() error
}

// DecoderOption function.
type DecoderOption func(Decoder) error

// NewDecoder variadic constructor.
func NewDecoder(algo string, r io.Reader, opts ...DecoderOption) (Decoder, error) {
	a, err := Registered(algo)
	if err != nil {
		return nil, err
	}

	dec, err := a.NewDecoder(r)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(dec); err != nil {
			return nil, err
		}
	}

	return dec, nil
}

// Decode method.
func Decode(algo string, encoded []byte, opts ...DecoderOption) ([]byte, error) {
	buf := bytes.NewBuffer(encoded)
	dec, err := NewDecoder(algo, buf, opts...)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(os.Stdout, dec); err != nil {
		return nil, err
	}

	if err := dec.Close(); err != nil {
		return nil, err
	}

	/*
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
	*/
	return nil, nil
}
