package compress

import (
	"io"

	"github.com/pkg/errors"
)

var algorithms = make(map[string]Algorithm)

// Algorithm interface.
type Algorithm interface {
	NewEncoder(writer io.Writer) (Encoder, error)
	NewDecoder(reader io.Reader) (Decoder, error)
}

// Register algorithm.
func Register(name string, algorithm Algorithm) {
	algorithms[name] = algorithm
}

// Registered algorithm.
func Registered(name string) (Algorithm, error) {
	c, ok := algorithms[name]
	if !ok {
		return nil, errors.Errorf("algorithm not registered: %s", name)
	}
	return c, nil
}

// Algorithms registered.
func Algorithms() []string {
	l := []string{}
	for k := range algorithms {
		l = append(l, k)
	}
	return l
}
