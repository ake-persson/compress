package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/mickep76/compress"
	_ "github.com/mickep76/compress/gzip"
	_ "github.com/mickep76/compress/snappy"
	_ "github.com/mickep76/compress/xz"
)

func main() {
	file := flag.String("f", "", "File.")
	out := flag.String("o", "example", "Output.")
	algo := flag.String("a", "gzip", fmt.Sprintf("Algorithms: [%s].", strings.Join(compress.Algorithms(), ", ")))
	exp := flag.Bool("x", false, "Expand.")

	flag.Parse()

	if *exp {
		b, err := compress.FromFile(*algo, *file, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(b))
	} else {
		b, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}

		if err := compress.ToFile(*algo, fmt.Sprintf("%s.%s", *out, *algo), b); err != nil {
			log.Fatal(err)
		}
	}
}
