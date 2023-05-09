package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := make(chan string, 5)

	tasks <- "s"

	go func() {
		tasks <- "a"
		time.Sleep(time.Second * 3)
		tasks <- "b"
		time.Sleep(time.Second * 3)
		tasks <- "c"
		time.Sleep(time.Second * 3)
		tasks <- "d"
		time.Sleep(time.Second * 3)
		tasks <- "e"
		time.Sleep(time.Second * 3)
		tasks <- "end"
	}()

	for {
		if tasks == nil {
			fmt.Println("channel is nil")
			break
		}
		if len(tasks) == 0 {
			fmt.Println("channel is empty")
			time.Sleep(time.Second)
			continue
		}
		task, ok := <-tasks
		if !ok {
			fmt.Println("* channel closed")
			break
		}
		fmt.Println(task)
		if task == "end" {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
