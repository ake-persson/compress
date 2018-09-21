package xz

import (
	"io"

	"github.com/pkg/errors"
	"github.com/ulikunitz/xz"

	"github.com/mickep76/compress"
)

type encoder struct {
	writer *xz.Writer
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...compress.EncoderOption) (compress.Encoder, error) {
	e := &encoder{}
	var err error
	if e.writer, err = xz.NewWriter(w); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *encoder) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (e *encoder) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (e *encoder) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
