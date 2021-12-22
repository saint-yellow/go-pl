package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/saint-yellow/go-pl/toys/tcp-scanner/parameter"
)

var (
	host = parameter.Host
	start = parameter.StartPort
	end = parameter.EndPort
)

func main() {
	var wg sync.WaitGroup
	for i := *start; i <= *end; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", *host, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s is closed\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s is open\n", address)
		}(i)
	}
	wg.Wait()
}