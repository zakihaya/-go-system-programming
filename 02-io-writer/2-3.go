package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "world",
	}

	// responseWriterのWrite実行時にgzip変換する
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	// gzip->responseの実行と同タイミングでstdoutにも内容を出力
	writer := io.MultiWriter(gzipWriter, os.Stdout)

	// writer出力時に内容をjson変換
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	// json変換実行と同時にwriterのWriteが走る
	encoder.Encode(source)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
