package xz

import (
	"io"

	"github.com/pkg/errors"
	"github.com/ulikunitz/xz"

	"github.com/ake-persson/compress"
)

type xzAlgorithm struct{}

type xzEncoder struct {
	writer *xz.Writer
}

type xzDecoder struct {
	reader *xz.Reader
}

func (a *xzAlgorithm) NewAlgorithm() compress.Algorithm {
	return &xzAlgorithm{}
}

func (a *xzAlgorithm) Ext() string {
	return "xz"
}

func (a *xzAlgorithm) SetLevel(level compress.Level) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (a *xzAlgorithm) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (a *xzAlgorithm) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm xz")
}

func (a *xzAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	e := &xzEncoder{}
	var err error
	if e.writer, err = xz.NewWriter(w); err != nil {
		return nil, err
	}
	return e, nil
}

func (a *xzAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(a, v)
}

func (e *xzEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *xzEncoder) Close() error {
	return e.writer.Close()
}

func (a *xzAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	e := &xzDecoder{}
	var err error
	if e.reader, err = xz.NewReader(r); err != nil {
		return nil, err
	}
	return e, nil
}

func (a *xzAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(a, v)
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
