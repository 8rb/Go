package main

import "fmt"

var cha chan bool

func process() {

	cha <- true
	fmt.Println("holaa")
}

func main() {
	cha = make(chan bool)

	go process()

	<-cha

	fmt.Println("llego")
}
