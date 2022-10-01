package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// don't miss mutex for read
func main() {
	storage := make(map[int]struct{})
	mu := sync.RWMutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}

	wg := sync.WaitGroup{}
	uniqueIDs := make(chan int, capacity)
	for i := 0; i < capacity; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := storage[doubles[i]]; !ok {
				storage[doubles[i]] = struct{}{}
				// without defer mu.Unlock() code required
				// mu.Unlock()

				uniqueIDs <- doubles[i]
				// return
			}
			// mu.Unlock()
		}()
	}

	wg.Done()
	fmt.Printf("len of ids: %d\n", len(uniqueIDs))
	fmt.Println(uniqueIDs)
}
