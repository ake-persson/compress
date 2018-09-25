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
	//	_ "github.com/mickep76/compress/lzw"
	//	_ "github.com/mickep76/compress/snappy"
	//	_ "github.com/mickep76/compress/xz"
	//	_ "github.com/mickep76/compress/zlib"
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
