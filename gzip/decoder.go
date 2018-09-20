package gzip

import (
	"compress/gzip"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type decoder struct {
	reader *gzip.Reader
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...compress.DecoderOption) (compress.Decoder, error) {
	d := &decoder{}
	var err error
	if d.reader, err = gzip.NewReader(r); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *decoder) SetOrder(o int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm gzip")
}

func (d *decoder) SetLitWidth(w int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm gzip")
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *decoder) Close() error {
	return d.reader.Close()
}
