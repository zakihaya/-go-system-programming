package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	// 書き込みファイルをオープン
	file, err := os.Create("random.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	// rand.Readerを準備
	randReader := rand.Reader
	// LimitReaderを通す
	limitReader := io.LimitReader(randReader, 1024)
	// 書き込み実行
	io.Copy(file, limitReader)
}
