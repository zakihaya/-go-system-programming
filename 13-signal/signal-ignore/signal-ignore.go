package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 最初の5秒はCtrl + Cで止まる
	fmt.Println("Accept Ctrl + C for 5 seconds")
	time.Sleep(time.Second * 5)

	// 可変長引数で任意の数のシグナルを設定可能
	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	// 次の5秒はCtrl + Cを無視する
	fmt.Println("Ignore Ctrl + C for 5 seconds")
	time.Sleep(time.Second * 5)
}
