package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	n := 10
	file := "hoge.txt" // temp

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
