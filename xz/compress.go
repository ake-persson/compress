package xz

import (
	"io"

	"github.com/pkg/errors"
	"github.com/ulikunitz/xz"

	"github.com/mickep76/compress"
)

type xzAlgorithm struct{}

type xzEncoder struct {
	writer *xz.Writer
}

type xzDecoder struct {
	reader *xz.Reader
}

func (m *xzAlgorithm) NewAlgorithm() compress.Algorithm {
	return &xzAlgorithm{}
}

func (m *xzAlgorithm) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (m *xzAlgorithm) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (m *xzAlgorithm) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (m *xzAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	e := &xzEncoder{}
	var err error
	if e.writer, err = xz.NewWriter(w); err != nil {
		return nil, err
	}
	return e, nil
}

func (m *xzAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(m, v)
}

func (e *xzEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *xzEncoder) Close() error {
	return e.writer.Close()
}

func (m *xzAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	e := &xzDecoder{}
	var err error
	if e.reader, err = xz.NewReader(r); err != nil {
		return nil, err
	}
	return e, nil
}

func (m *xzAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(m, v)
}

func (d *xzDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *xzDecoder) Close() error {
	return nil
}

func init() {
	compress.Register("xz", &xzAlgorithm{})
}
