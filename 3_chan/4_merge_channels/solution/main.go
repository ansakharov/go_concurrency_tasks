package main

import "fmt"

// code: https://go.dev/play/p/mJQEs5Srpet
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func merge[T any](chns ...chan T) chan T {
	result := make(chan T)

	go func() {
		for _, ch := range chns {
			close(ch)
			for value := range ch {
				result <- value
			}
		}
		close(result)
	}()

	return result
}
