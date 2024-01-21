package main

import "fmt"

/*
	Loops
		"for" is the basic loop in go

		for INITIAL; CONDITION; AFTER {
			// do something
		}

		for sections:
			INITIAL: run once at the beginning. can create variables within the loop scope
			CONDITION: checked before each iteration. if the condition doesn't pass, loop breaks
			AFTER: run after each iteration
*/
func main() {
	// basic loop (for)
	for i := 0; i < 10; i++ { // from 0 to i < 10 (9)
		fmt.Println(i)
	}

	// sections can be omitted
	for i := 0; ; i++ { // omitting for condition
		if i%2 == 1 {
			continue // "continue" keyword is used to skip to the next loop cycle
		} else {
			fmt.Println(i)
		}
		if i == 10 {
			break // "break" keyword is used to break the loop (stop it)
		}
	}

	// for only with condition (while-like)
	// there is no while in go
	condition := true
	count := 0
	for condition { // like java while
		fmt.Println("for", count)
		count++
		if count == 5 {
			condition = false
			fmt.Println("for end")
		}
	}

	// "infinite" loop
	// useful to do things until certain condition is met or for real time things maybe?
	for {
		fmt.Println("infinite loop")
		break // don't forget break lol
	}
}
