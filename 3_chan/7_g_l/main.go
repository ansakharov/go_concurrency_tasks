package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// find problem in code and fix it
// code: https://go.dev/play/p/bcakSCXWQepo
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		work(ctx)
		// assume, that at i == 5 some error occurs
		if i == 5 {
			cancel()
		}
	}

	// server doesn't die. Imagine, it's doing useful work.
	for {
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
	return rand.Int()
}

func resCh() chan int {
	ch := make(chan int)

	return ch
}
