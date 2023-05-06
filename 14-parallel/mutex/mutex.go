package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock()をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	// sync.Mutex構造体の変数宣言
	var mu sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("id: %d\n", generateId(&mu))
			wg.Done()
		}()
	}
	wg.Wait()
}
