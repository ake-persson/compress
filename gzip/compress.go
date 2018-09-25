package gzip

import (
	"compress/gzip"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type gzipAlgorithm struct {
	level compress.Level
}

type gzipEncoder struct {
	writer *gzip.Writer
}

type gzipDecoder struct {
	reader *gzip.Reader
}

func (a *gzipAlgorithm) NewAlgorithm() compress.Algorithm {
	return &gzipAlgorithm{}
}

func (a *gzipAlgorithm) Ext() string {
	return "gz"
}

func (a *gzipAlgorithm) SetLevel(level compress.Level) error {
	a.level = level
	return nil
}

func (a *gzipAlgorithm) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (a *gzipAlgorithm) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (a *gzipAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	e := &gzipEncoder{}
	if a.level == 0 {
		e.writer = gzip.NewWriter(w)
	} else {
		var err error
		if e.writer, err = gzip.NewWriterLevel(w, int(a.level)); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (a *gzipAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(a, v)
}

func (e *gzipEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *gzipEncoder) Close() error {
	return e.writer.Close()
}

func (a *gzipAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	d := &gzipDecoder{}
	var err error
	if d.reader, err = gzip.NewReader(r); err != nil {
		return nil, err
	}
	return d, nil
}

func (a *gzipAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(a, v)
}

func (d *gzipDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *gzipDecoder) Close() error {
	return d.reader.Close()
}

func init() {
	compress.Register("gzip", &gzipAlgorithm{})
}
