package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type ByteCounter struct {
	writer io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	return &c, &c.written
}

func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Fprintf(w, "Saint-Yellow\n")
	fmt.Println(*c)
}