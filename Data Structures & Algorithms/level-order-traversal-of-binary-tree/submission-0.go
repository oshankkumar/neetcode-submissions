/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	var result [][]int
	for queue.Len() > 0 {
		var level []int
		n := queue.Len()
		for range n {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
