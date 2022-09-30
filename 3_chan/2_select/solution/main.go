package main

import (
	"fmt"
	"math/rand"
	"time"
)

// code: https://go.dev/play/p/QKQD8r7AnW_9
func main() {
	chanForResp := make(chan int)
	go RPCCall(chanForResp)

	select {
	case result := <-chanForResp:
		fmt.Println(result)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func RPCCall(ch chan<- int) {
	time.Sleep(time.Hour)

	ch <- rand.Int()
}
