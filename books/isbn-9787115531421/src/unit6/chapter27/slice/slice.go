package main

import "fmt"

func main() {
	var soup []string
	fmt.Println(soup == nil)

	// safe
	// for _, ingredient := range soup {
	// 	fmt.Println(ingredient)
	// }

	fmt.Println(len(soup))

	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
}