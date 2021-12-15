package main

import (
	"flag"
	"fmt"
	"time"
)

var peroid = flag.Duration("peroid", 1 * time.Second, "sleep peroid")

func main() {
	flag.Parse()

	fmt.Printf("sleeping for %v...", *peroid)
	time.Sleep(*peroid)
	fmt.Println()
}