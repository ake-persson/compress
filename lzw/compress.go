package lzw

import (
	"compress/lzw"
	"io"

	"github.com/pkg/errors"

	"github.com/ake-persson/compress"
)

type lzwAlgorithm struct {
	endian   compress.Endian
	litWidth int
}

type lzwEncoder struct {
	writer io.WriteCloser
}

type lzwDecoder struct {
	reader io.ReadCloser
}

func (a *lzwAlgorithm) NewAlgorithm() compress.Algorithm {
	return &lzwAlgorithm{}
}

func (a *lzwAlgorithm) Ext() string {
	return "lzw"
}

func (a *lzwAlgorithm) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm lzw")
}

func (a *lzwAlgorithm) SetEndian(endian compress.Endian) error {
	a.endian = endian
	return nil
}

func (a *lzwAlgorithm) SetLitWidth(width int) error {
	a.litWidth = width
	return nil
}

func (a *lzwAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &lzwEncoder{
		writer: lzw.NewWriter(w, lzw.Order(a.endian), a.litWidth),
	}, nil
}

func (a *lzwAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(a, v)
}

func (e *lzwEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *lzwEncoder) Close() error {
	return e.writer.Close()
}

func (a *lzwAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &lzwDecoder{
		reader: lzw.NewReader(r, lzw.Order(a.endian), a.litWidth),
	}, nil
}

func (a *lzwAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(a, v)
}

func (d *lzwDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *lzwDecoder) Close() error {
	return d.reader.Close()
}

func init() {
	compress.Register("lzw", &lzwAlgorithm{})
}
