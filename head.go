package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Head struct {
	file   io.Reader
	output io.Writer
	lines  uint64
	chars  uint64
}

// ReadLines display first line of a file
func (h *Head) ReadLines() error {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(h.file)
	for i := uint64(0); i < h.lines; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		buf.Write(line)
	}

	fmt.Fprint(h.output, buf)
	return nil
}

// ReadCharacter display first bytes of a file
func (h *Head) ReadCharacter() error {
	buf := new(bytes.Buffer)
	reader := bufio.NewReader(h.file)
	for i := uint64(0); i < h.chars; i++ {
		b, err := reader.ReadByte()
		if err != nil && err != io.EOF {
			return err
		} else if err == io.EOF {
			break
		}
		buf.WriteByte(b)
	}

	fmt.Fprint(h.output, buf)
	return nil
}
