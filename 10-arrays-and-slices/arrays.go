package main

import (
	"fmt"

	variadicFunctions "github.com/braista/go-intro-course/5-functions/variadic"
)

/*
Arrays

	Fixed-size groups of variables of the same type
	The type [n]T is an array of n values of type T
*/
func main() {
	// declaring an array of 10 integers
	var myIntArray [10]int                      // initialized with zero values
	fmt.Println("declared array\t", myIntArray) // printed as [...]

	// or declaring an initialized  literal
	primes := [...]int{2, 3, 5, 7, 11, 13} // using "..." as size = infere size
	fmt.Println("initialized literal\t", primes)
	fmt.Printf("primes type:\t %T\n", primes)

	// access to an element using the index (position) between square-brackets
	fmt.Println("first prime:\t", primes[0]) // starts from index 0

	// assign a value as an element
	fmt.Println("before assignment:\t", myIntArray[2])
	myIntArray[2] = 5
	fmt.Println("after assignment:\t", myIntArray[2]) // prints 5 (was 0 before)

	// checking array size with "len()"
	fmt.Println("primes length:\t", len(primes))

	/*
		Slices
			Dynamically-sized & flexible view of the elements of an array
			A more common alternative to arrays

			Syntax:
				array[lowIndex:highIndex]	from lowIndex (inclusive) to highIndex (exclusive)
				array[lowIndex:]					from lowIndex (inclusive) to the end
				array[:highIndex]					from start to highIndex (exclusive)
				array[:]									from start to the end (every element)

			IMPORTANT NOTE:
				Slices are NOT copies! they references to the original array
				They don't store any data, it just describes a section of an underlying array
				So if you change an element from the slice, you are changing it on the array too
	*/
	// declaring a slice (zero value is nil)
	var emptySlice []int
	if emptySlice == nil {
		fmt.Println("slice zero value:", emptySlice) // prints "[]", but its zero value is still "nil"
	}

	// creating an slice of "primes" array
	// 	array 	[2, 3, 5, 7, 11, 13]
	// 	index	[0, 1, 2, 3, 4,   5]
	slicedPrimes := primes[1:4] //from index 1 (inclusive) to index 4 (exclusive)
	// element from index "4" is NOT included
	fmt.Println(slicedPrimes)

	// checking slice type
	fmt.Printf("slice type:\t %T\n", slicedPrimes) // prints []int (no fixed size compared to arrays)

	// trying [:]
	fmt.Println(primes[:])

	// changing an element from the slice (and therefore, from the original array too)
	slicedPrimes[0] = 99
	fmt.Println("slice-array", slicedPrimes, primes) // it changed in "slicedPrimes" and "primes"

	// creating a slice literal
	mySliceLiteral := []string{"el1", "el2", "el3"}
	fmt.Println("my slice literal:\t", mySliceLiteral)

	/*
		Make
			built-in function that helps us create and initialize slices, maps and channels
	*/
	// creating a slice using make
	sliceMade := make([]int, 3, 10) // arguments: type, length, capacity (optional)
	fmt.Println(sliceMade)

	// slice capacity
	// a slice has both length and capacity
	// 	length 		-> num of elements contained
	//	capacity 	-> num of elements in the underlying array, counting from the first element in the slice
	fmt.Println("slicedPrimes length:", len(slicedPrimes), "slicedPrimes capacity:", cap(slicedPrimes))

	// spread operator (...)
	// used to pass a slice into a variadic function
	fmt.Println(variadicFunctions.Sum(slicedPrimes...))

	// appending an element to a slice (append)
	//	built-in variadic function used to dynamically add elements to the end of a slice
	//	returns a slice containing all the elements of the original slice plus the appended
	//	if slice has no space, the returned slice will point to the newly allocated array
	sliceMade = append(sliceMade, 69)
	fmt.Println("appended slice:\t", sliceMade)

	// 2d slices (slice of slices)
	matrix := make([][]int, 2)
	fmt.Println("matrix (slice of slices):\t", matrix)
}
