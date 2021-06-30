package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8000")
	handleError(err)
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, connection)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(connection, os.Stdin)
	connection.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}