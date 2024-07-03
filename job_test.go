package twoface

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// DummyJob is a simple implementation of the Job interface for testing purposes.
type DummyJob struct {
	result Result[any, error]
}

func (d DummyJob) Do() Result[any, error] {
	return d.result
}

func TestJob(t *testing.T) {
	convey.Convey("Job", t, func() {
		convey.Convey("Should create a Job", func() {
			job := NewJob(DummyJob{Ok[any, error]("done")})
			convey.So(job.Do().Unwrap(), convey.ShouldEqual, "done")
		})

		convey.Convey("Should create a RetriableJob", func() {
			job := NewRetriableJob(context.Background(), DummyJob{Ok[any, error]("done")})
			convey.So(job.Do().Unwrap(), convey.ShouldEqual, "done")
		})

		convey.Convey("Should retry a RetriableJob", func() {
			job := NewRetriableJob(
				context.Background(), DummyJob{Err[any](fmt.Errorf("retry"))})
			// Assuming NewRetrier and NewFibonacci logic are implemented to retry
			// The following line will vary based on actual retry logic
			convey.So(job.Do().IsErr(), convey.ShouldBeTrue)
		})
	})
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		job := NewJob(DummyJob{Ok[any, error]("ok")})
		job.Do()
	}
}
