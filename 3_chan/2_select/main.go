package main

import (
	"fmt"
	"math/rand"
	"time"
)

// add timeout to avoid long waiting
func main() {
	rand.Seed(time.Now().Unix())

	chanForResp := make(chan int)
	go RPCCall(chanForResp)

	result := <-chanForResp
	fmt.Println(result)
}

func RPCCall(ch chan<- int) {
	// sleep 0-4 sec
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))

	ch <- rand.Int()
}
