package compress

import (
	"bytes"
	"io"
	"testing"

	"github.com/pkg/errors"
)

type algorithm struct{}

func (a *algorithm) NewEncoder(w io.Writer) (Encoder, error) {
	return &encoder{encoder: w}, nil
}

func (a *algorithm) NewDecoder(r io.Reader) (Decoder, error) {
	return &decoder{decoder: r}, nil
}

func init() {
	Register("mock", &algorithm{})
}

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
		t.Error(errors.Wrap(err, "test encode"))
	}

	if got, err := Decode("mock", encoded); err != nil {
		t.Error(errors.Wrap(err, "test decode"))
	} else if !bytes.Equal(exp, got) {
		t.Error(errors.Errorf("test decode doesn't match expected value"))
	}
}
