package main

import "fmt"

/*
	Pointers
		A pointer is a variable that stores the memor address of another variable.
		It points to the location of where the data is stored, NOT de actual data itself.
		Pointer syntax can be used in function arguments to pass them by reference (they are passed as value by default)
*/
func main() {
	// pointer syntax
	var pointer *string // zero value is nil

	// assigning an adress to the pointer variable
	aVariable := "some random variable"
	pointer = &aVariable // it stores the "aVariable" memory address using the & operator
	fmt.Println(pointer) //

	// accessing the value stored in the memory address
	fmt.Println(*pointer) // * operator is used to access to the stored value of the referenced memory address

	// changing the pointee value
	*pointer = "new value" // * operator again to modify the value of the referenced memory address
	fmt.Println("these two will show same output:", aVariable, "|", *pointer)

	// passing variable by reference
	fmt.Println("aVariable old value:", aVariable)
	referenceArgumentsFunction(&aVariable)
	fmt.Println("aVariable new value:", aVariable)
}

/*
	Pointer receivers
		Use the * operator in the argument's type to pass it by reference.

		Mostly used when you pass struct objects to functions
*/
func referenceArgumentsFunction(variable *string) {
	*variable = "change variable value without having to return it"
}
