package main


func reverseLeftWords(s string, n int) string {
	r := []rune(s)
	var left int = 0
	var right int = n - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left += 1
		right -= 1
	}

	left = n
	right = len(s) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left += 1
		right -= 1
	}

	left = 0
	right = len(s) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left += 1
		right -= 1
	}

	return string(r)
}