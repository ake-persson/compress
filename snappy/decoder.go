package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type decoder struct {
	decoder *snappy.Reader
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...compress.DecoderOption) (compress.Decoder, error) {
	return &decoder{decoder: snappy.NewReader(r)}, nil
}

func (d *decoder) SetOrder(o int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (d *decoder) SetLitWidth(w int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
}

func (d *decoder) Close() error {
	return nil
}
