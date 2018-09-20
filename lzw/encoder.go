package lzw

import (
	"compress/lzw"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type encoder struct {
	writer   io.WriteCloser
	order    int
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

	e.writer = lzw.NewWriter(w, lzw.Order(e.order), e.litWidth)
	return e, nil
}

func (e *encoder) SetLevel(o int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm lzw")
}

func (e *encoder) SetOrder(o int) error {
	e.order = o
	return nil
}

func (e *encoder) SetLitWidth(w int) error {
	e.litWidth = w
	return nil
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
