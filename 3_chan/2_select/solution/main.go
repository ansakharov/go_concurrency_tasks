package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// add ctx with timeout
func main() {
	chanForResp := make(chan int)
	go RPCCall(chanForResp)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	select {
	case result := <-chanForResp:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("timeout ctx")
	}
}

func RPCCall(ch chan<- int) {
	time.Sleep(time.Hour)

	ch <- rand.Int()
}
