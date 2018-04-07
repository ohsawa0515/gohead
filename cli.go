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
	var quiet, verbose bool

	flags := flag.NewFlagSet("head", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Uint64Var(&lines, "n", 10, "print the first K lines of each file")
	flags.Uint64Var(&chars, "c", 0, "print the first K bytes of each file")
	flags.BoolVar(&quiet, "q", false, "never print headers giving file names")
	flags.BoolVar(&verbose, "v", false, "always print headers giving file names")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	files := flags.Args()
	for _, file := range files {
		if status := cli.headFile(file, lines, chars, isShowFileName(len(files), quiet, verbose)); status != ExitCodeOK {
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

func isShowFileName(num int, quiet, verbose bool) bool {
	if verbose {
		return true
	}
	if quiet {
		return false
	}
	if num > 1 {
		return true
	}
	return false
}
