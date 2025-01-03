package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestPrintByteLen(t *testing.T) {
	f, _ := os.Open("./test.txt")

	var expect int64 = 342190
	got := getByteLen(f)

	if got != expect {
		t.Fatalf("got %v, want %v", got, expect)
	}
}

func TestGetLineCount(t *testing.T) {
	f, _ := os.Open("./test.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	expect := 7145
	got := getLineCount(s)

	if got != expect {
		t.Fatalf("got %v, want %v", got, expect)
	}
}

func TestGetWordCount(t *testing.T) {
	f, _ := os.Open("./test.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	expect := 58164
	got := getWordCount(s)

	if got != expect {
		t.Fatalf("got %v, want %v", got, expect)
	}
}

func TestIsStdin(t *testing.T) {
	args := [][]string{
		{"wc", "-l"},
		{"wc", "-l", "filename"},
		{"wc"},
		{"wc", "filez"},
	}
	expTable := []bool{true, false, true, false}

	for i, arg := range args {
		got := isStdin(arg)
		if got != expTable[i] {
			log.Fatalf("case %v: %v %v", i+1, got, expTable[i])
		}
	}
}

func TestCharCount(t *testing.T) {
	fname := "./test.txt"
	f, _ := os.Open(fname)
	b, _ := io.ReadAll(f)
	s := string(b)

	cmd := "wc -m ./test.txt | cut -f1 -d \" \""
	c := exec.Command("bash", "-c", cmd)
	o, _ := c.Output()
	expect, err := strconv.Atoi(strings.TrimSpace(string(o)))
	if err != nil {
		fmt.Println(err)
	}

	got := getCharCount(s)

	if got != expect {
		t.Fatalf("got %v, want %v", got, expect)
	}
}
