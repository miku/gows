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
