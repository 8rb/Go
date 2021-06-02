package main

import (
	"fmt"
	"sync"
)

func Genero1(personas1 int, b1, b2, b3 sync.Mutex) {
	for {
		b1.Lock()
		personas1++
		if personas1 == 1 {
			b2.Lock()
			b3.Lock()
		}
		b1.Unlock()
		fmt.Println("Soy género 1 y estoy en el baño")
		b1.Lock()
		fmt.Println("Soy género 1 y ya sali del baño")
		personas1--
		if personas1 == 0 {
			fmt.Println("Ya no hay personas del género 1, puede entrar otro género")
			b2.Unlock()
			b3.Unlock()
		}
		b1.Unlock()
	}
}

func Genero2(personas2 int, b1, b2, b3 sync.Mutex) {
	for {
		b2.Lock()
		personas2++
		if personas2 == 1 {
			b1.Lock()
			b3.Lock()
		}
		b2.Unlock()
		fmt.Println("Soy género 2 y estoy en el baño")
		b2.Lock()
		fmt.Println("Soy género 2 y ya sali del baño")
		personas2--
		if personas2 == 0 {
			fmt.Println("Ya no hay personas del género 2, puede entrar otro género")
			b1.Unlock()
			b3.Unlock()
		}
		b2.Unlock()
	}
}

func Genero3(personas3 int, b1, b2, b3 sync.Mutex) {
	for {
		b3.Lock()
		personas3++
		if personas3 == 1 {
			b1.Lock()
			b2.Lock()
		}
		b3.Unlock()
		fmt.Println("Soy género 3 y estoy en el baño")
		b3.Lock()
		fmt.Println("Soy género 3 y ya sali del baño")
		personas3--
		if personas3 == 0 {
			fmt.Println("Ya no hay personas del género 3, puede entrar otro género")
			b1.Unlock()
			b2.Unlock()
		}
		b3.Unlock()
	}
}

func main() {
	personas1 := 0
	personas2 := 0
	personas3 := 0
	b1 := &sync.Mutex{}
	b2 := &sync.Mutex{}
	b3 := &sync.Mutex{}

	go Genero1(personas1, *b1, *b2, *b3)
	go Genero1(personas1, *b1, *b2, *b3)
	go Genero1(personas1, *b1, *b2, *b3)
	go Genero1(personas1, *b1, *b2, *b3)
	go Genero1(personas1, *b1, *b2, *b3)
	go Genero1(personas1, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero2(personas2, *b1, *b2, *b3)
	go Genero3(personas3, *b1, *b2, *b3)
	go Genero3(personas3, *b1, *b2, *b3)
	go Genero3(personas3, *b1, *b2, *b3)
	go Genero3(personas3, *b1, *b2, *b3)
	go Genero3(personas3, *b1, *b2, *b3)
	Genero3(personas3, *b1, *b2, *b3)
}
