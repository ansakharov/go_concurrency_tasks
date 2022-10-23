package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	counter := 20

	for i := 0; i < counter; i++ {
		i := i
		// wg.Add(1) // add here

		go func() {
			// wg.Add(1) // or here
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}

	//time.Sleep(time.Second)
	wg.Wait()
}
