package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if cmd, err := line.Prompt(" "); err == nil {
			if cmd == "" {
				continue
			}
			// ここでコマンドを処理する
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
