package main

import (
	"fmt"
	"time"
)

var n int = 0

func p() {
	temp := n
	n = temp + 1
}

func q() {
	temp := n
	n = temp + 1
}

func main() {
	go p()

	time.Sleep(2 * time.Second)

	go q()

	time.Sleep(2 * time.Second)

	fmt.Println(n)
}
