package twoface

import "context"

/*
Job is an interface any type can implement if they want to be able to use the worker pool.

Example:

type MyJob struct{}

	func (m MyJob) Do() Result[any, error] {
	    return Ok("done")
	}
*/
type Job interface {
	Do() Result[any, error]
}

/*
NewJob is a convenience method to convert any incoming structured type to a Job interface.

Example:

job := NewJob(MyJob{})
*/
func NewJob(jobType Job) Job {
	return jobType
}

/*
RetriableJob provides boilerplate for quickly building jobs that retry based on a backoff delay strategy.

Example:

retriableJob := NewRetriableJob(context.Background(), MyJob{})
*/
type RetriableJob struct {
	ctx context.Context
	fn  Job
}

/*
NewRetriableJob creates a new retriable job.

Example:

retriableJob := NewRetriableJob(context.Background(), MyJob{})
*/
func NewRetriableJob(ctx context.Context, fn Job) Job {
	return NewJob(RetriableJob{
		ctx: ctx,
		fn:  fn,
	})
}

/*
Do the job and retry x amount of times when needed.

Example:

retriableJob := NewRetriableJob(context.Background(), MyJob{})
result := retriableJob.Do()
*/
func (job RetriableJob) Do() Result[any, error] {
	return NewRetrier(NewFibonacci(3)).Do(job.fn)
}
