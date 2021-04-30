package main

import (
	"fmt"
	"sort"
)

func sortString(s []string, f func(i, j int) bool) {
	if f == nil {
		// default
		f = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, f)
}

func main() {
	food := []string {
		"onion",
		"carrot",
		"celery",
		"cucumber",
	}

	// default
	var f1 func(i, j int) bool
	sortString(food, f1)
	fmt.Println(food)

	// customized
	var f2 = func(i, j int) bool {
		return food[i] > food[j]
	}
	sortString(food, f2)
	fmt.Println(food)

	// customized
	f3 := func(i, j int) bool {
		return len(food[i]) < len(food[j])
	}
	sortString(food, f3)
	fmt.Println(food)
}