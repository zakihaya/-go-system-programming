package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 2 seconds")
	after := time.After(2 * time.Second)
	fmt.Println("after")
	<-after
	fmt.Println("done")

	// time.Tick
	fmt.Println("notify each 2 seconds")
	var cnt int
	for now := range time.Tick(2 * time.Second) {
		cnt++
		fmt.Println(now)
		fmt.Println(cnt)
		if cnt > 3 {
			break
		}
	}
	fmt.Println("done")

	// time.NewTicker
	// 途中でstopが可能
	t := time.NewTicker(2 * time.Second)
	fmt.Println("notify each 2 seconds: time.NewTicker")
	var cnt2 int
	for {
		select {
		case gt := <-t.C:
			cnt2++
			fmt.Println(gt)
			fmt.Println(cnt2)
		}
		if cnt2 > 3 {
			break
		}
	}
	fmt.Println("done")
}
