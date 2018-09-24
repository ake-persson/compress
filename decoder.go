package compress

import (
	"bytes"
	"io"
)

// Decoder interface.
type Decoder interface {
	Read(v []byte) (int, error)
	Close() error
}

// Decode method.
func (m *Method) Decode(encoded []byte) ([]byte, error) {
	d, err := NewDecoder(bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, d); err != nil {
		return nil, err
	}

	if err := d.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
