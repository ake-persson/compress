package xz

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ake-persson/compress"
)

type MockWriter struct{}

func (m *MockWriter) Write(v []byte) (int, error) {
	return 0, errors.New("failed")
}

func TestNewEncoder(t *testing.T) {
	if _, err := compress.NewEncoder("xz", bytes.NewBuffer([]byte(""))); err != nil {
		t.Error(err)
	}

	if _, err := compress.NewEncoder("xz", &MockWriter{}); err == nil {
		t.Error("this should not fail")
	}
}

func TestNewDecoder(t *testing.T) {
	if _, err := compress.NewDecoder("gzip", &bytes.Buffer{}); err == nil {
		t.Error("this should have generated a EOF error")
	}
}
