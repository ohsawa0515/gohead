package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var n uint64
var c uint64

func main() {
	flag.Uint64Var(&n, "n", 10, "lines")
	flag.Uint64Var(&c, "c", 0, "bytes")
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
		// c オプションが指定された場合は N bytesまで表示する
		if err := HeadCharacter(f, c, output); err != nil {
			log.Fatal(err)
		}
		if c > 0 {

		} else {
			if err := HeadLine(f, n, output); err != nil {
				log.Fatal(err)
			}
		}
		if len(files) > 1 {
			fmt.Fprint(output, "\n")
		}
	}
}
