package main

import (
	"fmt"
)

/*
Interfaces:

	Collection of method signatures.
	A type can implement an interface if has all of the methods defined on it
	It never declares that it implements a given interface.. if a type has the proper methods defined,
	then the type automatically fulfills that interface
	Multiple interfaces can be implemented at the same time

	example:
		We define "shape" interface, with area() and perimeter()
		Any type that defines those methods, will implement "shape", and therefore could be treated as one "shape"
*/
type shape interface {
	area() float64
	perimeter() float64
	// name the arguments & return (if needed) so you add readability and clarity (better practice)
	compareArea(other shape) (result bool)
}

type rectangle struct { //rectangle implements shape because has defined area() and perimeter()
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (r rectangle) compareArea(other shape) (result bool) {
	if other == nil {
		return false
	}
	return r.area() == other.area()
}

type circle struct {
	radius float64
}

func main() {
	var rect shape = rectangle{width: 2.5, height: 3}
	fmt.Println("area:", rect.area())
	fmt.Println("perimeter:", rect.perimeter())
	// next line helps to check if a struct (rectangle) implements an interface (shape)
	// throws a compile error if not
	var _ shape = (*rectangle)(nil)

	/*
		Type assertion
			Provides access to an interface value's underlying concrete value (implementation)
			Basically it allows you to extract the actual value stored in an interface variable.
			Useful when you need to cast

			any = interface{} (empty interface) is used to assign whatever.. like typescript "any"

	*/
	var otherVariable any = rect
	c, ok := otherVariable.(circle) // c holds casted object, ok holds if cast was OK or not
	fmt.Println("assertion was ok?", ok, "casted object:", c)

	/*
		Type switches
			allows you to do several type assertions in a series

			("otherVariable" has to be type "interface{}" or "any")
	*/
	switch v := otherVariable.(type) {
	case int:
		fmt.Printf("int block: %T\n", v) // using %T on a Printf prints argument underlying type
	case string:
		fmt.Printf("string block: %T\n", v)
	default:
		fmt.Printf("default block: %T\n", v)
	}

	/*
		Clean interfaces (best practices)
			Some rules of thumb for keeping interfaces clean:

			1. Keep interfaces small
				Interfaces are meant to define the minimal behavior necessary to represent an idea
				Try to not do more than 5 method definitions

			2. Should have no knowledge of satisfying types
				Interfaces should define only what is necessary for other types to classify as a member

				For example:
					type car interface {
						Color() string
						Speed() int
						IsFiretruck() bool // WRONG!
						IsSedan() bool // WRONG!
					}

			3.Interface are NOT classes
					- Don't have constructors or deconstructors
					- They aren't hierarchical by nature
					- They define function signatures, but not behavior
	*/
}
