package main

import "fmt"

func main() {
	temperature := map[string]int {
		"Earth": 15,
		"Mars": -65,
	}

	temp := temperature["Earth"]

	fmt.Printf("On average the Earth is %v °C.\n", temp)

	temperature["Earth"] = 16
	temperature["Mars"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the Moon is %v °C.\n", moon)
	} else {
		fmt.Println("Where is the Moon?")
	}
}