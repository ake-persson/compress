package compress

import (
	"bytes"
	"io"
)

// Encoder interface.
type Encoder interface {
	Write(v []byte) (int, error)
	Close() error
}

// Encode method.
func (m *Method) Encode(v []byte) ([]byte, error) {
	var buf bytes.Buffer
	e, err := m.NewEncoder(&buf)
	if err != nil {
		return nil, err
	}

	if _, err := e.Write(v); err != nil {
		return nil, err
	}

	if err := e.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
