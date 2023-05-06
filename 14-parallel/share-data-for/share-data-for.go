package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func() {
			// goroutineが起動するときにはループが回りきって
			// 全部のタスクが最後のタスクになってしまう
			fmt.Println(task)
			wg.Done()
		}()

		// OK
		// go func(taskVal string) {
		// 	fmt.Println(taskVal)
		// 	wg.Done()
		// }(task)
	}
	wg.Wait()
}
