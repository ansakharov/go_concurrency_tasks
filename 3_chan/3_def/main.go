package main

import "fmt"

// avoid deadlock
func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	}
}
