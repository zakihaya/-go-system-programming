package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/google/shlex"
	"github.com/peterh/liner"
)

func parseCmd(cmdStr string) (cmd string, args []string, err error) {
	l := shlex.NewLexer(strings.NewReader(cmdStr))
	cmd, err = l.Next()
	if err != nil {
		return
	}
	for {
		token, nextErr := l.Next()
		if nextErr == io.EOF {
			return
		}
		args = append(args, token)
	}
}

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if cmdStr, err := line.Prompt(" "); err == nil {
			fmt.Println("cmd:", cmdStr)
			if cmdStr == "" {
				fmt.Println("continue")
				continue
			}
			// ここでコマンドを処理する
			fmt.Println("parse", cmdStr)
			cmd, args, err := parseCmd(cmdStr)
			fmt.Println("---")
			fmt.Println(cmd)
			fmt.Println(args)
			fmt.Println(err)
			fmt.Println("---")
		} else if errors.Is(err, io.EOF) {
			break
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Error reading line:", err)
		}
	}

	fmt.Println("== end")
}
