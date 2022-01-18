// LeetCode Problem Nr. 383: 赎金信

package main

func canConstruct(ransom string, magazine string) bool {
	return method1(ransom, magazine)
}

func method1(ransom, magazine string) bool {
	table := make([]int, 26)

	for _, s := range magazine {
		table[s-'a']++
	}

	for _, s := range ransom {
		if table[s-'a'] == 0 {
			return false
		}

		table[s-'a']--
	}
	return true
}

func method2(ransom, magazine string) bool {
	mapping := make(map[rune]int)

	for _, s := range magazine {
		if _, ok := mapping[s]; ok {
			mapping[s]++
		} else {
			mapping[s] = 1
		}
	}

	for _, s := range ransom {
		if v, ok := mapping[s]; ok && v > 0 {
			mapping[s]--
		} else {
			return false
		}
	}

	return true
}
