package twoface

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestOption(t *testing.T) {
	convey.Convey("Testing Option type", t, func() {
		convey.Convey("Some and IsSome", func() {
			opt := Some(42)
			convey.So(opt.IsSome(), convey.ShouldBeTrue)
		})

		convey.Convey("None and IsNone", func() {
			opt := None[int]()
			convey.So(opt.IsNone(), convey.ShouldBeTrue)
		})

		convey.Convey("Unwrap", func() {
			opt := Some(42)
			val, err := opt.Unwrap()
			convey.So(err, convey.ShouldBeNil)
			convey.So(val, convey.ShouldEqual, 42)
		})

		convey.Convey("Unwrap with None", func() {
			opt := None[int]()
			_, err := opt.Unwrap()
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("UnwrapOr", func() {
			opt := None[int]()
			val := opt.UnwrapOr(99)
			convey.So(val, convey.ShouldEqual, 99)
		})

		convey.Convey("Match", func() {
			opt := Some(42)
			called := false
			opt.Match(MatchHandlers[int]{
				Some: func(value int) {
					called = true
					convey.So(value, convey.ShouldEqual, 42)
				},
				None: func() {
					t.Error("None should not be called")
				},
			})
			convey.So(called, convey.ShouldBeTrue)
		})

		convey.Convey("Map", func() {
			opt := Some(42)
			mapped := opt.Map(func(value int) int {
				return value + 1
			})
			convey.So(mapped.UnwrapOr(0), convey.ShouldEqual, 43)
		})

		convey.Convey("FlatMap", func() {
			opt := Some(42)
			flatMapped := opt.FlatMap(func(value int) Option[int] {
				return Some(value + 1)
			})
			convey.So(flatMapped.UnwrapOr(0), convey.ShouldEqual, 43)
		})
	})
}

func BenchmarkOption(b *testing.B) {
	b.Run("Benchmark Some and IsSome", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			opt := Some(42)
			if !opt.IsSome() {
				b.Fatal("IsSome should be true")
			}
		}
	})

	b.Run("Benchmark None and IsNone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			opt := None[int]()
			if !opt.IsNone() {
				b.Fatal("IsNone should be true")
			}
		}
	})

	b.Run("Benchmark Map", func(b *testing.B) {
		opt := Some(42)
		for i := 0; i < b.N; i++ {
			mapped := opt.Map(func(value int) int {
				return value + 1
			})
			if mapped.UnwrapOr(0) != 43 {
				b.Fatal("Map result should be 43")
			}
		}
	})

	b.Run("Benchmark FlatMap", func(b *testing.B) {
		opt := Some(42)
		for i := 0; i < b.N; i++ {
			flatMapped := opt.FlatMap(func(value int) Option[int] {
				return Some(value + 1)
			})
			if flatMapped.UnwrapOr(0) != 43 {
				b.Fatal("FlatMap result should be 43")
			}
		}
	})
}
