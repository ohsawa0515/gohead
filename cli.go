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
	var lines, chars uint64 // lines, bytes

	flags := flag.NewFlagSet("head", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Uint64Var(&lines, "n", 10, "lines")
	flags.Uint64Var(&chars, "c", 0, "bytes")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	files := flags.Args()
	showFileName := false
	for _, file := range files {
		if len(files) > 1 {
			showFileName = true
		}
		if status := cli.headFile(file, lines, chars, showFileName); status != ExitCodeOK {
			return status
		}
	}

	return ExitCodeOK
}

func (cli *CLI) headFile(file string, lines, chars uint64, showFileName bool) int {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
		return ExitCodeParseFlagError
	}
	defer f.Close()

	head := &Head{
		file:   f,
		lines:  lines,
		chars:  chars,
		output: cli.outStream,
	}

	if showFileName {
		fmt.Fprintf(cli.outStream, "==> %s <==\n", file)
	}
	// chars オプションが指定された場合は N bytesまで表示する
	// lines オプションより優先する
	if chars > 0 {
		if err := head.ReadCharacter(); err != nil {
			log.Println(err)
			return ExitError
		}
	} else {
		if err := head.ReadLines(); err != nil {
			log.Println(err)
			return ExitError
		}
	}
	if showFileName {
		fmt.Fprint(cli.outStream, "\n")
	}

	return ExitCodeOK
}
