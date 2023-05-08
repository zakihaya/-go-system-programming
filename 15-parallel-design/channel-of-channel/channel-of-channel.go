package main

import (
	"fmt"
	"sync"
	"time"
)

// 終了した順に書き出し
// チャネルに結果が投入された順に処理される
func writeResultForEnd(channel chan int) {
	// 順番に取り出す
	for channel := range channel {
		if channel == 9 {
			break
		}
		fmt.Println(channel)
	}
}

// 開始した順に書き出し
// チャネルにチャネルを入れた（開始した順に処理される）
func writeResultForBegin(channels chan chan int) {
	// 順番に取り出す
	for ch := range channels {
		result := <-ch
		if result == 9 {
			break
		}
		fmt.Println(result)
	}
}

func tasks() chan int {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4)

	go func() {
		fmt.Println("task1 begin 3 sec")
		time.Sleep(3 * time.Second)
		fmt.Println("task1 end")
		ch <- 1
		wg.Done()
	}()

	go func() {
		fmt.Println("task2 begin 1 sec")
		time.Sleep(1 * time.Second)
		fmt.Println("task2 end")
		ch <- 2
		wg.Done()
	}()

	go func() {
		fmt.Println("task3 begin 4 sec")
		time.Sleep(4 * time.Second)
		fmt.Println("task3 end")
		ch <- 3
		wg.Done()
	}()

	go func() {
		fmt.Println("task4 begin 2 sec")
		time.Sleep(2 * time.Second)
		fmt.Println("task4 end")
		ch <- 4
		wg.Done()
	}()

	go func() {
		wg.Wait()
		ch <- 9
	}()

	return ch
}

func tasks2() chan chan int {
	var wg sync.WaitGroup
	channels := make(chan chan int)

	wg.Add(4)

	go func() {
		ch := make(chan int)
		channels <- ch
		fmt.Println("task1 begin 3 sec")
		time.Sleep(3 * time.Second)
		fmt.Println("task1 end")
		ch <- 1
		wg.Done()
	}()

	go func() {
		ch := make(chan int)
		channels <- ch
		fmt.Println("task2 begin 1 sec")
		time.Sleep(1 * time.Second)
		fmt.Println("task2 end")
		ch <- 2
		wg.Done()
	}()

	go func() {
		ch := make(chan int)
		channels <- ch
		fmt.Println("task3 begin 4 sec")
		time.Sleep(4 * time.Second)
		fmt.Println("task3 end")
		ch <- 3
		wg.Done()
	}()

	go func() {
		ch := make(chan int)
		channels <- ch
		fmt.Println("task4 begin 2 sec")
		time.Sleep(2 * time.Second)
		fmt.Println("task4 end")
		ch <- 4
		wg.Done()
	}()

	go func() {
		wg.Wait()
		ch := make(chan int)
		channels <- ch
		ch <- 9
	}()

	return channels
}

func main() {
	// 終了した順に書き出す
	ch := tasks()
	writeResultForEnd(ch)

	fmt.Println("-----")

	// 開始した順に書き出す
	ch2 := tasks2()
	writeResultForBegin(ch2)
}
