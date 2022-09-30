package main

import (
	"fmt"
	"sync"
)

// code: https://go.dev/play/p/pek5x9LdQCx
func main() {
	var storage map[int]int

	wg := sync.WaitGroup{}
	writes := 1000

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
