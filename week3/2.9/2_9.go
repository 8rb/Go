package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		// fmt.Println("ptemp = n")
		n = temp + 1
		// fmt.Println("n = ptemp + 1")
		sleepy()
	}
}

func q() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		// fmt.Println("qtemp = n")
		n = temp + 1
		// fmt.Println("n = qtemp + 1")
		sleepy()
	}
}

func main() {
	go p()
	go q()

	time.Sleep(1900 * time.Millisecond)
	fmt.Println(n)
}

func sleepy() {
	time.Sleep(2 * time.Second)
}
