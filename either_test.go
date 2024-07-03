package twoface

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestEither(t *testing.T) {
	convey.Convey("Testing Either type", t, func() {
		convey.Convey("Left and IsLeft", func() {
			either := Left[int, string](1)
			convey.So(either.IsLeft(), convey.ShouldBeTrue)
		})

		convey.Convey("Right and IsRight", func() {
			either := Right[int]("hello")
			convey.So(either.IsRight(), convey.ShouldBeTrue)
		})

		convey.Convey("UnwrapLeft", func() {
			either := Left[int, string](42)
			val := either.UnwrapLeft()
			convey.So(val, convey.ShouldEqual, 42)
		})

		convey.Convey("UnwrapLeft with Right", func() {
			either := Right[int]("hello")
			convey.So(func() { either.UnwrapLeft() }, convey.ShouldPanicWith, "called `UnwrapLeft` on a `Right` value")
		})

		convey.Convey("UnwrapRight", func() {
			either := Right[int]("hello")
			val := either.UnwrapRight()
			convey.So(val, convey.ShouldEqual, "hello")
		})

		convey.Convey("UnwrapRight with Left", func() {
			either := Left[int, string](1)
			convey.So(func() { either.UnwrapRight() }, convey.ShouldPanicWith, "called `UnwrapRight` on a `Left` value")
		})
	})
}

func BenchmarkEither(b *testing.B) {
	b.Run("Benchmark Left and IsLeft", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			either := Left[int, string](1)
			if !either.IsLeft() {
				b.Fatal("IsLeft should be true")
			}
		}
	})

	b.Run("Benchmark Right and IsRight", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			either := Right[int]("hello")
			if !either.IsRight() {
				b.Fatal("IsRight should be true")
			}
		}
	})
}
