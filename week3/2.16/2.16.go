package main

import (
	"fmt"
)

var C = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var D = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

func p() {
	var myNumber int
	for i := 0; i < len(C); i++ {
		// P1
		myNumber = C[i]
		// P2
		count := 0
		for j := 0; j < len(C); j++ {
			if C[j] < myNumber {
				count++
			}
		}
		// P3
		if count+1 < len(D) {
			D[count+1] = myNumber
		}

	}
}

func main() {
	p()
	for _, e := range D {
		fmt.Println(e)
	}
}
