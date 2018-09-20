package lzw

import "github.com/mickep76/compress"

const (
	LzwLSB = 0
	LzwMSB = 1
)

type algorithm struct{}

func init() {
	compress.Register("lzw", &algorithm{})
}
