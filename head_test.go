package main

import (
	"bytes"
	"testing"
)

func TestHeadLine(t *testing.T) {
	cases := []struct {
		file     string
		n        uint64
		expected string
	}{
		{file: "A\nBBB\nC\nD\nF\n", n: 2, expected: "A\nBBB\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", n: 10, expected: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		if err := HeadLine(f, c.n, buf); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		actual := buf.Bytes()
		expected := []byte(c.expected)
		if bytes.Compare(actual, expected) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}

func TestHeadCharacter(t *testing.T) {
	cases := []struct {
		file     string
		c        uint64
		expected string
	}{
		{file: "A\nBBB\nC\nD\nF\n", c: 10, expected: "A\nBBB\nC\nD\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", c: 200, expected: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
		{file: "あいうえお\n一二三四五\n", c: 6, expected: "あい"},
		{file: "あいうえお\n一二三四五\n", c: 22, expected: "あいうえお\n一二"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		if err := HeadCharacter(f, c.c, buf); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		actual := buf.Bytes()
		expected := []byte(c.expected)
		if bytes.Compare(actual, expected) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}
