package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//list := createMap(100000)
	toFind := rand.Intn(10000)
	fmt.Println(toFind)
	count := 0
	for {
		count++
		if toFind == rand.Intn(10000) {
			fmt.Printf("found on the %v\n", count)
			break
		}
	}
	/*
		for index := 0; ; index++ {
			if list[index] == toFind {
				fmt.Println(index)
				break
			}
		}
	*/

}

func createMap(count int) map[int]int {

	list := make(map[int]int)
	for i := 0; i < count; i++ {
		list[i] = rand.Intn(10000)
	}
	return list
}
