package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	out := make(chan job)

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
	out := make(chan job)

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
	out := make(chan job)

	go func() {
		for j := range in {
			j.value = int64(float64(j.value) / float64(rand.Intn(10)))
			j.state = FinishedStage

			fmt.Println(j)
			out <- j
		}
		close(out)
	}()

	return out
}

func main() {
	length := 50_000_000
	jobs := make([]job, length)
	in := make(chan job, len(jobs))
	for i := 0; i < length; i++ {
		jobs[i].value = int64(i)
		in <- jobs[i]
	}
	close(in)

	start := time.Now()

	result := LastProcessing(
		SecondProcessing(
			FirstProcessing(in),
		),
	)

	_ = result
	finished := time.Since(start)

	fmt.Println(finished)
}
