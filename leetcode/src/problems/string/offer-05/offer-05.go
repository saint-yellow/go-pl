// 剑指 Offer Problem Nr. 5: 替换空格

package main

import "fmt"

func replaceSpace(s string) string {
	return method1(s)
}

func method1(s string) string {
	bytes := []byte(s)
	result := make([]byte, 0)
	for _, b := range bytes {
		if b == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b)
		}
	}
	return string(result)
}

func method2(s string) string {
	bytes := []byte(s)
	length := len(bytes)
	spaceCount := 0
	for _, b := range bytes {
		if b == ' ' {
			spaceCount++
		}
	}

	extendCount := spaceCount * 2
	temp := make([]byte, extendCount)
	bytes = append(bytes, temp...)
	i, j := length - 1, len(bytes) - 1
	for i >= 0 {
		if bytes[i] != ' ' {
			bytes[j] = bytes[i]
			i--
			j--
		} else {
			bytes[j], bytes[j-1], bytes[j-2] = '0', '2', '%'
			i--
			j -= 3
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(replaceSpace("I am so happy."))
	fmt.Println(method2("I am so happy."))
}