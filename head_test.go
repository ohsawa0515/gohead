package main

import (
	"bytes"
	"testing"
)

func TestReadLines(t *testing.T) {
	cases := []struct {
		file     string
		lines    uint64
		expected string
	}{
		{file: "A\nBBB\nC\nD\nF\n", lines: 2, expected: "A\nBBB\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", lines: 10, expected: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		head := &Head{
			file:   f,
			lines:  c.lines,
			output: buf,
		}
		if err := head.ReadLines(); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		actual := buf.Bytes()
		expected := []byte(c.expected)
		if bytes.Compare(actual, expected) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}

func TestReadCharacter(t *testing.T) {
	cases := []struct {
		file     string
		chars    uint64
		expected string
	}{
		{file: "A\nBBB\nC\nD\nF\n", chars: 10, expected: "A\nBBB\nC\nD\n"},
		{file: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n", chars: 200, expected: "A\nBBB\nCCCC\nDDDDD\nFFFFFF\n"},
		{file: "あいうえお\n一二三四五\n", chars: 6, expected: "あい"},
		{file: "あいうえお\n一二三四五\n", chars: 22, expected: "あいうえお\n一二"},
	}

	for _, c := range cases {
		f := bytes.NewBufferString(c.file)
		buf := new(bytes.Buffer)
		head := &Head{
			file:   f,
			chars:  c.chars,
			output: buf,
		}
		if err := head.ReadCharacter(); err != nil {
			t.Errorf("unextected error: %v", err)
		}

		actual := buf.Bytes()
		expected := []byte(c.expected)
		if bytes.Compare(actual, expected) != 0 {
			t.Errorf("not matched; actual %v, expected %v", string(actual), string(expected))
		}
	}
}
