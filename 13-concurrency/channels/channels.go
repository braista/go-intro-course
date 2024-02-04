package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Channels

	They are typed, thread-safe queues that allow different goroutines to communicate with
	each other.
	We can put data from one place and read it in another one.

	Data sent through channels is called "tokens"
*/
func main() {
	// creates a channel
	ch := make(chan int)

	// sends data to a channel (blocking operation)
	go func() {
		// this is a blocking operation. it'll block the routine until someone reads the data
		// that's why is executed in a parallel goroutine
		ch <- 69
	}()

	// receive data from a channel (blocking operation)
	value := <-ch // it'll block the routine until someone sends data to the channel
	fmt.Println(value)

	// check example
	example()

	// closing a channel
	// it isn't necessary because there's nothing wrong with leaving channels open
	// only close it to indicate there's nothing else to send
	close(ch) // should always be closed from the sending side

	// checking if a channel is closed
	// sending data to a closed channel will panic
	v, ok := <-ch // ok will be "false" if channel is empty and closed
	fmt.Println("checking if channel is closed:", ok, v)

	// channels can be ranged over
	// will exit loop once channel is closed
	for ele := range ch { // does nothing if ch is closed, but blocks if it's open and empty
		fmt.Println(ele)
	}

	/*
		Channels can optionally be buffered.
		Sending on a buffered channel only blocks when the buffer is full.
		Receiving blocks only when the buffer is empty.
		Useful for limit throughput.
	*/
	buffer := make(chan int, 5) // buffered channels can be created passing limit to make()
	close(buffer)               // inmediately closed because won't use it

	// select
	// listens to multiple channels and process data in the order it comes through each channel
	// if multiple channels are ready at the same time, one is chosen randomly.
	select {
	case e, ok := <-ch:
		fmt.Println(ok, e)
	case <-buffer: // this case doesn't care about value, just fires block below if receives stuff
		fmt.Println("do something")
	default:
		// executes if nothing is received
		// nice way to retry to receive an element after a period of time
		time.Sleep(1 * time.Second)
	}

	/*
		A receive from a nil channel blocks forever
			var c chan string // c is nil
			fmt.Println(<-c) 	// blocks
	*/

	/*
		A send to a nil channel blocks forever
			var c chan string // c is nil
			c <- "whoo" 			// blocks
	*/
}

/*
Channels example

	Code below runs a producer/consumer simulation, where an anonymous function sends data
	through a channel and other code reads it
*/
func example() {
	isOldChan := make(chan bool)

	// concurrent anonymous function
	go func() {
		for _, e := range []string{"old", "old", "new"} {
			// this line below simulates a delay
			time.Sleep(time.Duration(randomIntBetween(1, 5)) * time.Second)
			if e == "old" {
				isOldChan <- true // sends "true" value and block until read
				continue
			}
			isOldChan <- false // sends "false" value and block until read
		}
	}()

	isOld := <-isOldChan // reads and blocks until send
	fmt.Println("element 1 is old:", isOld)
	isOld = <-isOldChan // reads and blocks until send
	fmt.Println("element 2 is old:", isOld)
	isOld = <-isOldChan // reads and blocks until send
	fmt.Println("element 3 is old:", isOld)
}

/* Function that returns a random int between two numbers */
func randomIntBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// readonly channel
func readonly(ch <-chan int) {
	// do something
}

// readonly channel
func writeonly(ch chan<- int) {
	// do something
}
