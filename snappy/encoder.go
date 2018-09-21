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

func (e *encoder) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (e *encoder) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (e *encoder) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
