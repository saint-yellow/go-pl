package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(10 * time.Second)
}

func trace(message string) func() {
	start := time.Now()
	log.Printf("enter %s", message)
	return func() {
		log.Printf("exit %s (%s)", message, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}