package main

import (
	"fmt"
	"play-with-go/typedef/org"
)

func main() {
	p := org.NewPerson("James", "Bond")
	fmt.Println(p.ID())
	fmt.Println(p)

	p.SetTwitter("@james")
	fmt.Println(p)

	fmt.Printf("type of twitterhandler is %T\n", org.TwitterHandler("test"))

}
