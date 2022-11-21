package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// add ctx with timeout
func main() {
	chanForResp := RPCCall()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	select {
	case result := <-chanForResp:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("timeout ctx")
	}

	go func() {
		fmt.Println(<-chanForResp)
	}()

	time.Sleep(time.Second * 10)
}

func RPCCall() <-chan int {
	chanForResp := make(chan int)

	go func() {
		defer close(chanForResp)

		time.Sleep(time.Second * 5)

		chanForResp <- rand.Int()
	}()

	return chanForResp
}
