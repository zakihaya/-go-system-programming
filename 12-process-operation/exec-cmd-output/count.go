package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, time.Now())
		time.Sleep(time.Second)
	}
}
