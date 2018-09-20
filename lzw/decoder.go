package lzw

import (
	"compress/lzw"
	"io"

	"github.com/mickep76/compress"
)

type decoder struct {
	reader   io.ReadCloser
	order    int
	litWidth int
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...compress.DecoderOption) (compress.Decoder, error) {
	e := &decoder{
		litWidth: 8,
	}

	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	e.reader = lzw.NewReader(r, lzw.Order(e.order), e.litWidth)
	return e, nil
}

func (d *decoder) SetOrder(o int) error {
	d.order = o
	return nil
}

func (d *decoder) SetLitWidth(w int) error {
	d.litWidth = w
	return nil
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *decoder) Close() error {
	return d.reader.Close()
}
