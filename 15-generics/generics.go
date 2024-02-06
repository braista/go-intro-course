package main

import "fmt"

/*
Generics (Type Parameters)
	Allow variables to refer to specific types.
	Make it possible to write abstract functions that reduce code duplication
*/
func splitAnySlice[T any](s []T) ([]T, []T) { // set T to "any" (empty interface = anything)
	// the function body doesn't care about the types of things stored in the slice
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

/*
	Constraints
		They are just interfaces that allow us to write generics that only operate within the constraint
		of a given interface type.
		Example:
*/
type stringer interface { // interface that will be use as T
	String() string
}

// only elements that implement "stringer" can be used. This is a constraint of T
func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

/*
	Interface Type Lists
		It's a way of writing interfaces reusing existing ones.
		They've to be declared inside an interface and using "|" as separator.
		Example:
*/
type Ordered interface {
	~int | ~int8 | ~int16 // etc... use "|" to separate elements of a interface type list
}

func main() {
	firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3, 4}) // T will be int inside function
	fmt.Println(firstInts, secondInts)
}
