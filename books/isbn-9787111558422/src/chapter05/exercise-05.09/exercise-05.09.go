package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "saint-yellow $foo"
	fmt.Printf("%s\n", expand(s, f1))
}

func expand(s string, f func(string)string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func f1(s string) string {
	return "f1"
}