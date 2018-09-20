package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type encoder struct {
	writer *snappy.Writer
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...compress.EncoderOption) (compress.Encoder, error) {
	return &encoder{writer: snappy.NewWriter(w)}, nil
}

func (e *encoder) SetOrder(o int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm snappy")
}

func (e *encoder) SetLitWidth(w int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm snappy")
}

func (e *encoder) SetLevel(l int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm snappy")
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
