package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// code: https://go.dev/play/p/AIKRroy-rpV
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

func FirstProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) * math.Pi)
		job.state = FirstStage
		result = append(result, job)
	}

	return result
}

func SecondProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) * math.E)
		job.state = SecondStage
		result = append(result, job)
	}

	return result
}

func LastProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) / float64(rand.Intn(10)))
		job.state = FinishedStage
		result = append(result, job)
	}

	return result
}

func main() {
	length := 50_000_000
	jobs := make([]job, length)
	for i := 0; i < length; i++ {
		jobs[i].value = int64(i)
	}

	start := time.Now()
	jobs = LastProcessing(
		SecondProcessing(
			FirstProcessing(jobs),
		),
	)
	finished := time.Since(start)

	fmt.Println(finished)
}
