package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	chLen = 100
)

// channels allows concurrent processing
type job struct {
	value int64
	state State
}

type State int

const (
	InitialState State = iota
	FirstStage
	SecondStage
	FinishedStage
)

func FirstProcessing(in <-chan job) chan job {
	out := make(chan job, chLen)

	go func() {
		for j := range in {
			j.value = int64(float64(j.value) * math.Pi)
			j.state = FirstStage

			out <- j
		}
		close(out)
	}()

	return out
}

func SecondProcessing(in <-chan job) chan job {
	out := make(chan job, chLen)

	go func() {
		for j := range in {
			j.value = int64(float64(j.value) * math.E)
			j.state = SecondStage

			out <- j
		}
		close(out)
	}()

	return out
}

func LastProcessing(in <-chan job) chan job {
	out := make(chan job, chLen)

	go func() {
		for j := range in {
			j.value = int64(float64(j.value) / float64(rand.Intn(10)))
			j.state = FinishedStage
			out <- j
		}
		close(out)
	}()

	return out
}

func main() {
	length := 50_000_000
	in := make(chan job, chLen)

	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		result := LastProcessing(
			SecondProcessing(
				FirstProcessing(in),
			),
		)

		for _ = range result {
		}
	}()

	for i := 0; i < length; i++ {
		in <- job{value: int64(i)}
	}
	close(in)

	wg.Wait()

	finished := time.Since(start)

	fmt.Println(finished)
}
