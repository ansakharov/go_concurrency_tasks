package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type resp struct {
	id  int
	err error
}

// add ctx with timeout
func main() {
	rand.Seed(time.Now().Unix())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	chanForResp := make(chan resp)
	go RPCCall(ctx, chanForResp)

	select {
	case result := <-chanForResp:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("timeout ctx")
	}
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("request aborted due timeout"),
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- resp{
			id: rand.Int(),
		}
	}
}
