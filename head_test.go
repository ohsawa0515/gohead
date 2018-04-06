package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHeadLine(t *testing.T) {
	cases := []struct {
		file   string
		n      uint64
		actual string
	}{
		{file: "A\nBBB\nC\nD\nF\n", n: 2, actual: "A\nBBB\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", n: 10, actual: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		if err := HeadLine(f, c.n, buf); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		expected := buf.Bytes()
		actual := []byte(c.actual)
		if bytes.Compare(expected, actual) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}

func TestHeadCharacter(t *testing.T) {
	cases := []struct {
		file   string
		c      uint64
		actual string
	}{
		{file: "A\nBBB\nC\nD\nF\n", c: 10, actual: "A\nBBB\nC\nD\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", c: 200, actual: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
		{file: "あいうえお\n一二三四五\n", c: 6, actual: "あい"},
		{file: "あいうえお\n一二三四五\n", c: 22, actual: "あいうえお\n一"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		if err := HeadCharacter(f, c.c, buf); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		expected := buf.Bytes()
		actual := []byte(c.actual)
		fmt.Println(bytes.Compare(expected, actual))
		if bytes.Compare(expected, actual) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}
