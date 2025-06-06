// https://leetcode.com/problems/add-two-numbers/description/
package leetcode

import . "leetcode/datastructures"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	curr := &ListNode{}
	head.Next = curr
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + carry
		carry = sum / 10
		curr.Val = sum % 10

		if l1 != nil || l2 != nil || carry != 0 {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}
	return head.Next
}
