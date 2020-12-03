package org

import (
	"fmt"
	"strings"
)

// IDInf is interface
type IDInf interface {
	ID() string
}

// SocialSecurityNumber string
type SocialSecurityNumber string

// ID of SSN
func (ssn SocialSecurityNumber) ID() string {
	return "23456"
}

// TwitterHandler type
type TwitterHandler string

// Redirect to twitter
func (th TwitterHandler) Redirect() string {
	cleanHandler := strings.Trim(string(th), "@")
	return fmt.Sprintf("http://www.twitter.com/%s", cleanHandler)
}

// Name struct
type Name struct {
	first string
	last  string
}

// FullName return first, last
func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", strings.Title(n.first), strings.Title(n.last))
}

// Employee struct
type Employee struct {
	Name
}

// Person struct
type Person struct {
	Name
	twitter TwitterHandler
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
		Name{
			first: newFirstName,
			last:  newLastName,
		},
		"",
	}
}

// ID Person implements IDInf
func (p Person) ID() string {
	return "12345"
}
