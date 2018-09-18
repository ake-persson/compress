package gzip

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"

	"github.com/mickep76/compress"
)

func TestEncodeDecode(t *testing.T) {
	exp := []byte("abc123\ndef456\nabc123\ndef456\nabc123\ndef456\n")

	encoded, err := compress.Encode("gzip", exp)
	if err != nil {
		t.Error(errors.Wrap(err, "test encode"))
	}

	if got, err := compress.Decode("gzip", encoded); err != nil {
		t.Error(errors.Wrap(err, "test decode"))
	} else if !bytes.Equal(exp, got) {
		t.Error(errors.Errorf("test decode doesn't match expected value"))
	}
}
