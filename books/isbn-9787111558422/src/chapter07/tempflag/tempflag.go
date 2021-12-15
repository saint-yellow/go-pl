package main

import (
	"flag"
	"fmt"
	"isbn-9787111558422/src/chapter07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}