package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// HeadLine display first line of a file
func HeadLine(file io.Reader, n uint64, w io.Writer) error {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(file)
	for i := uint64(0); i < n; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		buf.Write(line)
	}

	fmt.Fprint(w, buf)
	return nil
}

// HeadCharacter display first bytes of a file
func HeadCharacter(file io.Reader, c uint64, w io.Writer) error {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(file)
	for i := uint64(0); i < c; i++ {
		b, err := reader.ReadByte()
		if err != nil && err != io.EOF {
			return err
		} else if err == io.EOF {
			break
		}
		buf.WriteByte(b)
	}

	fmt.Fprint(w, buf)
	return nil
}
