package org

type IDInf interface {
	ID() string
}

type Person struct {
	FirstName string
	LastName string
}


func (p Person) ID() string {
	return "12345"
}