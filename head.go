package main

import (
	"bufio"
	"bytes"
	"io"
)

func head(file io.Reader, n int, w io.Writer) (io.Writer, error) {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(file)
	for i := 0; i < n; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		buf.Write(line)
	}

	return buf, nil
}
