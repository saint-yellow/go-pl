package main

import (
	"fmt"
	"io"
)

type safeWriter struct {
	writer io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return 
	}

	_, sw.err = fmt.Fprintln(sw.writer, s)
}