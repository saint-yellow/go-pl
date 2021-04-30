package main

import "time"

func main() {
	r := newRoverDriver()
	time.Sleep(3 * time.Second)
	r.left()
	time.Sleep(3 * time.Second)
	r.right()
	time.Sleep(3 * time.Second)
	r.stop()
	time.Sleep(3 * time.Second)
	r.start()
	time.Sleep(3 * time.Second)
}