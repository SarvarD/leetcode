package leetcode

import "testing"

func Test_LongestCommonPrefix(t *testing.T) {
	strs := []string{"hello", "hell", "hi"}
	res := LongestCommonPrefix(strs)
	expected := "h"
	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func Test_LongestCommonPrefixNoCommonPrefix(t *testing.T) {
	strs := []string{"hello", "hell", "goodbye"}
	res := LongestCommonPrefix(strs)
	expected := ""
	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func Test_LongestCommonPrefixFailedLeetcodeCase(t *testing.T) {
	strs := []string{"cir", "car"}
	res := LongestCommonPrefix(strs)
	expected := "c"
	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
