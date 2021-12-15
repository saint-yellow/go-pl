package main

import "fmt"

func SelectExam1() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	c1 <- 1
	c2 <- 2

	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	}
}

func SelectExam2() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("readable")
	case c<-1:
		fmt.Println("writable")
	}
}

func SelectExam3() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	select {
		case d := <-c:
			fmt.Println(d)
		default:
			fmt.Println("default")
	}
}

func SelectExam4() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	select {
		case d, ok := <-c:
			if !ok {
				fmt.Println("no data received")
				break
			}
			fmt.Println(d)
		default:
			fmt.Println("default")
	}
}

func SelectExam5() {
	select {}
}

func SelectExam6() {
	var c chan string
	select {
	case c <- "Hello":
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}

func main() {
	SelectExam5()
}