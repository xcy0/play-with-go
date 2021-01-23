package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(add(1, 2))

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Microsecond * 10)
	}
	time.Sleep(time.Second)

}

func add(p1 float64, p2 float64) float64 {
	return p1 + p2
}
