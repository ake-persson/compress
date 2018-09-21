package zlib

import (
	"compress/zlib"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type decoder struct {
	reader io.ReadCloser
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...compress.DecoderOption) (compress.Decoder, error) {
	d := &decoder{}
	var err error
	if d.reader, err = zlib.NewReader(r); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *decoder) SetOrder(o int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm zlib")
}

func (d *decoder) SetLitWidth(w int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm zlib")
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *decoder) Close() error {
	return d.reader.Close()
}
