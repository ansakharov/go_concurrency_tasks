package main

import (
	"fmt"
	"time"
)

// print square of range 0...20 in random order
// code: https://go.dev/play/p/Mwmn0RJEBxN
func main() {
	counter := 20
	for i := 0; i < counter; i++ {
		go func() {
			fmt.Println(i * i)
		}()
	}

	time.Sleep(time.Second)
}
