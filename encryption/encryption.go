package main

import "fmt"

func main() {
	str := "This is a string"
	fmt.Println(str)

	encrypted := encrypt(str)
	fmt.Println(encrypted)

	decrypted := decrypt(encrypted)
	fmt.Println(decrypted)

}

func encrypt(s string) string {
	return s
}

func decrypt(s string) string {
	return s
}
