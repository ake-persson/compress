package xz

import "github.com/ulikunitz/xz"

type decoder struct {
	decoder *xz.Reader
}

func (d *decoder) Decode(v []byte) (int, error) {
	return d.decoder.Read(v)
}
