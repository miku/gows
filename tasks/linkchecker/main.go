package main

import (
	"log"
)

// Let's write a program that takes a list of URLs and retrieves them in
// parallel. We want to check the status code of each link.
//
// The tasks may ask to rewrite parts of the program.
//
// ----------------

// 1. Write a non-concurrent version, first - using, e.g. http.Get

// 2. Start each check in a goroutine.

// 3. Use a WaitGroup to wait for all goroutines to finish.

// 4. Implement a worker pool style processing: Start a fixed number of
// "workers" and create a channel (e.g. of type string) to send a URL to the
// worker. The worker function can perform the link check and print out the
// results. You can iterate over incoming "work" with a for loop.

// 5. Create a fan-in function, that collect all results from the worker
// functions: The worker does not print out the result any more, but sends it
// to a channel to a function that can then output the results.

func main() {
	urls := []string{
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.qq.com",
		"http://www.google.com",
		"http://www.twitter.com",
		"http://www.live.com",
	}
	log.Println("checking %d urls", len(urls))
}
