package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Head display first line of a file
func Head(file io.Reader, n uint64, w io.Writer) error {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(file)
	for i := uint64(0); i < n; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		buf.Write(line)
	}

	fmt.Fprint(w, buf)
	return nil
}
