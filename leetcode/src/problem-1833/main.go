package main

import "sort"

func maxIceCream(costs []int, coins int) (count int) {
	sort.Ints(costs)
	for _, c := range costs {
		if c <= coins {
			count += 1
			coins -= c
		}
	}
	return
}