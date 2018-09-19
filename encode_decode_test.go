package compress

import (
	"bytes"
	"io"
	"testing"
)

type encoder struct {
	encoder io.Writer
}

func (e *encoder) Write(v []byte) (int, error) {
	return e.encoder.Write(v)
}

func (e *encoder) Close() error {
	return nil
}

type decoder struct {
	decoder io.Reader
}

func (d *decoder) Read(v []byte) (int, error) {
	return d.decoder.Read(v)
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
