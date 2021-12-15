package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	name    string
	address    string
	message string
}

func main() {
	servers :=  parse(os.Args[1:])

	for _, srv := range servers {
		
		connection, err := net.Dial("tcp", srv.address)
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close()


		go func(s *server) {
			sc := bufio.NewScanner(connection)
			for sc.Scan() {
				s.message = sc.Text()
			}
		}(srv)
	}

	for {
		fmt.Printf("\n")
		for _, server := range servers {
			fmt.Printf("%s (%s): %s\n", server.name, server.address, server.message)
		}
		fmt.Print("--------")

		time.Sleep(time.Second)
	}
}


func parse(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{s[0], s[1], ""})
	}
	return servers
}