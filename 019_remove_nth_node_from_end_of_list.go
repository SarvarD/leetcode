package leetcode

import . "leetcode/datastructures"

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// First travel n nodes into the list
	leader := head
	for range n {
		leader = leader.Next
	}

	// If n == length of list, we'll reach the end of the list before
	// we can iterate the "follower" pointer. In this case, we just need to
	// remove the head.
	if leader == nil {
		return head.Next
	}

	follower := head
	for leader.Next != nil {
		leader = leader.Next
		follower = follower.Next
	}
	follower.Next = follower.Next.Next

	return head
}
