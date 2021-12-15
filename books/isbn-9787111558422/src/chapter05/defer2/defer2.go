package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buffer [4096]byte
	n := runtime.Stack(buffer[:], false)
	os.Stdout.Write(buffer[:n])
}

func f(x int) {
	if x < 0 {
		return
	}
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}