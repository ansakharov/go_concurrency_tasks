package main

import (
	"fmt"
	"sync"
)

// create map with make and use mutex
func main() {
	storage := make(map[int]int, 1000)
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	writes := 1000

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
