package main

import "fmt"

func stringsJoin(concactor string, content ...string) string {
	if len(content) == 0 {
		return concactor
	}

	var result string
	for _, c := range content[:len(content)-1] {
		result += c + concactor
	}
	result += content[len(content)-1]

	return result
}

func main() {
	fmt.Println(stringsJoin("-", "saint", "yellow"))
	fmt.Println(stringsJoin("-", "yellow"))
	fmt.Println(stringsJoin("-"))
}