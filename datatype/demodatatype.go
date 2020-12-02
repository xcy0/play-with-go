package main

import (
	"play-with-go/datatype/org"
	"fmt"
)

func main() {
	p := org.NewPerson("James", "Bond")
	fmt.Println(p.ID())
	fmt.Println(p)

	var p2 org.IDInf = p
	fmt.Println(p2.ID())

}
