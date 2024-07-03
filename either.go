package twoface

/*
Either represents a value that can be one of two possible types.

Example:

	either := Left
	if either.IsLeft() {
	    fmt.Println(either.UnwrapLeft()) // Output: 42
	}
*/
type Either[L any, R any] struct {
	left  *L
	right *R
}

/*
Left creates an Either with a left value.

Example:

	either := Left
	if either.IsLeft() {
	    fmt.Println(either.UnwrapLeft()) // Output: 42
	}
*/
func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: &value}
}

/*
Right creates an Either with a right value.

Example:

	either := Right[int, string]("hello")
	if either.IsRight() {
	    fmt.Println(either.UnwrapRight()) // Output: hello
	}
*/
func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: &value}
}

/*
IsLeft returns true if the Either contains a left value.

Example:

	either := Left
	fmt.Println(either.IsLeft()) // Output: true
*/
func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

/*
IsRight returns true if the Either contains a right value.

Example:

	either := Right[int, string]("hello")
	fmt.Println(either.IsRight()) // Output: true
*/
func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

/*
UnwrapLeft returns the left value or panics if the Either contains a right value.

Example:

	either := Left
	fmt.Println(either.UnwrapLeft()) // Output: 42
*/
func (e Either[L, R]) UnwrapLeft() L {
	if e.IsRight() {
		panic("called `UnwrapLeft` on a `Right` value")
	}
	return *e.left
}

/*
UnwrapRight returns the right value or panics if the Either contains a left value.

Example:

	either := Right[int, string]("hello")
	fmt.Println(either.UnwrapRight()) // Output: hello
*/
func (e Either[L, R]) UnwrapRight() R {
	if e.IsLeft() {
		panic("called `UnwrapRight` on a `Left` value")
	}
	return *e.right
}
