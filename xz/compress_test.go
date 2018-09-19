package xz

import (
	"bytes"
	"testing"

	"github.com/mickep76/compress"
)

func TestNewEncoder(t *testing.T) {
	if _, err := compress.NewEncoder("xz", &bytes.Buffer{}); err != nil {
		t.Error(err)
	}
}

func TestNewDecoder(t *testing.T) {
	if _, err := compress.NewDecoder("xz", &bytes.Buffer{}); err != nil {
		t.Error(err)
	}
}
