package main

import "fmt"

// code: https://go.dev/play/p/asOfB4XAXyw
func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("no one will write to chan")
	}
}
