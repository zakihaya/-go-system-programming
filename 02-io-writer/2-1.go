package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(f, "%d %s", 1, "a")
}
