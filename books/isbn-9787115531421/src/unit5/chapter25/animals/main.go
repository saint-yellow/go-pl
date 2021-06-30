package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sunrise, sunset := 8, 18

	rand.Seed(time.Now().UnixNano())

	animals := []animal {
		honeyBee{name: "Buzz Lightyear"},
		gopher{name: "Go gopher"},
	}

	var sol, hour int

	for {
		fmt.Printf("[%2d:00] ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("The animals are sleeping.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol ++
			if sol >= 3 {
				break
			}
		}
	}
}
