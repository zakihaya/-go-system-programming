package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() started")
		time.Sleep(time.Second)
		fmt.Println("sub() finished")
		// 終了を通知
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
