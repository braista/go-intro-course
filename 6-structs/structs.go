package main

import (
	privatePackage "github.com/braista/go-intro-course/6-structs/private" // import alias "privatePackage"

	"fmt"
)

/*
Structs:

	represent structured data.
	they are kinda like classes, but THEY AREN'T classes
	example of a car struct definition:
		("tys" vscode shortcut)
		type Car struct {
			Make string
			Model string
			Height int
			Width int
		}

Struct Naming Convention

	struct name should be capitalized if you want to make it public, otherwise it's private
	same rule applies to struct's variables
	private variables need getters to be accesed from outside
*/
type Car struct {
	privateCode string // private variable
	Make        string // public variable
	Model       string
	Height      int
	Width       int
}

func main() {
	// struct instantiation
	myCar := Car{} // attributes are initialized with "zero values" by default
	fmt.Println(myCar.Make)
	fmt.Println(myCar.privateCode) // private variable can be accesed from the same package
	// assign value to an object's attribute
	myCar.Make = "Tesla"
	fmt.Println(myCar.Make)

	// using a private struct from outside this package
	// myPrivateStruct := privatePackage.privateStruct // WRONG! cant be imported because it's private
	myPublicStruct := privatePackage.PublicStruct{}
	// fmt.Println(myPublicStruct.privateVariable) // WRONG! getter doesn't exist, so it can't be accesed
	fmt.Println(myPublicStruct.PublicVariable) // works!

	/*
		Anonymous structs
			just like a normal struct, but it's defined without a name
			to create one just instantiate the instance immediately using a second pair of brackets
			anonymous structs can be nested

			when should I use it?
				only when you only need to create 1 instance of the struct

			advantages?
				they prevent you from re-using a struct definition you never intended to re-use
	*/
	myBike := struct {
		Make   string
		Model  string
		Engine struct {
			CC string
		}
	}{
		Make:  "Honda", // you can initialize its values this way
		Model: "CBR",
		Engine: struct { // anon struct must be re-declared here if u want to initialize
			CC string
		}{
			CC: "150cc", // trailing comma seems to be mandatory
		},
	}
	fmt.Println(myBike.Make, myBike.Model, myBike.Engine.CC)

	/*
		Embedded vs Nested structs
	*/
	type model struct {
		name string
		year string
	}
	type car struct {
		make  string
		model // model is embedded, so its  definition ("name" and "year") is now part of "car" struct too
	}
	laneTruck := car{
		make: "Toyota",
		/* name: "Corolla", WRONG!
		year: "2024", */
		model: model{ // correct way of initializing an embedded struct
			name: "Corolla",
			year: "2024",
		},
	}
	fmt.Println(laneTruck.name)       // can be accessed directly because they are promoted to the top-level
	fmt.Println(laneTruck.model.name) // or through the embedded reference
	calculateArea()
}

/*
Structs methods

	while go is NOT object oriented, it supports methods which can be defined on structs
	they are just functions that have a receiver: a special parameter that syntactically goes before the name of the function
*/
type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int { // area() has a receiver of (r rectangle)
	return r.width * r.height
}

func calculateArea() {
	r := rectangle{
		width:  5,
		height: 10,
	}
	fmt.Println(r.area())
}
