package main

import (
	"bytes"
	"testing"
)

func TestHead(t *testing.T) {
	cases := []struct {
		file   string
		n      uint64
		actual string
	}{
		{file: "A\nB\nC\nD\nF\n", n: 2, actual: "A\nB\n"},
		{file: "A\nB\nC\nD\nF\n", n: 10, actual: "A\nB\nC\nD\nF\n"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		if err := Head(f, c.n, buf); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		actual := buf.Bytes()
		expected := []byte(c.actual)
		if bytes.Compare(expected, actual) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}
