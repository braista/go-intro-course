package private

import "fmt"

// can't be accesed from outside this package
type privateStruct struct {
	privateVarible string
}

// can be imported and accesed from outside this package
type PublicStruct struct {
	privateVariable string
	PublicVariable  string
}

// getter definition (horrible)
func (p PublicStruct) PrivateVariable() string {
	return p.privateVariable
}

// setter definition
func (p *PublicStruct) SetPrivateVariable(privateVariable string) {
	p.privateVariable = privateVariable
}

func main() {
	variable := PublicStruct{}
	fmt.Println(variable.privateVariable) // can be accesed because it's the same package
}
