package leetcode

import "strings"

func ZigZagConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]rune, numRows)
	forward := true
	currRow := 0
	for _, chr := range s {
		rows[currRow] = append(rows[currRow], chr)
		if currRow == (numRows - 1) {
			forward = false
		} else if currRow == 0 {
			forward = true
		}

		if forward {
			currRow += 1
		} else {
			currRow -= 1
		}
	}

	var sb strings.Builder

	for _, row := range rows {
		for _, elem := range row {
			sb.WriteRune(elem)
		}
	}
	return sb.String()
}
