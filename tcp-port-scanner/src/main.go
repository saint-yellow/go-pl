package main

import (
	"fmt"
	"net"
)

func scan(host string, fromPort, toPort int) {
	for i := fromPort; i <= toPort; i++ {
		address := fmt.Sprintf("%v:%d", host, i)
		connection, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%v is closed.\n", address)
			continue
		}
		connection.Close()
		fmt.Printf("%v is open!\n", address)
	}
}

func main() {
	scan("127.0.0.1", 21, 120)
}