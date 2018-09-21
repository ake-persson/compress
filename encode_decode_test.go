package compress

import (
	"bytes"
	"io"
	"testing"
)

type encoder struct {
	writer io.Writer
}

func (a *algorithm) NewEncoder(w io.Writer, opts ...EncoderOption) (Encoder, error) {
	return &encoder{writer: w}, nil
}

func (e *encoder) SetEndian(endian Endian) error {
	return nil
}

func (e *encoder) SetLitWidth(width int) error {
	return nil
}

func (e *encoder) SetLevel(level Level) error {
	return nil
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.writer.Write(v)
}

func (e *encoder) Close() error {
	return nil
}

type decoder struct {
	reader io.Reader
}

func (a *algorithm) NewDecoder(r io.Reader, opts ...DecoderOption) (Decoder, error) {
	return &decoder{reader: r}, nil
}

func (d *decoder) SetOrder(o int) error {
	return nil
}

func (d *decoder) SetLitWidth(w int) error {
	return nil
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.reader.Read(v)
}

func (d *decoder) Close() error {
	return nil
}

func TestEncodeDecode(t *testing.T) {
	exp := []byte("abc123\ndef456\nabc123\ndef456\nabc123\ndef456\n")

	encoded, err := Encode("mock", exp)
	if err != nil {
		t.Error(err)
	}

	if got, err := Decode("mock", encoded); err != nil {
		t.Error(err)
	} else if !bytes.Equal(exp, got) {
		t.Error("decode response doesn't match what was encoded")
	}
}

func TestNewEncoder(t *testing.T) {
	if _, err := NewDecoder("foo", bytes.NewBuffer([]byte(""))); err == nil {
		t.Error("foo should not be a registered algorithm")
	}
}
