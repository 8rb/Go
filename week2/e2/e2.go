package main

import (
	"fmt"
	"time"
)

var n int = 0

func p() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		n = temp + 1
	}
}

func main() {
	go p()
	go p()

	time.Sleep(2 * time.Second)

	fmt.Println(n)
}
