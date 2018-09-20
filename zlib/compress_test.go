package zlib

import (
	"bytes"
	"testing"

	"github.com/mickep76/compress"
)

func TestNewDecoder(t *testing.T) {
	if _, err := compress.NewDecoder("zlib", &bytes.Buffer{}); err == nil {
		t.Error("this should have generated a EOF error")
	}
}
