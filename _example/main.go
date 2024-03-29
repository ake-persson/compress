package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ake-persson/compress"
	_ "github.com/ake-persson/compress/gzip"
	_ "github.com/ake-persson/compress/lzw"
	_ "github.com/ake-persson/compress/snappy"
	_ "github.com/ake-persson/compress/xz"
	_ "github.com/ake-persson/compress/zlib"
)

func usage() {
	fmt.Printf("Usage: example [options] file\n\nOptions:\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	out := flag.String("out", "example", "Output.")
	algo := flag.String("algo", "gzip", fmt.Sprintf("Algorithms: [%s].", strings.Join(compress.Algorithms(), ", ")))
	dec := flag.Bool("dec", false, "Decode.")

	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
	}
	file := flag.Args()[0]

	a, err := compress.NewAlgorithm(*algo)
	if err != nil {
		log.Fatal(err)
	}

	if *dec {
		encoded, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		b, err := a.Decode(encoded)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(b))
	} else {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		encoded, err := a.Encode(b)
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(*out+"."+*algo, encoded, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
