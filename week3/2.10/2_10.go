package main

import (
	"fmt"
	"time"
)

var n int
var k int = 10

func p() {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		n = temp + 1
	}
}

func q() {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		n = temp - 1
	}
}

func main() {
	go q()
	go p()

	time.Sleep(1990 * time.Millisecond)
	fmt.Println(n)
}
