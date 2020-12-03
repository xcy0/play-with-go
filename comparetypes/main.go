package main

import (
	"fmt"
	"play-with-go/embed/org"
)

func main() {

	idInterface := org.NewSocialSecurityNumber("111-222-333-444")
	p := org.NewPerson("James", "Bond", idInterface)
	fmt.Println(p.ID())
	fmt.Println(p)

	p.SetTwitter("@james")
	fmt.Println(p)
	fmt.Println(p.GetTwitter().Redirect())

	fmt.Printf("type of twitterhandler is %T\n", org.TwitterHandler("test"))

	name1 := Name{First: "f", Last: "S"}
	name2 := Name{First: "f", Last: "S"}

	if name1 == name2 {
		fmt.Println("equal")
	}

}

// Name struct
type Name struct {
	First string
	Last  string
	// Middle []string
}

// OtherName struct
type OtherName = Name
