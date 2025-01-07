package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	isLineNo := flag.Bool("n", false, "Show line numbers")
	ignNewline := flag.Bool("b", false, "Ignore newline numbering")
	flag.Parse()

	if os.Args[1] == "-" {
		f = os.Stdin
		catFile(f, *isLineNo, 0, *ignNewline)
		return
	}

	lineNo := 0
	fnames := os.Args[1:]
	for _, fname := range fnames {
		f, _ = os.Open(fname)
		lineNo = catFile(f, *isLineNo, lineNo, *ignNewline)
	}
}

func catFile(f *os.File, isLine bool, lineNo int, ignoreNewline bool) int {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		if sc.Text() == "" && ignoreNewline {
			fmt.Println()
			continue
		}

		lineNo++
		if isLine {
			fmt.Printf("%v ", lineNo)
		}
		fmt.Printf("%v\n", sc.Text())
	}

	return lineNo
}
