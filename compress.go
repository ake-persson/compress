package compress

import (
	"io"

	"github.com/pkg/errors"
)

var algorithms = make(map[string]Algorithm)

// Algorithm interface.
type Algorithm interface {
	NewEncoder(w io.Writer, opts ...EncoderOption) (Encoder, error)
	NewDecoder(r io.Reader, opts ...DecoderOption) (Decoder, error)
}

// Register algorithm.
func Register(name string, algorithm Algorithm) {
	algorithms[name] = algorithm
}

// Registered algorithm.
func Registered(name string) (Algorithm, error) {
	a, ok := algorithms[name]
	if !ok {
		return nil, errors.Errorf("algorithm not registered: %s", name)
	}
	return a, nil
}

// Algorithms registered.
func Algorithms() []string {
	l := []string{}
	for a := range algorithms {
		l = append(l, a)
	}
	return l
}
