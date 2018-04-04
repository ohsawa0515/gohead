package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var n uint64

func main() {
	flag.Uint64Var(&n, "n", 10, "lines")
	flag.Parse()
	files := flag.Args()
	output := os.Stdout

	for _, file := range files {
		f, err := os.OpenFile(file, os.O_RDONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if len(files) > 1 {
			fmt.Fprintf(output, "==> %s <==\n", file)
		}
		if err := Head(f, n, output); err != nil {
			log.Fatal(err)
		}
		if len(files) > 1 {
			fmt.Fprint(output, "\n")
		}
	}
}
