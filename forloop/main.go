package main

import "fmt"

// LinkedList struct
type LinkedList struct {
	value int
	p     *LinkedList
}

func main() {

	list := createLinkedList(10, nil)
	fmt.Println(list)

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
