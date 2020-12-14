// Package main is the entry point.
package main

import (
	"fmt"
	lib "play-with-go/packagetest/lib/doc.go"
)

func init() {
	fmt.Println("from init")
}
func main() {
	fmt.Println(lib.Aconst)

}
