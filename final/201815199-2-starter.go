package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var remotehost string
var coordenadas string

func main() {
	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese puerto remoto: (8001, 8004, 8007):")
	port, _ := gin.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port)
	for {
		fmt.Print("Ingrese coordenadas (separadas por espacios): ")
		coordenadas, _ := gin.ReadString('\n')
		fmt.Print("Enviando coordenadas: ", coordenadas)
		sendCapture(coordenadas)
	}
}

func sendCapture(coordenadas string) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", coordenadas)
}
