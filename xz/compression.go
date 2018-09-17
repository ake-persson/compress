package xz

import (
	"io"

	"github.com/ulikunitz/xz"

	"github.com/mickep76/compression"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(w io.Writer) (compression.Encoder, error) {
	enc, err := xz.NewWriter(w)
	if err != nil {
		return nil, err
	}
	return &encoder{encoder: enc}, nil
}

func (a *algorithm) NewDecoder(r io.Reader) (compression.Decoder, error) {
	dec, err := xz.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &decoder{decoder: dec}, nil
}

func init() {
	compression.Register("xz", &algorithm{})
}
