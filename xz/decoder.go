package xz

import (
	"github.com/pkg/errors"
	"io"

	"github.com/ulikunitz/xz"

	"github.com/mickep76/compress"
)

type decoder struct {
	reader *xz.Reader
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...compress.DecoderOption) (compress.Decoder, error) {
	e := &decoder{}
	var err error
	if e.reader, err = xz.NewReader(r); err != nil {
		return nil, err
	}
	return e, nil
}

func (d *decoder) SetOrder(o int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm xz")
}

func (d *decoder) SetLitWidth(w int) error {
	return errors.Wrap(compress.UnsupportedOption, "algorithm xz")
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *decoder) Close() error {
	return nil
}
