package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// add simple limiter, read from gorotuine
func main() {
	const (
		limitPerSecond = 25
		requestCount   = 50
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int, requestCount)

	limiter := make(chan struct{})
	ticker := time.NewTicker(time.Second / limitPerSecond)
	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				limiter <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	wg := sync.WaitGroup{}
	wg.Add(requestCount)
	for i := 0; i < requestCount; i++ {
		go func() {
			defer wg.Done()

			<-limiter
			ch <- RPCCall()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}
}

func RPCCall() int {
	return rand.Int()
}

	