package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func expandPath(path, workDir string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(workDir, path)
}

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if ipt, err := line.Prompt("$ "); err == nil {
			if ipt == "" {
				continue
			}
			// ここでコマンドを処理する
			cmdStr, args, err := parseCmd(ipt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			cmd := exec.Command(cmdStr, args...)
			o, err := cmd.CombinedOutput()
			fmt.Println(string(o))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
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
