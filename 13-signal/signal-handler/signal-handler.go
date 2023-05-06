package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())

	ctx := context.Background()
	sigctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	toctx, cancel2 := context.WithTimeout(ctx, time.Second*20)
	defer cancel2()

	select {
	case <-sigctx.Done():
		fmt.Println("signal")
	case <-toctx.Done():
		fmt.Println("timeout")
	}
}
