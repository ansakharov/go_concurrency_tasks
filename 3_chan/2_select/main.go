package main

import (
	"fmt"
	"math/rand"
	"time"
)

// code: https://go.dev/play/p/LoMOFIuHbnF
func main() {
	chanForResp := make(chan int)
	go RPCCall(chanForResp)

	result := <-chanForResp
	fmt.Println(result)
}

func RPCCall(ch chan<- int) {
	time.Sleep(time.Hour)

	ch <- rand.Int()
}
