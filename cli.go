// refer https://deeeet.com/writing/2014/12/18/golang-cli-test/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var n, c uint64 // lines, bytes

	flags := flag.NewFlagSet("head", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Uint64Var(&n, "n", 10, "lines")
	flags.Uint64Var(&c, "c", 0, "bytes")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	files := flags.Args()
	result := 0
	showFileName := false
	for _, file := range files {
		if len(files) > 1 {
			showFileName = true
		}
		if result != cli.headFile(file, n, c, showFileName) {
			return result
		}
	}

	return ExitCodeOK
}

func (cli *CLI) headFile(file string, n, c uint64, showFileName bool) int {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if showFileName {
		fmt.Fprintf(cli.outStream, "==> %s <==\n", file)
	}
	// c オプションが指定された場合は N bytesまで表示する
	// n オプションより優先する
	if c > 0 {
		if err := HeadCharacter(f, c, cli.outStream); err != nil {
			log.Println(err)
			return ExitError
		}
	} else {
		if err := HeadLine(f, n, cli.outStream); err != nil {
			log.Println(err)
			return ExitError
		}
	}
	if showFileName {
		fmt.Fprint(cli.outStream, "\n")
	}

	return ExitCodeOK
}
