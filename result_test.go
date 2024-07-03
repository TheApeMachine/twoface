package twoface

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestResult(t *testing.T) {
	convey.Convey("Result", t, func() {
		convey.Convey("Should create Ok result", func() {
			r := Ok[int, error](42)
			convey.So(r.IsOk(), convey.ShouldBeTrue)
			convey.So(r.IsErr(), convey.ShouldBeFalse)
			convey.So(r.Unwrap(), convey.ShouldEqual, 42)
		})

		convey.Convey("Should create Err result", func() {
			err := fmt.Errorf("an error")
			r := Err[int, error](err)
			convey.So(r.IsOk(), convey.ShouldBeFalse)
			convey.So(r.IsErr(), convey.ShouldBeTrue)
			convey.So(r.UnwrapErr(), convey.ShouldEqual, err)
		})

		convey.Convey("Should panic on Unwrap of Err result", func() {
			err := fmt.Errorf("an error")
			r := Err[int, error](err)
			convey.So(func() { r.Unwrap() }, convey.ShouldPanic)
		})

		convey.Convey("Should panic on UnwrapErr of Ok result", func() {
			r := Ok[int, error](42)
			convey.So(func() { r.UnwrapErr() }, convey.ShouldPanic)
		})

		convey.Convey("Should map Ok result", func() {
			r := Ok[int, error](42)
			newR := r.Map(func(v int) int { return v * 2 })
			convey.So(newR.Unwrap(), convey.ShouldEqual, 84)
		})

		convey.Convey("Should flat map Ok result", func() {
			r := Ok[int, error](42)
			newR := r.FlatMap(func(v int) Result[int, error] { return Ok[int, error](v * 2) })
			convey.So(newR.Unwrap(), convey.ShouldEqual, 84)
		})

		convey.Convey("Should and then Ok result", func() {
			r := Ok[int, error](42)
			newR := r.AndThen(func(v int) Result[int, error] { return Ok[int, error](v * 2) })
			convey.So(newR.Unwrap(), convey.ShouldEqual, 84)
		})
	})
}

func BenchmarkResult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := Ok[int, error](42)
		_ = r.Map(func(v int) int { return v * 2 })
	}
}
