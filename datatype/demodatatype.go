package main

import (
	"fmt"
	"play-with-go/datatype/org"
)

func main() {
	p := org.NewPerson("James", "Bond")
	fmt.Println(p.ID())
	fmt.Println(p)

	var p2 org.IDInf = p
	fmt.Println(p2.ID())

}
