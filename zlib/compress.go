package zlib

import "github.com/mickep76/compress"

type algorithm struct{}

func init() {
	compress.Register("zlib", &algorithm{})
}
