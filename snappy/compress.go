package snappy

import (
	"io"

	"github.com/golang/snappy"

	"github.com/mickep76/compress"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(writer io.Writer) (compress.Encoder, error) {
	return &encoder{encoder: snappy.NewWriter(writer)}, nil
}

func (a *algorithm) NewDecoder(reader io.Reader) (compress.Decoder, error) {
	return &decoder{decoder: snappy.NewReader(reader)}, nil
}

func init() {
	compress.Register("snappy", &algorithm{})
}
