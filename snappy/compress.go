package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type snappyMethod struct{}

type snappyEncoder struct {
	writer *snappy.Writer
}

type snappyDecoder struct {
	reader *snappy.Reader
}

func (m *snappyMethod) NewMethod() compress.Method {
	return &snappyMethod{}
}

func (m *snappyMethod) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (m *snappyMethod) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (m *snappyMethod) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm snappy")
}

func (m *snappyMethod) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &snappyEncoder{writer: snappy.NewWriter(w)}, nil
}

func (m *snappyMethod) Encode(v []byte) ([]byte, error) {
	return compress.Encode(m, v)
}

func (e *snappyEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *snappyEncoder) Close() error {
	return e.writer.Close()
}

func (m *snappyMethod) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &snappyDecoder{reader: snappy.NewReader(r)}, nil
}

func (m *snappyMethod) Decode(v []byte) ([]byte, error) {
	return compress.Decode(m, v)
}

func (d *snappyDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *snappyDecoder) Close() error {
	return nil
}

func init() {
	compress.Register("snappy", &snappyMethod{})
}
