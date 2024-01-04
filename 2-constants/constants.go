package main

/*
	Constants
		they store values, like variables, but this cannot be changed after declaration
		declared using "const" keyword
		can't be declared using short syntax
		value must be set from start
*/
func main() {
	const stringConstant = "text that cannot be modified"
	println(stringConstant)
	//stringConstant = "this re-asignation gives compilation error"
}
