package gzip

import (
	"bytes"
	"testing"

	"github.com/ake-persson/compress"
)

func TestNewDecoder(t *testing.T) {
	if _, err := compress.NewDecoder("gzip", &bytes.Buffer{}); err == nil {
		t.Error("this should have generated a EOF error")
	}
}
