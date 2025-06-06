package leetcode

import (
	"strconv"
	"strings"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return CountAndSayRepr(CountAndSay(n - 1))
}

func CountAndSayRepr(s string) string {
	var sb strings.Builder
	currVal := s[0]
	currCount := 1
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == currVal {
			currCount += 1
		} else {
			sb.WriteString(strconv.Itoa(currCount))
			sb.WriteByte(currVal)
			currVal = s[idx]
			currCount = 1
		}
	}
	sb.WriteString(strconv.Itoa(currCount))
	sb.WriteByte(currVal)
	return sb.String()
}
