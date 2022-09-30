package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// code: https://go.dev/play/p/0PvwVzg3fYo
func main() {
	storage := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}

	uniqueIDs := make(chan int, capacity)
	for i := 0; i < capacity; i++ {
		i := i

		go func() {
			if _, ok := storage[doubles[i]]; !ok {
				mu.Lock()
				storage[doubles[i]] = struct{}{}
				mu.Unlock()

				uniqueIDs <- doubles[i]
			}
		}()
	}

	fmt.Printf("len of ids: %d\n", len(uniqueIDs))
	fmt.Println(uniqueIDs)
}
