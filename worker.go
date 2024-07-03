package twoface

import (
	"context"
	"fmt"
	"time"
)

// Worker processes jobs from the job channel.
type Worker struct {
	ID           int
	WorkerPool   chan chan Job
	JobChannel   chan Job
	ctx          context.Context
	lastUse      time.Time
	lastDuration int64
	drain        bool
}

// NewWorker creates a new worker.
func NewWorker(ID int, workerPool chan chan Job, ctx context.Context) *Worker {
	return &Worker{
		ID:         ID,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		ctx:        ctx,
		lastUse:    time.Now(),
		drain:      false,
	}
}

// Start the worker to be ready to accept jobs from the job queue.
func (worker *Worker) Start() *Worker {
	go func() {
		for {
			worker.WorkerPool <- worker.JobChannel

			select {
			case job := <-worker.JobChannel:
				worker.lastUse = time.Now()
				result := job.Do()
				if result.IsErr() {
					fmt.Printf("Worker %d: Job failed with error: %v\n", worker.ID, result.UnwrapErr())
				}
				worker.lastDuration = time.Since(worker.lastUse).Nanoseconds()

				if worker.drain {
					return
				}
			case <-worker.ctx.Done():
				return
			}
		}
	}()
	return worker
}

// Drain the worker, which means it will finish its current job first before it will stop.
func (worker *Worker) Drain() {
	worker.drain = true
}
