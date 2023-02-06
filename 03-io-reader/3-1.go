package main

import (
	"io"
	"os"
)

func main() {
	// 古いファイルオープン
	oldFile, readErr := os.Open("old.txt")
	defer oldFile.Close()
	if readErr != nil {
		panic(readErr)
	}

	// 新しいファイルオープン
	newFile, writeErr := os.Create("new.txt")
	defer newFile.Close()
	if writeErr != nil {
		panic(writeErr)
	}

	// 内容を新しいファイルに書き込み
	io.Copy(newFile, oldFile)
}
