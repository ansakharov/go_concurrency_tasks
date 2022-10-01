package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// add limiter with limit 10 per second
// read values concurrently
func main() {
	count := 50
	ch := make(chan int, count)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			ch <- RPCCall()
		}()
	}

	wg.Wait()
	close(ch)
	for value := range ch {
		fmt.Println(value)
	}
}

func RPCCall() int {
	return rand.Int()
}
