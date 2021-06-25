package main

import (
	"fmt"

	"net"

	"strings"
)

func main() {

	interfs, _ := net.Interfaces()
	for _, interf := range interfs {
		if strings.HasPrefix(interf.Name, "Wi-Fi") {

			fmt.Println(interf.Name)

			direcciones, _ := interf.Addrs()

			for _, direccion := range direcciones {

				fmt.Println(direccion.String())

				switch d := direccion.(type) {

				case *net.IPNet:

					if strings.HasPrefix(d.IP.String(), "192") {

						fmt.Println(d.IP.String())

					}

				}

			}

		}

	}

}
