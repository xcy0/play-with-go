package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stringChannel := make(chan string)
	tower1Channel := make(chan string)
	tower2Channel := make(chan string)

	var offset int32 = 3

	go tower1(stringChannel, tower1Channel, offset)
	go tower2(stringChannel, tower2Channel, -offset)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-tower1Channel:
			fmt.Printf("\nMessage from tower1 - %v\n", msg)
		case msg := <-tower2Channel:
			fmt.Printf("\nMessage from tower1 - %v\n", msg)
		}
	}
}

func tower1(s chan string, t1 chan string, offset int32) {
	inputStream := bufio.NewReader(os.Stdin)
	fmt.Println("Tower1: enter your message for Tower2")
	userInput, _ := inputStream.ReadString('\n')

	userInput = strings.Replace(userInput, "\r\n", "", -1)
	fmt.Printf("Tower1: original message is - %v\n", userInput)

	encrypted := encrypt(userInput, offset)
	fmt.Printf("Tower1: encrypted message - %v\n", encrypted)

	s <- encrypted
	t1 <- "message sent to tower2"
}

func tower2(s chan string, t2 chan string, offset int32) {

	//fmt.Println("Tower2: receiving")

	encrypted := <-s
	fmt.Printf("Tower2: encrypted message - %v\n", encrypted)

	decrypted := encrypt(encrypted, offset)
	fmt.Printf("Tower2: decrypted message - %v\n", decrypted)

	t2 <- "message received"

}

func encrypt(s string, offset int32) string {
	var encrypted string
	for _, c := range s {
		encrypted += string(c + offset)
	}
	return encrypted
}
