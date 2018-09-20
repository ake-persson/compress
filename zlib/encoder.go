package zlib

import (
	"compress/zlib"
	"io"

	"github.com/mickep76/compress"
)

const (
	ZlibNoCompression      = 0
	ZlibBestSpeed          = 1
	ZlibBestCompression    = 9
	ZlibDefaultCompression = -1
	ZlibHuffmanOnly        = -2
)

type encoder struct {
	writer io.WriteCloser
	level  int
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...compress.EncoderOption) (compress.Encoder, error) {
	e := &encoder{}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	if e.level == 0 {
		e.writer = zlib.NewWriter(w)
	} else {
		var err error
		if e.writer, err = zlib.NewWriterLevel(w, e.level); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func (e *encoder) SetLevel(l int) error {
	e.level = l
	return nil
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return e.writer.Close()
}
