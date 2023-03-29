package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main() begin")
	timer := time.After(3 * time.Second)
	fmt.Println(<-timer)
	// time.Sleep(time.Second)
	fmt.Println("main() end")
}
