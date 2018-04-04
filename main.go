package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var n int

func main() {
	flag.IntVar(&n, "n", 10, "count")
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
