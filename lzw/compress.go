package lzw

import "github.com/mickep76/compress"

const (
	// LzwLSB Least Significant Bit.
	LzwLSB = 0

	// LzwMSB Most Significant Bit.
	LzwMSB = 1
)

type algorithm struct{}

func init() {
	compress.Register("lzw", &algorithm{})
}
