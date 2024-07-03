package twoface

import (
	"errors"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestFuture(t *testing.T) {
	convey.Convey("Future and Promise", t, func() {
		convey.Convey("Should set and get result", func() {
			p, f := NewPromise[int]()
			p.Set(42, nil)
			result, err := f.Result()
			convey.So(result, convey.ShouldEqual, 42)
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("Should handle Then", func() {
			p, f := NewPromise[int]()
			ch := make(chan int)
			f.Then(func(result int) { ch <- result })
			p.Set(42, nil)
			convey.So(<-ch, convey.ShouldEqual, 42)
		})

		convey.Convey("Should handle Catch", func() {
			p, f := NewPromise[int]()
			ch := make(chan error)
			f.Catch(func(err error) { ch <- err })
			p.Set(0, errDummy)
			convey.So(<-ch, convey.ShouldEqual, errDummy)
		})

		convey.Convey("Should handle Finally", func() {
			p, f := NewPromise[int]()
			ch := make(chan bool)
			f.Finally(func() { ch <- true })
			p.Set(42, nil)
			convey.So(<-ch, convey.ShouldBeTrue)
		})
	})
}

func BenchmarkFuture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p, f := NewPromise[int]()
		go func() {
			time.Sleep(time.Millisecond)
			p.Set(42, nil)
		}()
		f.Result()
	}
}

var errDummy = errors.New("dummy error")
