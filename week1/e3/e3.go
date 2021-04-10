package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menu := `
		Bienvenido
		Selecciona tu comida:
		[1] picsa
		[2] empanada
		[3] anvorgesa
	`
	fmt.Println(menu)
	bufferIn := bufio.NewReader(os.Stdin)
	input, _ := bufferIn.ReadString('\n')
	optionSTR := strings.TrimRight(input, "\r\n")

	option, _ := strconv.Atoi(optionSTR)
	switch option {
	case 1:
		fmt.Println("elegiste picsa")
	case 2:
		fmt.Println("elegiste empanada")
	default:
		fmt.Println("elegiste anvorgesa")
	}
	fmt.Println("Ingresa la cantidad que quieres")

	var quantity int
	fmt.Scanf("%d", &quantity)
	fmt.Println("cuesta ", quantity*3, " soles")
}
