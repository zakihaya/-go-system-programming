package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))

	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM MST"))
	fmt.Println(now.Format("2006/01/02 15:04:05 MST"))
}
