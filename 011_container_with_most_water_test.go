package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	res := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expected := 49

	if !(res == expected) {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestMaxAreaInCenter(t *testing.T) {
	res := MaxArea([]int{1, 8, 6, 2000, 2000, 5, 4, 8, 3, 7})
	expected := 2000

	if !(res == expected) {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestMaxArea3(t *testing.T) {
	res := MaxArea([]int{2000, 2000, 1, 1})
	expected := 2000

	if !(res == expected) {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
