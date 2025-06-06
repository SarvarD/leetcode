package leetcode

func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = longestPrefix(prefix, strs[i])
	}
	return prefix
}

func longestPrefix(s1 string, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	maxLen := min(len(runes1), len(runes2))

	matchUpTo := 0
	for idx := range maxLen {
		if runes1[idx] == runes2[idx] {
			matchUpTo++
		} else {
			break
		}
	}
	return string(runes1[:matchUpTo])
}
