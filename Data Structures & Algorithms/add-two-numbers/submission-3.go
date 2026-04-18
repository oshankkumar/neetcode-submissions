/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	var carry int
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		curr.Next = &ListNode{Val: val % 10}
		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	remain := l1
	if remain == nil {
		remain = l2
	}

	for remain != nil {
		val := carry + remain.Val
		carry = val / 10
		curr.Next = &ListNode{Val: val % 10}
		curr = curr.Next
		remain = remain.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return head.Next
}