package main

import (
	"fmt"
	"os"
)

func main() {
	inputs := make(chan []byte)

	go func() {
		a, _ := os.ReadFile("a.txt")
		inputs <- a
	}()

	fmt.Println(string(<-inputs))

	go func() {
		b, _ := os.ReadFile("b.txt")
		inputs <- b
	}()

	fmt.Println(string(<-inputs))
}
