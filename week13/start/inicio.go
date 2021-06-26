package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var remotehost string

func enviar(num int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	//enviar el numero
	fmt.Fprintf(conn, "%d\n", num)
}

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la IP remota: ")
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	remotehost = fmt.Sprintf("%s:%d", ip, 8002)

	//enviar los datos al nodo 1 (8000)

	enviar(18)

	//enviar(35)
	// enviar(10)
	// enviar(25)
	// enviar(15)
	// enviar(30)

}
