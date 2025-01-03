package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func main() {
	readStdin := isStdin(os.Args)
	var f *os.File
	if readStdin {
		f = os.Stdin
	} else {
		f, _ = os.Open(os.Args[len(os.Args)-1])
	}
	opts := extractOpts(os.Args)
	choose(f, opts)
}

func isStdin(args []string) bool {
	return len(args) == 1 || strings.HasPrefix(args[len(args)-1], "-")
}

func extractOpts(args []string) []string {
	if len(args) <= 2 {
		return []string{"-l", "-w", "-c"}
	}

	return args[1 : len(args)-1]
}

func choose(f *os.File, opts []string) {
	output := make([]string, 3)
	bytes, _ := io.ReadAll(f)
	strContent := string(bytes)

	for _, opt := range opts {
		switch opt {
		case "-c":
			output = append(output, strconv.Itoa(len(bytes)))
		case "-l":
			output = append(output, strconv.Itoa(getLineCount(strContent)))
		case "-w":
			output = append(output, strconv.Itoa(getWordCount(strContent)))
		case "-m":
			output = append(output, strconv.Itoa(getCharCount(strContent)))
		}
	}

	for _, o := range output {
		fmt.Printf("%v ", o)
	}
	fmt.Printf("%v\n", f.Name())
}

func getCharCount(s string) int {
	charLen := 0
	for range s {
		charLen++
	}
	return charLen
}

func getWordCount(s string) int {
	return len(strings.Fields(s))
}

func getLineCount(s string) int {
	return strings.Count(s, "\n")
}

func getByteLen(file fs.File) int64 {
	ft, _ := file.Stat()
	return ft.Size()
}
