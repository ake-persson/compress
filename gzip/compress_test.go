package gzip

import (
	"bytes"
	"testing"

	"github.com/mickep76/compress"
)

func TestNewDecoder(t *testing.T) {
	if _, err := compress.NewDecoder("gzip", &bytes.Buffer{}); err != nil {
		t.Error(err)
	}
}
