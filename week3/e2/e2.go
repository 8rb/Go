package main

import (
	"fmt"
	"time"
)

func message(cha chan string, text string) {

	cha <- text
}

func main() {

	cha := make(chan string)

	go message(cha, "Saludos")
	time.Sleep(2 * time.Second)
	go message(cha, "Hola")

	for {
		fmt.Println(<-cha)
	}
}
