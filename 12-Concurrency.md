# Concurrency

## Classic primitives and CSP

* go supports classic and CSP style concurrency primitives
* three main concepts are goroutines, channels and select statement

## Runtime Scheduler

* the go runtime implements a scheduler, akin to an OS to manage goroutines
* there is always one goroutine running (i.e. main thread)
* M:N scheduler, utilizes multiple cores

> At any time, M goroutines need to be scheduled on N OS threads that runs on at
> most GOMAXPROCS numbers of processors. -- https://rakyll.org/scheduler/

* there are some guaruantee, but in general you can determine the scheduling
  order - not should you rely on it

Just as in an OS, goroutines will be in different states:

* running
* runnable
* blocked

Blocking can occur through:

* Sending and Receiving on Channel.
* Network I/O.
* Blocking System Call.
* Timers.
* Mutexes.

Previously, a variable `GOMAXPROCS` used to be 1, which would only use one OS
thread. It now defaults to the number of CPUs.

> The GOMAXPROCS variable limits the number of operating system threads that can
> execute user-level Go code simultaneously. -- https://pkg.go.dev/runtime

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
* a goroutine consumes about 2kb of memory
* you can start 1 million goroutines in about 2 seconds (and they'd consume 1/5
  of the RAM of an 16GB machine)

[embedmd]:# (x/manygoroutines/main.go)
```go
package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

var (
	N     = flag.Int("n", 10, "number of goroutines to start")
	sleep = flag.Duration("s", 100*time.Millisecond, "how long each goroutine sleeps")
)

func main() {
	flag.Parse()
	log.Printf("starting %d goroutines", *N)
	started := time.Now()
	for i := 0; i < *N; i++ {
		go f()
	}
	log.Printf("%d/%d goroutines started/running in %v", *N, runtime.NumGoroutine(), time.Since(started))
}

func f() {
	time.Sleep(*sleep)
}
```

* there is no outside control, you cannot "stop" a goroutine
* a goroutine will stop if the function it started ends
* a goroutine may run on one thread or many different threads (see: [LockOSThread](https://pkg.go.dev/runtime#LockOSThread))

### How goroutines may end

* function done successfully
* function failed and returned (there is no return value captured by the caller)
* function hangs, endless loop, resource or goroutine leak

## Channels

* channels are typed conduits between goroutines
* can be unidirectional or bidirectional
* can be unbuffered or buffered

More on channels:

* Use channels to orchestrate and coordinate goroutines.
    * Focus on the signaling semantics and not the sharing of data.
    * Signaling with data or without data.
    * Question their use for synchronizing access to shared state.
        * _There are cases where channels can be simpler for this but initially question._
* Unbuffered channels:
    * Receive happens before the Send.
    * Benefit: 100% guarantee the signal being sent has been received.
    * Cost: Unknown latency on when the signal will be received.
* Buffered channels:
    * Send happens before the Receive.
    * Benefit: Reduce blocking latency between signaling.
    * Cost: No guarantee when the signal being sent has been received.
        * The larger the buffer, the less guarantee.
        * Buffer of 1 can give you one delayed send of guarantee.
* Closing channels:
    * Close happens before the Receive. (like Buffered)
    * Signaling without data.
    * Perfect for signaling cancellations and deadlines.
* NIL channels:
    * Send and Receive block.
    * Turn off signaling
    * Perfect for rate limiting or short-term stoppages.



## Select
