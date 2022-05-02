# Concurrency

## Classic primitives and CSP

* go supports classic and CSP style concurrency primitives
* three main concepts are goroutines, channels and select

## Goroutines

> A "go" statement starts the execution of a function call as an independent
> concurrent thread of control, or goroutine, within the same address space. 

> The function value and parameters are evaluated as usual in the calling
> goroutine, but unlike with a regular call, program execution does not wait for
> the invoked function to complete. Instead, the function begins executing
> independently in a new goroutine. When the function terminates, its goroutine
> also terminates. If the function has any return values, they are discarded
> when the function completes. 

* goroutine is a lightweight thread of execution

[embedmd]:# (x/manygoroutines/main.go)
```go
package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println("Starting...")
	N := 500000
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(1 * time.Second)
		}()
	}
	log.Printf("started %d goroutines, running: %d", N, runtime.NumGoroutine())
}

// $ go run x/manygoroutines/main.go
// 2022/05/02 23:54:44 Starting...
// 2022/05/02 23:54:45 started 500000 goroutines, running: 490075
```


## Channels

## Select
