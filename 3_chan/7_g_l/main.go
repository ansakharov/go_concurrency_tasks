package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// why goroutines leaked? find problem and fix it
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		go work(ctx)
	}
	fmt.Printf("Goroutines running: %d\n", runtime.NumGoroutine())

	cancel()

	// server doesn't die. Imagine, it's doing useful work.
	for {
		fmt.Printf("Goroutines leaks: %d\n", runtime.NumGoroutine()-1)
		fmt.Printf("i do some useful work, print num: %d\n", rand.Int())
		time.Sleep(time.Second)
	}
}

func work(ctx context.Context) {
	ch := resCh()

	go func() {
		ch <- rpcCall()
	}()

	select {
	case value := <-ch:
		fmt.Printf("result of rpcCall: %d\n", value)
	case <-ctx.Done():
		return
	}
}

func rpcCall() int {
	<-time.After(time.Minute)
	return rand.Int()
}

func resCh() chan int {
	ch := make(chan int)

	return ch
}
