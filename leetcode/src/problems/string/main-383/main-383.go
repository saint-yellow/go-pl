// LeetCode Problem Nr. 383: 赎金信

package main

func canConstruct(ransomNote string, magazine string) bool {
	return method1(ransomNote, magazine)
}

func method1(ransomNote, magazine string) bool {
	table := make([]int, 26)

	for _, s := range magazine {
		table[s-'a']++
	}

	for _, s := range ransomNote {
		if table[s-'a'] == 0 {
			return false
		}

		table[s-'a']--
	}
	return true
}