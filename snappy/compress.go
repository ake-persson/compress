package snappy

import (
	"io"

	"github.com/golang/snappy"

	"github.com/mickep76/compress"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &encoder{encoder: snappy.NewWriter(w)}, nil
}

func (a *algorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &decoder{decoder: snappy.NewReader(r)}, nil
}

func init() {
	compress.Register("snappy", &algorithm{})
}
