package main

import (
	"fmt"
	"time"
)

func finished() chan bool {
	b := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		b <- true
	}()
	return b
}

func countUp() chan int {
	cnt := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second)
			cnt <- i
		}
	}()
	return cnt
}

func main() {
	fmt.Println("begin main")

	reader := countUp()
	exit := finished()

	func() {
		for {
			select {
			case n := <-reader:
				fmt.Println(n)
			case <-exit:
				fmt.Println("case exit")
				return
			}
		}
	}()

	fmt.Println("finish main")
}
