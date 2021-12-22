package main

import (
	"fmt"
	"net"
	"sort"
	"sync"

	"github.com/saint-yellow/go-pl/toys/tcp-scanner/parameter"
)

var (
	host = parameter.Host
	start = parameter.StartPort
	end = parameter.EndPort
)

func worker(ports chan int, results chan []int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", *host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- []int{p, -1}
			continue
		}
		conn.Close()
		results <- []int{p, 0}
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan []int)
	var openPorts []int
	var closedPorts []int
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := *start; i <= *end; i++ {
			ports <- i
		}
	}()

	for i := *start; i <= *end; i++ {
		port := <-results
		if port[1] == 0 {
			openPorts = append(openPorts, port[0])
		} else {
			closedPorts = append(closedPorts, port[0])
		}
	}

	wg.Wait()
	close(ports)
	close(results)

	sort.Ints(openPorts)
	sort.Ints(closedPorts)

	for _, port := range openPorts {
		fmt.Printf("%s:%d is open\n", *host, port)
	}

	for _, port := range closedPorts {
		fmt.Printf("%s:%d is closed\n", *host, port)
	}
}
