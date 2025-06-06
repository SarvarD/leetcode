package leetcode

import (
	. "leetcode/datastructures"
	"testing"
)

func TestAddTwo(t *testing.T) {
	num1 := NewLinkedList([]int{1, 2, 3})
	num2 := NewLinkedList([]int{2, 2, 4})
	res := addTwoNumbers(num1, num2)
	expected := NewLinkedList([]int{3, 4, 7})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected.ToString(), res.ToString())
	}
}

func TestAddTwoWithCarry(t *testing.T) {
	num1 := NewLinkedList([]int{9, 9})
	num2 := NewLinkedList([]int{4, 3})
	res := addTwoNumbers(num1, num2)
	expected := NewLinkedList([]int{3, 3, 1})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected.ToString(), res.ToString())
	}
}

func TestAddTwoDifferentLength(t *testing.T) {
	num1 := NewLinkedList([]int{9, 9})
	num2 := NewLinkedList([]int{4})
	res := addTwoNumbers(num1, num2)
	expected := NewLinkedList([]int{3, 0, 1})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected.ToString(), res.ToString())
	}
}
