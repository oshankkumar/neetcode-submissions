/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */



func copyRandomList(head *Node) *Node {
	var newHead *Node
	newCurr := &newHead

	oldToNewNodes := make(map[*Node]*Node)
	for curr := head; curr != nil; curr = curr.Next {
		*newCurr = &Node{Val: curr.Val}
		oldToNewNodes[curr] = *newCurr
		newCurr = &((*newCurr).Next)

	}

	for n1, n2 := head, newHead; n1 != nil; n1, n2 = n1.Next, n2.Next {
		n2.Random = oldToNewNodes[n1.Random]
	}

	return newHead
}
