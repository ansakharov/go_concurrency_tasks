package main

import "fmt"

// merge two channels
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

type s struct {
	
}
func (si s[T any]) merge(chans ...chan T) chan T {
	ch := make(chan T)
	close(ch)

	return ch
}
