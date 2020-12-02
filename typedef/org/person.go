package org

import (
	"fmt"
	"strings"
)

// IDInf is interface
type IDInf interface {
	ID() string
}

// TwitterHandler type
type TwitterHandler string

// Redirect to twitter
func (th TwitterHandler) Redirect() string {
	cleanHandler := strings.Trim(string(th), "@")
	return fmt.Sprintf("http://www.twitter.com/%s", cleanHandler)
}

// Person struct
type Person struct {
	firstName string
	lastName  string
	twitter   TwitterHandler
}

// SetTwitter setter
func (p *Person) SetTwitter(newTwitter TwitterHandler) error {
	p.twitter = newTwitter
	return nil
}

// GetTwitter getter
func (p *Person) GetTwitter() TwitterHandler {
	return p.twitter
}

// NewPerson constructor
func NewPerson(newFirstName string, newLastName string) Person {
	return Person{
		firstName: newFirstName,
		lastName:  newLastName,
	}
}

// ID Person implements IDInf
func (p Person) ID() string {
	return "12345"
}
