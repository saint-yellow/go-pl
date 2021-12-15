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

	// 类型断言
	tcpConnection := connection.(*net.TCPConn)

	go func() {
		io.Copy(os.Stdout, connection)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(connection, os.Stdin)

	// 关闭TCP连接的写半边
	tcpConnection.CloseWrite()

	<-done

	// 关闭TCP连接的读半边
	tcpConnection.CloseRead()
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