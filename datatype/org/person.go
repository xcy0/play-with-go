package org

type IDInf interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
}

func NewPerson(newFirstName string, newLastName string) Person {
	return Person{ 
		firstName: newFirstName, 
		lastName: newLastName,
	}
}


func (p Person) ID() string {
	return "12345"
}