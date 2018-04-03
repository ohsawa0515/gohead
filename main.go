package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

	reader := bufio.NewReader(f)
	for i := 0; i < n; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			break
		}
		fmt.Print(string(line))
	}
}
