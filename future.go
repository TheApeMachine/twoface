package twoface

import (
	"sync"
)

/*
Future represents a value that will be available at some point in the future.

Example:

f := NewFuture[int]()
result, err := f.Result()
*/
type Future[T any] struct {
	result T
	err    error
	done   chan struct{}
	mu     sync.Mutex
}

/*
NewFuture creates a new Future.

Example:

f := NewFuture[int]()
*/
func NewFuture[T any]() *Future[T] {
	return &Future[T]{
		done: make(chan struct{}),
	}
}

/*
Result blocks until the future value is available and returns it.

Example:

f := NewFuture[int]()
result, err := f.Result()
*/
func (f *Future[T]) Result() (T, error) {
	<-f.done
	return f.result, f.err
}

/*
Then adds a handler to be called when the Future is completed successfully.

Example:

f := NewFuture[int]()
f.Then(func(result int) { fmt.Println(result) })
*/
func (f *Future[T]) Then(handler func(T)) *Future[T] {
	go func() {
		<-f.done
		f.mu.Lock()
		defer f.mu.Unlock()
		if f.err == nil {
			handler(f.result)
		}
	}()
	return f
}

/*
Catch adds a handler to be called if the Future is completed with an error.

Example:

f := NewFuture[int]()
f.Catch(func(err error) { fmt.Println(err) })
*/
func (f *Future[T]) Catch(handler func(error)) *Future[T] {
	go func() {
		<-f.done
		f.mu.Lock()
		defer f.mu.Unlock()
		if f.err != nil {
			handler(f.err)
		}
	}()
	return f
}

/*
Finally adds a handler to be called when the Future is completed, regardless of success or failure.

Example:

f := NewFuture[int]()
f.Finally(func() { fmt.Println("Completed") })
*/
func (f *Future[T]) Finally(handler func()) *Future[T] {
	go func() {
		<-f.done
		f.mu.Lock()
		defer f.mu.Unlock()
		handler()
	}()
	return f
}

/*
Promise represents a writable, single-assignment container for a Future.

Example:

p, f := NewPromise[int]()
p.Set(42, nil)
result, err := f.Result()
*/
type Promise[T any] struct {
	future *Future[T]
	once   sync.Once
}

/*
NewPromise creates a new Promise and its associated Future.

Example:

p, f := NewPromise[int]()
*/
func NewPromise[T any]() (*Promise[T], *Future[T]) {
	future := NewFuture[T]()
	return &Promise[T]{future: future}, future
}

/*
Set sets the value of the Future, completing it.

Example:

p, f := NewPromise[int]()
p.Set(42, nil)
*/
func (p *Promise[T]) Set(result T, err error) {
	p.once.Do(func() {
		p.future.mu.Lock()
		defer p.future.mu.Unlock()
		p.future.result = result
		p.future.err = err
		close(p.future.done)
	})
}
