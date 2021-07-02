package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := minInt(height[left], height[right]) * (right - left)
	for left < right {
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
		newArea := minInt(height[left], height[right]) * (right - left)
		area = maxInt(area, newArea)
	}
	return area
}

func minInt(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}