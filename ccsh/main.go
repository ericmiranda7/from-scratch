package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Printf("ccsh> ")
		exit := !s.Scan()
		if exit {
			os.Exit(0)
		}

		cmdReq := strings.Split(s.Text(), " ")
		switch cmdReq[0] {
		case "exit":
			os.Exit(0)
		case "cd":
			os.Chdir(cmdReq[1])
		default:
			exc := exec.Command(cmdReq[0], cmdReq[1:]...)
			exc.Stdout = os.Stdout
			exc.Run()
		}

	}
}
