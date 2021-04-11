package main

import "fmt"

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("P: SNC LINE 1")
		fmt.Println("P: SNC LINE 2")
		for wantq {

		}
		wantp = true
		fmt.Println("P: SC LINE 1")
		fmt.Println("P: SC LINE 2")
		wantp = false
	}
}

func q() {
	for {
		fmt.Println("Q: SNC LINE 1")
		fmt.Println("Q: SNC LINE 2")
		for wantp {

		}
		wantq = true
		fmt.Println("P: SC LINE 1")
		fmt.Println("P: SC LINE 2")
		wantq = false
	}
}

func main() {

	go p()
	q()
}
