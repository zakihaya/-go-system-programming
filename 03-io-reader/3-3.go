package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	// zipファイル作成
	zipIoWriter, _ := os.Create("3-3.zip")
	defer zipIoWriter.Close()

	zipWriter := zip.NewWriter(zipIoWriter)
	defer zipWriter.Close()

	str := "this is content"
	reader := strings.NewReader(str)

	writer, _ := zipWriter.Create("newfile3-3.txt")
	io.Copy(writer, reader)
}
