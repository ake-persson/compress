package snappy

import (
	"io"

	"github.com/golang/snappy"

	"github.com/mickep76/compression"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(writer io.Writer) compression.Encoder {
	return &encoder{encoder: snappy.NewWriter(writer)}
}

func (a *algorithm) NewDecoder(reader io.Reader) compression.Decoder {
	return &decoder{decoder: snappy.NewReader(reader)}
}

func init() {
	compression.Register("snappy", &algorithm{})
}
