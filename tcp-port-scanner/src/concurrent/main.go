package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scan(host string, fromPort, toPort int) {
	start := time.Now()
	var wg sync.WaitGroup
	for i := fromPort; i <= toPort; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%v:%d", host, j)
			connection, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%v is closed.\n", address)
				return
			}
			connection.Close()
			fmt.Printf("%v is open!\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}

func main() {
	scan("127.0.0.1", 21, 65535)
}