package zlib

import (
	"compress/zlib"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type zlibAlgorithm struct {
	level compress.Level
}

type zlibEncoder struct {
	writer io.WriteCloser
}

type zlibDecoder struct {
	reader io.ReadCloser
}

func (a *zlibAlgorithm) NewAlgorithm() compress.Algorithm {
	return &zlibAlgorithm{}
}

func (a *zlibAlgorithm) SetLevel(level compress.Level) error {
	a.level = level
	return nil
}

func (a *zlibAlgorithm) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (a *zlibAlgorithm) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (a *zlibAlgorithm) NewEncoder(w io.Writer) (compress.Encoder, error) {
	e := &zlibEncoder{}
	if a.level == 0 {
		e.writer = zlib.NewWriter(w)
	} else {
		var err error
		if e.writer, err = zlib.NewWriterLevel(w, int(a.level)); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (a *zlibAlgorithm) Encode(v []byte) ([]byte, error) {
	return compress.Encode(a, v)
}

func (e *zlibEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *zlibEncoder) Close() error {
	return e.writer.Close()
}

func (a *zlibAlgorithm) NewDecoder(r io.Reader) (compress.Decoder, error) {
	d := &zlibDecoder{}
	var err error
	if d.reader, err = zlib.NewReader(r); err != nil {
		return nil, err
	}
	return d, nil
}

func (a *zlibAlgorithm) Decode(v []byte) ([]byte, error) {
	return compress.Decode(a, v)
}

func (d *zlibDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *zlibDecoder) Close() error {
	return d.reader.Close()
}

func init() {
	compress.Register("zlib", &zlibAlgorithm{})
}
