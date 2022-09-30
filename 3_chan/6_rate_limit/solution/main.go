package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// code: https://go.dev/play/p/Tjy1TmZviXn
func main() {

	limiter := make(chan struct{}, 10)

	go func() {
		for {
			<-time.After(time.Second)
			for i := 0; i < 10; i++ {
				limiter <- struct{}{}
			}
		}
	}()

	count := 100
	ch := make(chan int, count)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ch <- RPCCallWithLimiter(limiter)
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()

		for i := 0; i < count; i++ {
			fmt.Println(<-ch)
		}
	}()

	wg.Wait()
}

func RPCCall() int {
	return rand.Int()
}

func RPCCallWithLimiter(limiter chan struct{}) int {
	<-limiter
	return RPCCall()
}
