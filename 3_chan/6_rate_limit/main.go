package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 1 find bug
// 2 add limiter
// code: https://go.dev/play/p/6HWhEeaxi2S
func main() {
	count := 100
	ch := make(chan int, count)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {

		wg.Add(1)
		go func() {
			ch <- RPCCall(&wg)
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
	close(ch)
	for value := range ch {
		fmt.Println(value)
	}
}

func RPCCall(wg *sync.WaitGroup) int {
	defer wg.Done()

	return rand.Int()
}
