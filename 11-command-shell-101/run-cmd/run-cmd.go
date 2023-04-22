package main

import (
	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
}
