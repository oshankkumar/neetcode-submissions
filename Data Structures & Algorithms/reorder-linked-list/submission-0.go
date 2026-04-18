func reorderList(head *ListNode) {
	var nodes []*ListNode
	for curr := head; curr != nil; curr = curr.Next {
		nodes = append(nodes, curr)
	}

	curr, i := head, len(nodes)-1
	for curr != nodes[i] && curr.Next != nodes[i] {
		temp := curr.Next
		curr.Next = nodes[i]
		nodes[i].Next = temp
		curr = temp
		i--
	}
	nodes[i].Next = nil
}
