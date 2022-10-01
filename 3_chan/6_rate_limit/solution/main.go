package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// add simple limiter, read from gorotuine
func main() {

	limiter := make(chan struct{}, 10)

	go func() {
		for {
			t := time.NewTicker(time.Second)
			<-t.C
			for i := 0; i < 10; i++ {
				limiter <- struct{}{}
			}
		}
	}()

	count := 50
	ch := make(chan int, count)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ch <- RPCCallWithLimiter(limiter)
		}()
	}

	wg.Add(1)
	go func() {
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
