package datastructures

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(elems []int) *ListNode {
	head := ListNode{}
	curr := &ListNode{}
	head.Next = curr
	for _, value := range elems {
		curr.Next = &ListNode{Val: value}
		curr = curr.Next
	}
	return head.Next.Next
}

func (l1 *ListNode) Equals(l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) || (l1.Val != l2.Val) {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func (l1 *ListNode) ToString() string {
	var sb strings.Builder

	head := l1
	for head != nil {
		sb.WriteString(strconv.Itoa(head.Val))
		sb.WriteString(" ")
		if head.Next != nil {
			sb.WriteString("-> ")
		}
		head = head.Next
	}
	return sb.String()
}
