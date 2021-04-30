package main

import "fmt"

 func recovery() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}

func main() {
	defer recovery()

	panic("I forgot my towel!")
}