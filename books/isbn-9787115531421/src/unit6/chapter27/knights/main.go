package main

import "fmt"

func main() {
	arthur := &character {
		name: "Arthur",
	}

	shrubbery := &item {
		name: "shrubbery",
	}

	fmt.Println(arthur)
	arthur.pickup(shrubbery)
	fmt.Println(arthur)

	lancelot := &character {
		name: "Lancelot",
	}
	fmt.Println(lancelot)

	arthur.give(lancelot)

	fmt.Println(arthur)
	fmt.Println(lancelot)
}