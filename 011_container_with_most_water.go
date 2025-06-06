package leetcode

func MaxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxVal := min(height[start], height[end]) * (end - start)
	for start < end {
		if height[end] >= height[start] {
			start += 1
		} else {
			end -= 1
		}
		currVal := min(height[start], height[end]) * (end - start)
		maxVal = max(currVal, maxVal)
	}

	return maxVal
}
