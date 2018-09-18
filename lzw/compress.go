package lzw

import (
	"io"

	"compress/lzw"

	"github.com/mickep76/compress"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &encoder{encoder: lzw.NewWriter(w, lzw.LSB, 2)}, nil
}

func (a *algorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &decoder{decoder: lzw.NewReader(r, lzw.LSB, 2)}, nil
}

func init() {
	compress.Register("lzw", &algorithm{})
}
