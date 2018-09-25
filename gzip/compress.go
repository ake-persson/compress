package gzip

import (
	"compress/gzip"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

type gzipMethod struct {
	level compress.Level
}

type gzipEncoder struct {
	writer *gzip.Writer
}

type gzipDecoder struct {
	reader *gzip.Reader
}

func (m *gzipMethod) NewMethod() compress.Method {
	return &gzipMethod{}
}

func (m *gzipMethod) SetLevel(level compress.Level) error {
	m.level = level
	return nil
}

func (m *gzipMethod) SetEndian(endian compress.Endian) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (m *gzipMethod) SetLitWidth(width int) error {
	return errors.Wrap(compress.ErrUnsupportedOption, "algorithm gzip")
}

func (m *gzipMethod) NewEncoder(w io.Writer) (compress.Encoder, error) {
	e := &gzipEncoder{}
	if m.level == 0 {
		e.writer = gzip.NewWriter(w)
	} else {
		var err error
		if e.writer, err = gzip.NewWriterLevel(w, int(m.level)); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (m *gzipMethod) Encode(v []byte) ([]byte, error) {
	return compress.Encode(m, v)
}

func (e *gzipEncoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *gzipEncoder) Close() error {
	return e.writer.Close()
}

func (m *gzipMethod) NewDecoder(r io.Reader) (compress.Decoder, error) {
	d := &gzipDecoder{}
	var err error
	if d.reader, err = gzip.NewReader(r); err != nil {
		return nil, err
	}
	return d, nil
}

func (m *gzipMethod) Decode(v []byte) ([]byte, error) {
	return compress.Decode(m, v)
}

func (d *gzipDecoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *gzipDecoder) Close() error {
	return d.reader.Close()
}

func init() {
	compress.Register("gzip", &gzipMethod{})
}
