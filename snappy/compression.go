package snappy

import (
	"io"

	"github.com/golang/snappy"

	"github.com/mickep76/compression"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(writer io.Writer) (compression.Encoder, error) {
	return &encoder{encoder: snappy.NewWriter(writer)}, nil
}

func (a *algorithm) NewDecoder(reader io.Reader) (compression.Decoder, error) {
	return &decoder{decoder: snappy.NewReader(reader)}, nil
}

func init() {
	compression.Register("snappy", &algorithm{})
}
