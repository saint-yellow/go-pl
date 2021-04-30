package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		connection, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%v is closed.\n", address)
			results <- 0
			continue
		}
		connection.Close()
		fmt.Printf("%v is open!\n", address)
		results <- p
	}
}

func main() {
	start := time.Now()

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts, closePorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closePorts = append(closePorts, port)
		}
		ports <- i
	}
	
	close(ports)
	close(results)

	sort.Ints(openPorts)
	// sort.Ints(closePorts)

	// for _, port := range closePorts {
	// 	fmt.Printf("%d is close.\n", port)
	// }

	for _, port := range openPorts {
		fmt.Printf("%d is open.\n", port)
	}

	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)

}