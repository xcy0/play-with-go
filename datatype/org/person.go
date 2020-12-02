package org

// IDInf is interface
type IDInf interface {
	ID() string
}

// Person struct
type Person struct {
	firstName string
	lastName string
}

// NewPerson constructor 
func NewPerson(newFirstName string, newLastName string) Person {
	return Person{ 
		firstName: newFirstName, 
		lastName: newLastName,
	}
}

// ID Person implements IDInf
func (p Person) ID() string {
	return "12345"
}