package main

import (
	"fmt"
)

func Function(param string) {
	fmt.Println("Function called", param)
}

func main() {
	fmt.Println("begin")

	Function("1")

	func() {
		fmt.Println("exec noname func 1")
	}()

	go Function("2")

	go func() {
		fmt.Println("exec nonemae func 2")
	}()

	fmt.Println("end")
}
