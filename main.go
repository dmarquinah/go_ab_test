package main

import (
	"flag"

	"github.com/dmarquinah/go_ab_test/generator"
)

func main() {
	sample_size := flag.Int("size", 10000, "Number of rows of data generated")
	flag.Parse()                     // Parse provided values from input flags
	generator.Generate(*sample_size) // Generates the csv file to start the evaluation
}
