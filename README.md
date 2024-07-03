Sure, let's write a comprehensive and engaging `README.md` for the `twoface` package.

---

# twoface 🦇🃏

> Inspired by Two-Face, the Batman villain who embodies duality, `twoface` is a Go package designed to provide powerful functional and concurrency primitives. Just like Two-Face operates with two personalities concurrently, this package helps you manage complex functional and concurrent programming tasks seamlessly.

## Features ✨

- **Option**: Represents optional values, providing a safe way to handle absence.
- **Result**: Represents either success or failure, helping you manage errors gracefully.
- **Either**: Represents a value that can be one of two possible types.
- **Future & Promise**: Handle asynchronous computations with ease.
- **Worker Pool**: Manage a pool of workers for concurrent job processing.
- **Retrier**: Retry logic with customizable strategies.
- **Scaler**: Dynamically scale worker pools based on load.

## Scenarios and Usage Examples 📚

### Option Type

**Scenario**: Use `Option` when a value may or may not be present.

```go
package main

import (
	"fmt"
	"yourmodule/twoface"
)

func main() {
	opt := twoface.Some(42)

	opt.Match(twoface.MatchHandlers[int]{
		Some: func(value int) {
			fmt.Printf("Got a value: %d\n", value)
		},
		None: func() {
			fmt.Println("No value present")
		},
	})
}
```

### Result Type

**Scenario**: Use `Result` to represent operations that can succeed or fail.

```go
package main

import (
	"fmt"
	"yourmodule/twoface"
)

func compute(value int) twoface.Result[int, error] {
	if value < 0 {
		return twoface.Err[int, error](fmt.Errorf("negative value: %d", value))
	}
	return twoface.Ok[int, error](value * 2)
}

func main() {
	result := compute(10)

	result.Match(twoface.MatchHandlers[int, error]{
		Ok: func(value int) {
			fmt.Printf("Success: %d\n", value)
		},
		Err: func(err error) {
			fmt.Printf("Error: %v\n", err)
		},
	})
}
```

### Either Type

**Scenario**: Use `Either` when a value can be one of two types.

```go
package main

import (
	"fmt"
	"yourmodule/twoface"
)

func main() {
	either := twoface.Left

	either.Match(twoface.MatchHandlers[int, string]{
		Left: func(value int) {
			fmt.Printf("Left value: %d\n", value)
		},
		Right: func(value string) {
			fmt.Printf("Right value: %s\n", value)
		},
	})
}
```

### Future and Promise

**Scenario**: Use `Future` and `Promise` to handle asynchronous computations.

```go
package main

import (
	"fmt"
	"time"
	"yourmodule/twoface"
)

func main() {
	promise, future := twoface.NewPromise[int]()

	go func() {
		time.Sleep(2 * time.Second)
		promise.Set(42, nil)
	}()

	value, err := future.Result()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Future result: %d\n", value)
	}
}
```

### Worker Pool

**Scenario**: Use a worker pool to manage concurrent job processing efficiently.

```go
package main

import (
	"context"
	"fmt"
	"yourmodule/twoface"
)

type PrintJob struct {
	message string
}

func (pj PrintJob) Do() twoface.Result[any, error] {
	fmt.Println(pj.message)
	return twoface.Ok[any, error](nil)
}

func main() {
	ctx := context.Background()
	pool := twoface.NewPool(ctx, 5)

	for i := 0; i < 10; i++ {
		job := PrintJob{message: fmt.Sprintf("Job #%d", i)}
		pool.Submit(job)
	}

	pool.Shutdown()
	fmt.Println("All jobs completed.")
}
```

### Retrier

**Scenario**: Use `Retrier` to implement retry logic with customizable strategies.

```go
package main

import (
	"fmt"
	"yourmodule/twoface"
)

type FailingJob struct {
	attempts int
}

func (fj *FailingJob) Do() twoface.Result[any, error] {
	fj.attempts++
	if fj.attempts < 3 {
		return twoface.Err[any, error](fmt.Errorf("attempt %d failed", fj.attempts))
	}
	return twoface.Ok[any, error](nil)
}

func main() {
	job := &FailingJob{}
	retrier := twoface.NewFibonacci(5)
	result := retrier.Do(job)

	result.Match(twoface.MatchHandlers[any, error]{
		Ok: func(_ any) {
			fmt.Println("Job succeeded")
		},
		Err: func(err error) {
			fmt.Printf("Job failed: %v\n", err)
		},
	})
}
```

### Scaler

**Scenario**: Use `Scaler` to dynamically adjust the size of a worker pool based on load.

```go
package main

import (
	"context"
	"fmt"
	"yourmodule/twoface"
)

func main() {
	ctx := context.Background()
	pool := twoface.NewPool(ctx, 5)

	scaler := twoface.NewScaler(pool)
	scaler.Run()

	for i := 0; i < 20; i++ {
		job := twoface.NewJob(func() twoface.Result[any, error] {
			fmt.Printf("Processing job #%d\n", i)
			return twoface.Ok[any, error](nil)
		})
		pool.Submit(job)
	}

	pool.Shutdown()
	fmt.Println("All jobs completed.")
}
```

## License 📜

This project is licensed under the Unlicense.

## Contributing 🤝

Contributions are welcome! Feel free to open an issue or submit a pull request.
