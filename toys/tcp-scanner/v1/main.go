package main

import (
	"fmt"
	"net"

	"github.com/saint-yellow/go-pl/toys/tcp-scanner/parameter"
)

var (
	host = parameter.Host
	start = parameter.StartPort
	end = parameter.EndPort
)

func main() {
	for i := *start; i <= *end; i++ {
		address := fmt.Sprintf("%s:%d", *host, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s is closed\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s is open\n", address)
	}
}