package main

import (
	"errors"
	"fmt"

	"github.com/braista/go-intro-course/5-functions/variadic"
)

/*
Functions:

	can take zero or more arguments
	variable type comes AFTER the name
	for example:
*/
func sub(x int, y int) int { // function signature
	return x - y
}

// Multiple arguments of the same type may only declare it on the last one
func multiArgs(x, y float32) float32 {
	return x + y
}

/*
Multiple returns value
*/
func multiReturn() (x int, y int, z int) {
	return 3, 4, 5
}

/*
Named return values (naked return)

	return values are treated as new variables
	used when you want to document what are the return values meant to be used for

	function below is the same as:
		func nakedReturn() (int, int) {
			var x int
			var y int
			return x, y
		}
*/
func nakedReturn() (x, y int) {
	// x and y are initialized with zero values
	return // automatically returns x and y (implicit return)
}

/*
Early returns

	return early from a function
	can clean up code when used as guard clauses

Guard clause:

	an early return from a function when a given condition is met
*/
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by 0") // early return (guard clause)
	}
	return dividend / divisor, nil
}

/*
Advanced Functions (High-Order Functions)

	HOC are functions that receive other functions as parameters.
	They can be called from inside the function body whenever is necessary.
	The function passed as parameter is called "First-class functions".

	When should I use them?
		Basically whenever you need to run code in the future
		- HTTP API handlers
		- Pub/Sub handlers
		- Callbacks

	"Currying" is a practice where a function that receives a function returns a new function
*/
func hoFunction(num1, num2 int, operation func(num1, num2 int) int) int {
	return operation(num1, num2)
}

func main() {
	result := sub(3, 2) // function invocation
	println(result)

	var number float32 = 2.0
	println(multiArgs(number, 3)) // variable are passed by VALUE (there are exceptions)
	// "number" value cannot be mutated inside "multiArgs"

	// ignore returned value
	response, _, _ := multiReturn() // use "_" to ignore a returned value
	println(response)

	// naked return
	println(nakedReturn())

	// early return
	divisionResult, _ := divide(4, 2)
	println(divisionResult)

	// passing arguments to a variadic function
	variadic.Sum(1, 2, 3, 4) // can pass an undefined number of ints because "..." in its signature

	// calling high-order function
	fmt.Println("hof operation result:", hoFunction(5, 3, sub)) // we pass "sub" function as parameter
}
