package main

import (
	"fmt"
	"sync"
)

func sub(c int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("share by arguments:", c*c)
}

func main() {
	var wg sync.WaitGroup

	// 引数渡し
	wg.Add(1)
	go sub(10, &wg)

	// クロージャのキャプチャ渡し
	c := 20
	wg.Add(1)
	go func() {
		fmt.Println("share by capture:", c*c)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
}
