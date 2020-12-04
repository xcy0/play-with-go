package main

import "fmt"

type intSlice []int

func main() {
	intList := intSlice{24, 46, 65, 86, 97, 123, 241, 376, 651}
	fmt.Println(intList)
	for i := 0; i < len(intList); i++ {

	}
}
