package main

import (
	"fmt"
	"time"
)

/*
Concurrency

	Ability to perform multiple tasks at the same time.
	In golang, there is a "go" keyword to execute a function concurrently.

	Example:
		Run concurreny.go and check in the console how the code is executed at the same time
*/
func main() {
	// go keyword generates a new go routine (something like new Thread() in java)
	go doSomething()
	fmt.Println("code below and doSomething() will be executed at the same time")
	for i := 0; i < 10; i++ {
		fmt.Println("main code element", i)
		time.Sleep(2 * time.Second)
	}
}

func doSomething() {
	fmt.Println("something executed concurrently")
	for i := 0; i < 10; i++ {
		fmt.Println("concurrent function element", i)
		time.Sleep(1 * time.Second)
	}
}
