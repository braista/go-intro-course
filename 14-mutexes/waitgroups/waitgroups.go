package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Wait Groups

	Used to wait for multiple goroutines to finish.
*/
func main() {
	// used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup // declaring a sync.WaitGroup

	// launching several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1) // incrementing WaitGroup counter

		// avoid re-use the same i value in each goroutine closure
		// if not, every goroutine will use the same instance of i, causing bugs or unexpected results
		i := i

		// generating goroutine
		go func() {
			defer wg.Done() // done decrements counter by 1
			worker(i)       // do something
		}() // "i" can be passed here as argument as an alternative to create a new instance
	}

	// blocks until the WaitGroup counter goes back to 0 -> meaning workers are done
	wg.Wait()

	atomicCounters()
}

/*
Atomic Counters

	using sync/atomic package for atomic variables (thread-safe variables)
*/
func atomicCounters() {
	var ops atomic.Uint64 // declaring atomic counter using sync/atomic module
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(i) // incrementing wg counter
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1) // atomically adds (and returns value, but in this example it's not read)
			}
			wg.Done() // decrementing wg counter
		}()
	}
	wg.Wait()                       // await wg to finish
	fmt.Println("ops:", ops.Load()) // atomically loads and return the value
}

// This function will run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
