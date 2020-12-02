package org

// IDInf is interface
type IDInf interface {
	ID() string
}

// Person struct
type Person struct {
	firstName string
	lastName  string
	twitter   string
}

// SetTwitter setter
func (p *Person) SetTwitter(newTwitter string) error {
	p.twitter = newTwitter
	return nil
}

// GetTwitter getter
func (p *Person) GetTwitter() string {
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
