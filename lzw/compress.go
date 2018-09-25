package lzw

import (
	"compress/lzw"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type lzwMethod struct {
	endian   compress.Endian
	litWidth int
}

type lzwEncoder struct {
	writer io.WriteCloser
}

type lzwDecoder struct {
	reader io.ReadCloser
}

func (m *lzwMethod) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (m *lzwMethod) SetEndian(endian compress.Endian) error {
	m.endian = endian
	return nil
}

func (m *lzwMethod) SetLitWidth(width int) error {
	m.litWidth = width
	return nil
}

func (m *lzwMethod) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &lzwEncoder{
		writer: lzw.NewWriter(w, lzw.Order(m.endian), m.litWidth),
	}, nil
}

func (m *lzwMethod) Encode(v []byte) ([]byte, error) {
	return compress.Encode(m, v)
}

func (e *lzwEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *lzwEncoder) Close() error {
	return e.writer.Close()
}

func (m *lzwMethod) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &lzwDecoder{
		reader: lzw.NewReader(r, lzw.Order(m.endian), m.litWidth),
	}, nil
}

func (m *lzwMethod) Decode(v []byte) ([]byte, error) {
	return compress.Decode(m, v)
}

func (d *lzwDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *lzwDecoder) Close() error {
	return d.reader.Close()
}

func init() {
	compress.Register("lzw", &lzwMethod{})
}
