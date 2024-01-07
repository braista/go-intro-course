package main

/*
	Variables:
		used to store values of a certain type
*/
func main() {
	/*
		Declaration
			when value is not specified, "0" values is used
			zero values from common types:
				int: 			0
				float64:	0.000000
				bool:			false
				string:		""
	*/
	var number int
	println(number) // prints 0

	/*
		Initialization
			value can be set from start of the variable declaration
	*/
	var price float64 = 2.50
	println(price)

	/*
		Short variable declaration:
			variables can be initialized without specifying type using :=
			type will be infered from value
			in this example, assigning the text below as value sets "string" as its type
	*/
	text := "this a string value, so 'text' will be type 'string'"
	println(text)
	logic := true
	println(logic)                     // inferred as "bool"
	mileage, company := 80276, "Tesla" // multiple variables declaration
	println(mileage, company)

	/*
		Unless you have a good reason, it is recommended to stick to the "default" types:
			bool
			string
			int
			uint32
			byte
			rune
			float64
			complex128

		Use more specific types when you want to tweak your program performance and
		memory usage.
	*/

	/*
		The number on the types' names indicates the amount of bites.
		For example:
			uint32 = 32 bits unsigned number
	*/

	/*
		Memory allocation
			variable are pointers to a memory location
			when you create a variable using another one, it make a copy of its value
			(it DOESN'T point at the same location)
	*/
	var original string = "original"
	copied := original // creates a copy (new memory allocation) of "original" value
	println(original, copied)
}
