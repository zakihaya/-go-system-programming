package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// encoding/csvパッケージを使って、csvをファイルに出力する
	file, err := os.Create("sample.csv")
	if err != nil {
		panic(err)
	}
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{"a", "b", "c"})
	csvWriter.Flush()
}
