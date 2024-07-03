package twoface

import "fmt"

/*
Result represents a value that can be either Ok or Err.

Example:

okResult := Ok
errResult := Err[int, error](fmt.Errorf("an error"))
*/
type Result[T any, E error] struct {
	ok  *T
	err *E
}

/*
Ok creates a Result with an Ok value.

Example:

okResult := Ok
*/
func Ok[T any, E error](value T) Result[T, E] {
	return Result[T, E]{ok: &value}
}

/*
Err creates a Result with an Err value.

Example:

errResult := Err[int, error](fmt.Errorf("an error"))
*/
func Err[T any, E error](err E) Result[T, E] {
	return Result[T, E]{err: &err}
}

/*
IsOk returns true if the Result is Ok.

Example:

okResult := Ok
fmt.Println(okResult.IsOk()) // true
*/
func (r Result[T, E]) IsOk() bool {
	return r.ok != nil
}

/*
IsErr returns true if the Result is Err.

Example:

errResult := Err[int, error](fmt.Errorf("an error"))
fmt.Println(errResult.IsErr()) // true
*/
func (r Result[T, E]) IsErr() bool {
	return r.err != nil
}

/*
Unwrap returns the Ok value or panics if the Result is Err.

Example:

okResult := Ok
fmt.Println(okResult.Unwrap()) // 42
*/
func (r Result[T, E]) Unwrap() T {
	if r.IsErr() {
		panic(fmt.Sprintf("called `Unwrap` on an `Err` value: %v", *r.err))
	}
	return *r.ok
}

/*
UnwrapErr returns the Err value or panics if the Result is Ok.

Example:

errResult := Err[int, error](fmt.Errorf("an error"))
fmt.Println(errResult.UnwrapErr()) // "an error"
*/
func (r Result[T, E]) UnwrapErr() E {
	if r.IsOk() {
		panic("called `UnwrapErr` on an `Ok` value")
	}
	return *r.err
}

/*
Map transforms the Ok value using the provided function.

Example:

okResult := Ok
newResult := okResult.Map(func(v int) int { return v * 2 })
fmt.Println(newResult.Unwrap()) // 84
*/
func (r Result[T, E]) Map(f func(T) T) Result[T, E] {
	if r.IsOk() {
		return Ok[T, E](f(*r.ok))
	}
	return r
}

/*
FlatMap transforms the Ok value using the provided function that returns a Result.

Example:

okResult := Ok
newResult := okResult.FlatMap(func(v int) Result[int, error] { return Ok[int, error](v * 2) })
fmt.Println(newResult.Unwrap()) // 84
*/
func (r Result[T, E]) FlatMap(f func(T) Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return f(*r.ok)
	}
	return r
}

/*
AndThen applies a function to the Ok value, returning a new Result.

Example:

okResult := Ok
newResult := okResult.AndThen(func(v int) Result[int, error] { return Ok[int, error](v * 2) })
fmt.Println(newResult.Unwrap()) // 84
*/
func (r Result[T, E]) AndThen(f func(T) Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return f(*r.ok)
	}
	return r
}
