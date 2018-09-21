package compress

import (
	"reflect"
	"testing"
)

type algorithm struct{}

func init() {
	Register("mock", &algorithm{})
}

func TestRegistered(t *testing.T) {
	if _, err := Registered("mock"); err != nil {
		t.Error(err)
	}

	if _, err := Registered("foo"); err == nil {
		t.Error("foo reports as registered when it's not")
	}
}

func TestAlgorithms(t *testing.T) {
	if !reflect.DeepEqual(Algorithms(), []string{"mock"}) {
		t.Error("registered algorithms got unexpected response")
	}
}
