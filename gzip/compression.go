package gzip

import (
	"io"

	"compress/gzip"

	"github.com/mickep76/compression"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(writer io.Writer) (compression.Encoder, error) {
	return &encoder{encoder: gzip.NewWriter(writer)}, nil
}

func (a *algorithm) NewDecoder(reader io.Reader) (compression.Decoder, error) {
	dec, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return &decoder{decoder: dec}, nil
}

func init() {
	compression.Register("gzip", &algorithm{})
}
