package main

import (
	"fmt"
	"time"
)

/*
Timers

	Used to execute code once at some point in the future.
*/
func timers() {
	// creating a 2s timer
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // awaits & blocks until timer completion
	fmt.Println("Timer 1 fired")

	// timers can be canceled before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// Returns true if the call stops the timer,
	// false if the timer has already expired or been stopped
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

/*
Tickers

	Used to do something repeatedly at regular intervals, until told to stop.
*/
func tickers() {
	// creating a 500ms ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // channel to make it stop

	// ticking goroutine
	go func() {
		for {
			select {
			case <-done: // executed when done
				return
			case t := <-ticker.C: // executed every 500ms
				fmt.Println("Tick at:", t)
			}
		}
	}()
	time.Sleep(2 * time.Second) // delay to let ticker work for a while
	ticker.Stop()               // stopping ticker using Stop()
	done <- true                // used to exit goroutine (if not deadlock)
	fmt.Println("Ticker stopped")
}

func main() {
	timers()
	tickers()
}
