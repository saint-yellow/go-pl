package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(connection)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConnection(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	ch := make(chan struct{})
	go func() {
		for input.Scan() {
			ch <- struct{}{}
		}
	}()

	for {
		select {
		case <-time.After(10 * time.Second):
			return
		case <-ch:
			go echo(c, input.Text(), 1 * time.Second)
		}
	}
}