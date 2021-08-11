package main

func isValid(s string) bool {
	mapping := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0)

	for _, i := range s {
		if i == '{' || i == '(' || i == '[' {
			stack = append(stack, i)
		} else if len(stack) > 0 && stack[len(stack)-1] == mapping[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}