package main

import (
	"fmt"
	"sync"
)

func A(aArrived sync.Mutex, bArrived sync.Mutex) {
	for {
		fmt.Println("A1 \n")
		aArrived.Unlock()
		bArrived.Lock()
		fmt.Println("A2 \n")
	}
}

func B(aArrived sync.Mutex, bArrived sync.Mutex) {
	for {
		fmt.Println("B1 \n")
		bArrived.Unlock()
		aArrived.Lock()
		fmt.Println("B2 \n")
	}
}

func main() {
	aArrived := &sync.Mutex{}
	bArrived := &sync.Mutex{}
	aArrived.Lock()
	bArrived.Lock()

	go B(*aArrived, *bArrived)
	A(*aArrived, *bArrived)
}
