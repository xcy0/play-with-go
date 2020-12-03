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
}
