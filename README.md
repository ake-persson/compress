[![GoDoc](https://godoc.org/github.com/mickep76/compress?status.svg)](https://godoc.org/github.com/mickep76/compress)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/compress)](https://goreportcard.com/report/github.com/mickep76/compress)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/compress/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/mickep76/compress.svg?branch=master)](https://travis-ci.org/mickep76/compress)
[![codecov](https://codecov.io/gh/mickep76/compress/branch/master/graph/badge.svg)](https://codecov.io/gh/mickep76/compress)

# compress

Package provides a generic interface to compress and un-compress

## Example

```go
package main

import (
        "flag"
        "fmt"
        "io/ioutil"
        "log"
        "os"
        "strings"

        "github.com/mickep76/compress"
        _ "github.com/mickep76/compress/gzip"
        _ "github.com/mickep76/compress/lzw"
        _ "github.com/mickep76/compress/snappy"
        _ "github.com/mickep76/compress/xz"
        _ "github.com/mickep76/compress/zlib"
)

func usage() {
        fmt.Printf("Usage: example [options] file\n\nOptions:\n")
        flag.PrintDefaults()
        os.Exit(0)
}

func main() {
        out := flag.String("out", "example", "Output.")
        method := flag.String("method", "gzip", fmt.Sprintf("Methods: [%s].", strings.Join(compress.Methods(), ", ")))
        dec := flag.Bool("dec", false, "Decode.")

        flag.Parse()

        if len(flag.Args()) < 1 {
                usage()
        }
        file := flag.Args()[0]

        m, err := compress.GetMethod(*method)
        if err != nil {
                log.Fatal(err)
        }

        if *dec {
                encoded, err := ioutil.ReadFile(file)
                if err != nil {
                        log.Fatal(err)
                }

                b, err := m.Decode(encoded)
                if err != nil {
                        log.Fatal(err)
                }

                fmt.Print(string(b))
        } else {
                b, err := ioutil.ReadFile(file)
                if err != nil {
                        log.Fatal(err)
                }

                encoded, err := m.Encode(b)
                if err != nil {
                        log.Fatal(err)
                }

                if err := ioutil.WriteFile(*out+"."+*method, encoded, 0644); err != nil {
                        log.Fatal(err)
                }
        }
}
```
