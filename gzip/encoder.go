package gzip

import (
	"compress/gzip"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type encoder struct {
	writer *gzip.Writer
	level  int
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...compress.EncoderOption) (compress.Encoder, error) {
	e := &encoder{}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	if e.level == 0 {
		e.writer = gzip.NewWriter(w)
	} else {
		var err error
		if e.writer, err = gzip.NewWriterLevel(w, e.level); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func (e *encoder) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (e *encoder) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (e *encoder) SetLevel(level compress.Level) error {
	e.level = int(level)
	return nil
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
