package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LinkedList struct
type LinkedList struct {
	value int
	p     *LinkedList
}

func main() {

	rand.Seed(time.Now().UnixNano())

	list := createLinkedList(rand.Int(), nil)

	for i := 0; i < 5; i++ {
		list = createLinkedList(rand.Int(), list)
	}
	showLinkedList(list)
	fmt.Println(list)

}

func showLinkedList(list *LinkedList) {
	for curList := list; curList != nil; {
		fmt.Printf("%v\n", curList.value)
		curList = curList.p
	}
}

func createLinkedList(number int, list *LinkedList) *LinkedList {
	if list == nil {
		newList := &LinkedList{value: number, p: nil}
		return newList
	}

	curList := list
	for curList.p != nil {
		curList = curList.p
	}
	curList.p = &LinkedList{value: number, p: nil}
	return list

}
