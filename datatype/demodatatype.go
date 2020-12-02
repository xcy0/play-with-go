package main

import (
	"play-with-go/datatype/org"
	"fmt"
)

func main() {
	p := org.Person{}
	fmt.Println(p.ID())

	var p2 org.IDInf = org.Person{}
	fmt.Println(p2.ID())

}
