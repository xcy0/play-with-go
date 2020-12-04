package main

import (
	"fmt"
	"math"
)

type intSlice []int

func main() {
	intList := intSlice{24, 46, 65, 86, 97, 123, 241, 376, 651}
	fmt.Println(intList)

	for i := 0; i < len(intList); i++ {
		isPrime := true
		for j := 2; j < int(math.Sqrt(float64(intList[i]))); j++ {
			if intList[i]%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v is a prime number\n", intList[i])
		}
	}
}
