package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mickep76/compression"
	_ "github.com/mickep76/compression/snappy"
)

func main() {
	text := "abc123\ndef456\nabc123\ndef456\nabc123\ndef456\n"

	if err := compression.ToFile("snappy", "example.snappy", []byte(text)); err != nil {
		log.Fatal(err)
	}

	b, err := compression.FromFile("snappy", "example.snappy", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))

	if err := os.Remove("example.snappy"); err != nil {
		log.Fatal(err)
	}
}
