package main

import (
	"fmt"
	"math/rand"
	"time"
)

// add timeout to avoid long waiting
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
