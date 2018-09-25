package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type snappyAlgorithm struct{}

type snappyEncoder struct {
	writer *snappy.Writer
}

type snappyDecoder struct {
	reader *snappy.Reader
}

func (a *snappyAlgorithm) NewAlgorithm() compress.Algorithm {
	return &snappyAlgorithm{}
}

func (a *snappyAlgorithm) Ext() string {
	return "snappy"
}

func (a *snappyAlgorithm) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algoritha *snappy")
}

func (a *snappyAlgorithm) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algoritha *snappy")
}

func (a *snappyAlgorithm) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algoritha *snappy")
}

func (a *snappyAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	return &snappyEncoder{writer: snappy.NewWriter(w)}, nil
}

func (a *snappyAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(a, v)
}

func (e *snappyEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *snappyEncoder) Close() error {
	return e.writer.Close()
}

func (a *snappyAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	return &snappyDecoder{reader: snappy.NewReader(r)}, nil
}

func (a *snappyAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(a, v)
}

func (d *snappyDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *snappyDecoder) Close() error {
	return nil
}

func init() {
	compress.Register("snappy", &snappyAlgorithm{})
}
