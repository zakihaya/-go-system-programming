package main

import (
	"fmt"
)

type StringFuture struct {
	receiver chan string
	cache    string
}

func NewStringFuture() (*StringFuture, chan string) {
	f := &StringFuture{
		receiver: make(chan string, 1),
	}
	return f, f.receiver
}

func (f *StringFuture) Get() string {
	r, ok := <-f.receiver
	if ok {
		fmt.Println("# value from receiver")
		close(f.receiver)
		f.cache = r
	} else {
		fmt.Println("# value from cache")
	}
	return f.cache
}

func (f *StringFuture) Close() {
	close(f.receiver)
}

func main() {
	sf, ch := NewStringFuture()

	ch <- "a"

	fmt.Println(sf.Get())
	fmt.Println(sf.Get())
}
