package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Alienigena struct {
	Longitud float64 `json:"longitud"`
	Latitud  float64 `json:"latitud"`
	Nombre   string  `json:"nombre"`
}

type Aliado struct {
	Latitud  float64 `json:"latitud"`
	Longitud float64 `json:"longitud"`
	Nombre   string  `json:"nombre"`
}

var remotehost string

var allyList []string

var notificationPort string
var listenPort string
var processPort string

var ally Aliado

func initAlly() {
	fmt.Println("Initializing ally")
	ally.Nombre = "Aliado"
	ally.Longitud = 12.0
	ally.Latitud = 15.0
	fmt.Print("Inicializado el aliado:", ally)
}

func registerAllyLocally(newAllyAddress string) {
	allyList = append(allyList, newAllyAddress)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresar dirección: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)
	fmt.Println(port)
	initAlly()
	portNum, _ := strconv.Atoi(port)
	notificationPort = strconv.Itoa(portNum)
	listenPort = strconv.Itoa(portNum + 1)
	processPort = strconv.Itoa(portNum + 2)
	fmt.Println("LOS PUERTOS:", notificationPort, listenPort, processPort)
	registerAllyLocally(notificationPort)
	fmt.Print("Ingresar dirección aliado cercano: ")
	port, _ = reader.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port)
	if port != "" {
		fmt.Println("Registrando nuevo aliado")
		registerAllyLocally(port)
		for i := 0; i < len(allyList); i++ {
			notify(allyList[i], notificationPort)
		}
	}
	fmt.Println(allyList)
	for {
		go handleNotifications()
		handleCaptureProcess()
	}
}

func handleCaptureProcess() {
	listenHost := fmt.Sprintf("localhost:%s", listenPort)
	ln, _ := net.Listen("tcp", listenHost)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go captureManager(conn)
	}
}

func handleNotifications() {
	notificationHost := fmt.Sprintf("localhost:%s", notificationPort)
	ln, _ := net.Listen("tcp", notificationHost)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go notificationManager(conn)
	}
}

func captureManager(conn net.Conn) {
	fmt.Println("Recibí captura")
	defer conn.Close()
	bufferIn := bufio.NewReader(conn)
	coordenadas, _ := bufferIn.ReadString('\n')
	coordenadas = strings.TrimSpace(coordenadas)
	fmt.Println("Coordenadas de alienigena:", coordenadas)
}

func notificationManager(conn net.Conn) {
	fmt.Println("Recibí notificación")
	defer conn.Close()
	bufferIn := bufio.NewReader(conn)
	newPort, _ := bufferIn.ReadString('\n')
	newPort = strings.TrimSpace(newPort)
	registered := false
	for i := 0; i < len(allyList); i++ {
		if allyList[i] == newPort {
			registered = true
		}
	}
	if registered == true {
		fmt.Println(allyList)
	} else {
		fmt.Println("Registrando nuevo aliado")
		allyList = append(allyList, newPort)
		for i := 0; i < len(allyList); i++ {
			if allyList[i] != newPort {
				notify(newPort, allyList[i])
			}
		}
		fmt.Println(allyList)
	}
}

func notify(remotePort string, newPort string) {
	if remotePort != newPort {
		fmt.Println("Notificando aliado", remotePort, newPort)
		recieverHost := fmt.Sprintf("localhost:%s", remotePort)
		conn, _ := net.Dial("tcp", recieverHost)
		defer conn.Close()
		fmt.Fprintf(conn, "%s\n", newPort)
	}
}

func send(num int, port string) {
	recieverHost := fmt.Sprintf("localhost:%s", port)
	conn, _ := net.Dial("tcp", recieverHost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num)
}
