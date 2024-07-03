package twoface

import (
	"fmt"
)

/*
Option is a generic type that can hold a value or indicate absence.

Example:

	opt := Some(42)
	fmt.Println(opt.IsSome()) // Output: true
*/
type Option[T any] struct {
	value *T
}

/*
Some creates an Option with a value.

Example:

	opt := Some(42)
	fmt.Println(opt.IsSome()) // Output: true
*/
func Some[T any](v T) Option[T] {
	return Option[T]{value: &v}
}

/*
None creates an Option without a value.

Example:

	opt := None[int]()
	fmt.Println(opt.IsNone()) // Output: true
*/
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

/*
IsSome returns true if the Option contains a value.

Example:

	opt := Some(42)
	fmt.Println(opt.IsSome()) // Output: true
*/
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

/*
IsNone returns true if the Option does not contain a value.

Example:

	opt := None[int]()
	fmt.Println(opt.IsNone()) // Output: true
*/
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

/*
Unwrap returns the contained value if present, or an error if not.

Example:

	opt := Some(42)
	val, err := opt.Unwrap()
	if err != nil {
	    fmt.Println(err)
	} else {
	    fmt.Println(val) // Output: 42
	}
*/
func (o Option[T]) Unwrap() (T, error) {
	if o.IsNone() {
		var zero T
		return zero, fmt.Errorf("called `Unwrap` on a `None` value")
	}
	return *o.value, nil
}

/*
UnwrapOr returns the contained value if present, or a default value if not.

Example:

	opt := None[int]()
	val := opt.UnwrapOr(99)
	fmt.Println(val) // Output: 99
*/
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return *o.value
}

/*
MatchHandlers holds the functions to handle the Some and None cases.
*/
type MatchHandlers[T any] struct {
	Some func(T)
	None func()
}

/*
Match applies the appropriate function based on whether the Option contains a value.

Example:

	opt := Some(42)
	opt.Match(MatchHandlers[int]{
	    Some: func(value int) {
	        fmt.Println("Got a value:", value) // Output: Got a value: 42
	    },
	    None: func() {
	        fmt.Println("No value present")
	    },
	})
*/
func (o Option[T]) Match(handlers MatchHandlers[T]) {
	if o.IsSome() {
		handlers.Some(*o.value)
	} else {
		handlers.None()
	}
}

/*
Map transforms the Option value using the provided function.

Example:

	opt := Some(42)
	mapped := opt.Map(func(value int) int {
	    return value + 1
	})
	fmt.Println(mapped.UnwrapOr(0)) // Output: 43
*/
func (o Option[T]) Map(f func(T) T) Option[T] {
	if o.IsSome() {
		return Some(f(*o.value))
	}
	return None[T]()
}

/*
FlatMap transforms the Option value using the provided function that returns an Option.

Example:

	opt := Some(42)
	flatMapped := opt.FlatMap(func(value int) Option[int] {
	    return Some(value + 1)
	})
	fmt.Println(flatMapped.UnwrapOr(0)) // Output: 43
*/
func (o Option[T]) FlatMap(f func(T) Option[T]) Option[T] {
	if o.IsSome() {
		return f(*o.value)
	}
	return None[T]()
}
