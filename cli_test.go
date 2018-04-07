package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCliRunParseFlagError(t *testing.T) {
	cases := []struct {
		args string
	}{
		{args: "head -foo=1 ./test/hoge.txt"},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeParseFlagError {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeParseFlagError)
		}
	}
}

func TestCliRunFileNotFound(t *testing.T) {
	cases := []struct {
		args string
	}{
		{args: "head ./test/foo.txt"},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeParseFlagError {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeParseFlagError)
		}
	}
}

func TestCliLines(t *testing.T) {
	cases := []struct {
		args     string
		expected string
	}{
		{args: "head -n=2 ./test/hoge.txt", expected: "A\nB\n"},
		{args: "head ./test/fuga.txt", expected: "a\nbb\nccc\ndddd\neeeee\nffffff\nggggggg\n"},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeOK {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeOK)
		}

		if outStream.String() != c.expected {
			t.Errorf("not matched; actual %v, expected %v", outStream.String(), c.expected)
		}
	}
}

func TestCliBytes(t *testing.T) {
	cases := []struct {
		args     string
		expected string
	}{
		{args: "head -c=6 ./test/hoge.txt", expected: "A\nB\nC\n"},
		{args: "head -c=100 ./test/fuga.txt", expected: "a\nbb\nccc\ndddd\neeeee\nffffff\nggggggg\n"},
		{args: "head -n=2 -c=8 ./test/fuga.txt", expected: "a\nbb\nccc"},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeOK {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeOK)
		}

		if outStream.String() != c.expected {
			t.Errorf("not matched; actual %v, expected %v", outStream.String(), c.expected)
		}
	}
}

func TestCliMultiFiles(t *testing.T) {
	cases := []struct {
		args     string
		expected string
	}{
		{args: "head -n=1 ./test/hoge.txt ./test/fuga.txt", expected: `==> ./test/hoge.txt <==
A

==> ./test/fuga.txt <==
a

`},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeOK {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeOK)
		}

		if outStream.String() != c.expected {
			t.Errorf("not matched; actual %v, expected %v", outStream.String(), c.expected)
		}
	}
}

func TestCliQuiet(t *testing.T) {
	cases := []struct {
		args     string
		expected string
	}{
		{args: "head -n=1 -q ./test/hoge.txt ./test/fuga.txt", expected: "A\na\n"},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, " ")

		status := cli.Run(args)
		if status != ExitCodeOK {
			t.Errorf("ExitStatus=%d, expected %d", status, ExitCodeOK)
		}

		if outStream.String() != c.expected {
			t.Errorf("not matched; actual %v, expected %v", outStream.String(), c.expected)
		}
	}
}
