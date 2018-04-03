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
	file := flag.Args()[0]

	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, err := head(f, n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, w)
}
