package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Poolを作成。Newで新規作成時のコードを実装
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	// 追加した要素から受け取れるプールが空だと新規作成
	pool.Put("manually added: 1")
	pool.Put("manually added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // これは新規作成
	pool.Put("manually added: 3")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	// => go run sync_pool.go
	// manually added: 1
	// manually added: 2
	// created: 1
	// manually added: 3
	// created: 2
	pool.Put("removed 1")
	pool.Put("removed 2")
	runtime.GC()
	runtime.GC()
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
