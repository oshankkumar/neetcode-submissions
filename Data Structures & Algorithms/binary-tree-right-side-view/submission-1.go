/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	var result []int

	for queue.Len() > 0 {
		n := queue.Len()
		var i int
		for range n {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				result = append(result, node.Val)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			i++
		}
	}

	return result
}
