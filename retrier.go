package twoface

import (
	"fmt"
	"math"
	"time"
)

// Retrier interface to implement different retry strategies.
type Retrier interface {
	Do(Job) Result[any, error]
}

// NewRetrier creates a new retrier.
func NewRetrier(retrierType Retrier) Retrier {
	return retrierType
}

// Fibonacci is a RetryStrategy that retries a function n times with a Fibonacci interval in seconds between retries.
type Fibonacci struct {
	max int
	n   int
}

// NewFibonacci creates a new Fibonacci retrier.
func NewFibonacci(max int) Retrier {
	return NewRetrier(Fibonacci{
		max: max,
		n:   0,
	})
}

// Do retries the job with a Fibonacci backoff strategy.
func (strategy Fibonacci) Do(fn Job) Result[any, error] {
	if strategy.n > strategy.max {
		return Err[any, error](fmt.Errorf("maximum retries reached"))
	}

	result := fn.Do()
	if result.IsOk() {
		return result
	}

	strategy.n = int(math.Round((math.Pow(math.Phi, float64(strategy.n)) + math.Pow(math.Phi-1, float64(strategy.n))) / math.Sqrt(5)))
	time.Sleep(time.Duration(strategy.n) * time.Second)
	return strategy.Do(fn)
}
