package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

//variables globales
var bitacora []string //Ips de los nodos de la red
const (
	puerto_registro = 8000
	puerto_notifica = 8001
)

var direccionIP_Nodo string

//funciones
func ManejadorNotificacion(conn net.Conn) {
	defer conn.Close()
	//leer la notificación
	bufferIn := bufio.NewReader(conn)
	IpNuevoNodo, _ := bufferIn.ReadString('\n')
	IpNuevoNodo = strings.TrimSpace(IpNuevoNodo)
	//actualizar su bitácora
	bitacora = append(bitacora, IpNuevoNodo)
	fmt.Println(bitacora)
}
func AtenderNotificaciones() {
	//modo escucha
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_Nodo, puerto_notifica)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go ManejadorNotificacion(conn)
	}
}

func RegistrarSolicitud(ipConectar string) {
	hostremoto := fmt.Sprintf("%s:%d", ipConectar, puerto_registro)
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	//enviar la Ip del cliente al host remoto
	fmt.Fprintf(conn, "%s\n", direccionIP_Nodo)
	//leer la bitacora que envia el host remoto
	bufferIn := bufio.NewReader(conn)
	msgBitacora, _ := bufferIn.ReadString('\n')
	var arrAuxiliar []string
	json.Unmarshal([]byte(msgBitacora), &arrAuxiliar)
	bitacora = append(arrAuxiliar, ipConectar) //agregar la ip del host remoto a la bitacora del cliente
	fmt.Println(bitacora)
}

func Notificar(ipremoto, ipNuevoNodo string) {
	hostremoto := fmt.Sprintf("%s:%d", ipremoto, puerto_notifica)
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	//enviar la IP del nodo que se este uniendo a la red
	fmt.Fprintf(conn, "%s\n", ipNuevoNodo)
}

func NotificarTodos(ipNuevoNodo string) {
	//recorrer la bitácora y notificar
	for _, dirIp := range bitacora {
		Notificar(dirIp, ipNuevoNodo)
	}
}

func ManejadorSolicitudes(conn net.Conn) {
	defer conn.Close()
	//leer el IP que envia el nodo a unirse a la red
	bufferIn := bufio.NewReader(conn)
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	//devolvermos al nodo nuevo la bitacora del nodo actual
	//codificar en formato json la bitacora
	bytesBitacora, _ := json.Marshal(bitacora)
	//serializarlo en string
	fmt.Fprintf(conn, "%s\n", string(bytesBitacora)) //enviar respuesta
	//notificar al resto de nodos de la red del nuevo nodo
	NotificarTodos(ip)
	//actualizar la bitacora del nodo actual
	bitacora = append(bitacora, ip)
	fmt.Println(bitacora) //imprimir la bitácora
}

func AtenderSolicitudRegistro() {
	//modo escucha
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_Nodo, puerto_registro)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	//atención concurrente
	for {
		conn, _ := ln.Accept() //aceptar las conexiones
		//manejador
		go ManejadorSolicitudes(conn)
	}
}

func main() {
	direccionIP_Nodo = localAddress()
	fmt.Println("IP: ", direccionIP_Nodo)
	//rol de servidor
	go AtenderSolicitudRegistro()

	//rol de cliente
	//enviar la solicitud de registro
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la ip remota: ")
	ipConectar, _ := bufferIn.ReadString('\n')
	ipConectar = strings.TrimSpace(ipConectar)
	//siempre y cuando no sea el primer nodo de la red
	if ipConectar != "" {
		RegistrarSolicitud(ipConectar)
	}

	//rol de servidor
	AtenderNotificaciones()
}

func localAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
		return "127.0.0.1"
	}
	for _, oiface := range ifaces {

		//for _, dir := range oiface.Addrs() {
		//	fmt.Printf("%v %v\n", oiface.Name, dir)
		//}
		//fmt.Println(oiface.Name)

		if strings.HasPrefix(oiface.Name, "Wi-Fi") {
			addrs, err := oiface.Addrs()
			if err != nil {
				log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
				continue
			}
			for _, dir := range addrs {
				//fmt.Printf("%v %v\n", oiface.Name, dir)
				switch d := dir.(type) {
				case *net.IPNet:
					//fmt.Println(d.IP)
					if strings.HasPrefix(d.IP.String(), "192") {
						//fmt.Println(d.IP)
						return d.IP.String()
					}

				}
			}
		}
	}
	return "127.0.0.1"
}
