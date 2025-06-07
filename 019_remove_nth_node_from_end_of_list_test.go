package leetcode

import (
	. "leetcode/datastructures"
	"testing"
)

func TestSecondLast(t *testing.T) {
	list := NewLinkedList([]int{1, 2, 3, 4, 5})
	res := RemoveNthFromEnd(list, 2)
	expected := NewLinkedList([]int{1, 2, 3, 5})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestListLength(t *testing.T) {
	list := NewLinkedList([]int{1, 2, 3, 4, 5})
	res := RemoveNthFromEnd(list, 5)
	expected := NewLinkedList([]int{2, 3, 4, 5})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestListSingleElement(t *testing.T) {
	list := NewLinkedList([]int{1})
	res := RemoveNthFromEnd(list, 1)
	expected := NewLinkedList([]int{})

	if !(res.Equals(expected)) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
