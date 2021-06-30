package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
			case gopherID := <-c:
				fmt.Println("gopher ", gopherID, " has finished sleeping")
			case <-timeout:
				fmt.Println("my patience ran out")
				return
		}	
	}
}

func sleepGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id
}