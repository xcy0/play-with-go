package org

type IDInf interface {
	ID() string
}

type Person struct {
}

func (p Person) ID() string {
	return "12345"
}