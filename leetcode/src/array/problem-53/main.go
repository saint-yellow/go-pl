package main

func maxSubArray(numbers []int) int {
	max := func (x, y int) int {
		if x < y {
			return y
		} else {
			return x
		}
	}
	pre, ans := 0, numbers[0]
	for _, n := range numbers {
		pre = max(pre+n, n)
		ans = max(ans, pre)
	}
	return ans
}