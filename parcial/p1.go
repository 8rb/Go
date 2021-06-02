package main

import (
	"fmt"
	"sync"
)

func Reportar(alumnos int, dandoClase sync.Mutex, reportes int, numeroDesconect int) {
	for {
		dandoClase.Lock()
		fmt.Println("Profesor reporta a coordinador")
		reportes++
		if reportes == numeroDesconect {
			fmt.Println("Coordinador reporta a Directora")
			reportes = 0
		}
		dandoClase.Unlock()
	}
}

func Clase(alumnos int, reportar, dandoClase sync.Mutex) {
	for {
		reportar.Lock()
		alumnos++
		if alumnos == 1 {
			dandoClase.Lock()
		}
		reportar.Unlock()
		fmt.Println("Estoy en clase")

		reportar.Lock()
		fmt.Println("Me desconecté")
		alumnos--
		if alumnos == 0 {
			dandoClase.Unlock()
		}
		reportar.Unlock()
	}
}

func main() {
	alumnos := 0
	reportes := 0
	// Cantidad de alumnos que se deben desconectar para reportar a la directora
	numeroDesconect := 10

	reportar := &sync.Mutex{}
	dandoClase := &sync.Mutex{}

	go Clase(alumnos, *reportar, *dandoClase)
	go Clase(alumnos, *reportar, *dandoClase)
	go Clase(alumnos, *reportar, *dandoClase)
	go Reportar(alumnos, *dandoClase, reportes, numeroDesconect)
	Clase(alumnos, *reportar, *dandoClase)
}
