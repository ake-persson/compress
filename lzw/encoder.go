package lzw

import (
	"compress/lzw"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type encoder struct {
	writer   io.WriteCloser
	endian   int
	litWidth int
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...compress.EncoderOption) (compress.Encoder, error) {
	e := &encoder{
		litWidth: 8,
	}

	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	e.writer = lzw.NewWriter(w, lzw.Order(e.endian), e.litWidth)
	return e, nil
}

func (e *encoder) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm lzw")
}

func (e *encoder) SetEndian(endian compress.Endian) error {
	e.endian = int(endian)
	return nil
}

func (e *encoder) SetLitWidth(width int) error {
	e.litWidth = width
	return nil
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
