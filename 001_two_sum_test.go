package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{2, 5, 3, 7, 11, 15}, 9)
	expected := []int{0, 3}

	if !(res[0] == expected[0] && res[1] == expected[1]) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
