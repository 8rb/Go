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

	//lógica del nodo

	if numero == 0 {

		fmt.Println("Fin del proceso, llegó a cero la carga!!!")

	} else {

		enviar(numero - 1)

	}

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

	//establecer el modo escucha del nodo

	ln, _ := net.Listen("tcp", localhost)

	defer ln.Close()

	for {

		//manejador de conexiones

		conn, _ := ln.Accept()

		go manejador(conn) //soportar la carga

	}

}
