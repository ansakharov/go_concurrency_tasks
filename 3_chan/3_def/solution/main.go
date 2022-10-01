package main

import "fmt"

// use default
func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("no one will write to chan")
	}
}
