package compress

import (
	"io"
	"reflect"
	"testing"
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

func TestRegistered(t *testing.T) {
	if _, err := Registered("mock"); err != nil {
		t.Error(err)
	}

	if _, err := Registered("foo"); err == nil {
		t.Error("foo reports as registered when it's not")
	}
}

func TestAlgorithms(t *testing.T) {
	if !reflect.DeepEqual(Algorithms(), []string{"mock"}) {
		t.Error("registered algorithms got unexpected response")
	}
}
