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
	_ "github.com/mickep76/compress/snappy"
	_ "github.com/mickep76/compress/xz"
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

	if *dec {
		encoded, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		b, err := compress.Decode(*algo, encoded)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(b))
	} else {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		encoded, err := compress.Encode(*algo, b)
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(*out+"."+*algo, encoded, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
