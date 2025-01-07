package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("ccsh> ")
		s.Scan()

		prcs := strings.Split(s.Text(), "|")
		cmds := make([]*exec.Cmd, 0, len(prcs))
		var inpPipe io.ReadCloser = os.Stdin
		for i, cmd := range prcs {
			cmdReq := strings.Split(strings.TrimSpace(cmd), " ")
			switch cmdReq[0] {
			case "exit":
				os.Exit(0)
			case "cd":
				_ = os.Chdir(cmdReq[1])
			}

			exc := exec.Command(cmdReq[0], cmdReq[1:]...)
			exc.Stdin = inpPipe
			cmds = append(cmds, exc)

			if i < len(prcs)-1 {
				stdout, _ := exc.StdoutPipe()
				inpPipe = stdout
			} else {
				exc.Stdout = os.Stdout
			}
		}

		// cmd1 - cmd2 - cmd3

		runCmds(cmds)

	}
}

func runCmds(cmds []*exec.Cmd) {
	var wg sync.WaitGroup

	for _, exc := range cmds {
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := exc.Run()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
