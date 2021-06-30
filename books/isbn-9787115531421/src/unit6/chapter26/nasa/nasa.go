package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christopher J. Scholese"
	administrator = &scolese
	fmt.Println("administrator: ", *administrator)


	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println("administrator: ", *administrator)


	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println("administrator: ", bolden)


	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println("major: ", *major)
	fmt.Println(bolden)
	fmt.Println("administrator: ", *administrator)


	fmt.Println(administrator == major)


	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major)


	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)


	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	fmt.Println(&charles == &bolden)


}