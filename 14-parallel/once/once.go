package main

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("init")
}

func initialize() {
	fmt.Println("初期化処理")
}

var once sync.Once

func main() {
	// 3回呼び出しても1度しか呼ばれない
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
