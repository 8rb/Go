package main

import (
	"bufio"

	"fmt"

	"net"

	"os"

	"strconv"

	"strings"
)

//variables globales

var remotehost string

var chCont chan int

var n, min int

func enviar(num int) { //enviar el numero mayor al host remoto

	conn, _ := net.Dial("tcp", remotehost)

	defer conn.Close()

	//envio el número

	fmt.Fprintf(conn, "%d\n", num)

}

func manejador(conn net.Conn) {

	defer conn.Close()

	//recuperar el número

	bufferIn := bufio.NewReader(conn)

	numStr, _ := bufferIn.ReadString('\n')

	numStr = strings.TrimSpace(numStr)

	numero, _ := strconv.Atoi(numStr)

	fmt.Printf("Llegó el número %d\n", numero)

	contador := <-chCont

	//lógica del servicio

	if contador == 0 {

		min = numero

	} else if numero < min {

		//enviar el mayor al nodo siguiente

		enviar(min)

		min = numero //se actualiza el valor del minimo

	} else {

		enviar(numero)

	}

	contador++

	//definir cuando se muestra la respuesta Min

	if contador == n {

		fmt.Printf("El numero es: %d\n", min)

		contador = 0

	}

	chCont <- contador

}

func main() {

	//establecer el identificador del host local (IP:puerto)

	bufferIn := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el puerto local: ")

	puerto, _ := bufferIn.ReadString('\n')

	puerto = strings.TrimSpace(puerto)

	localhost := fmt.Sprintf("localhost:%s", puerto)

	//establecer el identificador del host remoto (IP:puerto)

	fmt.Print("Ingrese el puerto remoto:")

	puerto, _ = bufferIn.ReadString('\n')

	puerto = strings.TrimSpace(puerto)

	remotehost = fmt.Sprintf("localhost:%s", puerto)

	//Cantidad de numero a recibir x nodo

	fmt.Print("Cantidad:")

	nstr, _ := bufferIn.ReadString('\n')

	nstr = strings.TrimSpace(nstr)

	n, _ = strconv.Atoi(nstr)

	//canal para el contador

	chCont = make(chan int, 1) //canal asincrono

	chCont <- 0

	//establecer el modo escucha del nodo

	ln, _ := net.Listen("tcp", localhost)

	defer ln.Close()

	for {

		//manejador de conexiones

		conn, _ := ln.Accept()

		go manejador(conn)

	}

}
