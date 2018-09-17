package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mickep76/compression"
	_ "github.com/mickep76/compression/snappy"
	_ "github.com/mickep76/compression/xz"
)

// flags
// file
// output
// algo

func main() {
	text := "abc123\ndef456\nabc123\ndef456\nabc123\ndef456\n"
	algo := "snappy"
	file := "example." + algo

	if err := compression.ToFile(algo, file, []byte(text)); err != nil {
		log.Fatal(err)
	}

	b, err := compression.FromFile(algo, file, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
