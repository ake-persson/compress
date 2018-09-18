package gzip

import (
	"io"

	"compress/gzip"

	"github.com/mickep76/compress"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &encoder{encoder: gzip.NewWriter(w)}, nil
}

func (a *algorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	dec, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &decoder{decoder: dec}, nil
}

func init() {
	compress.Register("gzip", &algorithm{})
}