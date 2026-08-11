/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 

func preorderTraversal(root *TreeNode) []int {
	return PreorderTraversal(root)
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var visited []int
	visited = append(visited, root.Val)
	visited = append(visited, PreorderTraversal(root.Left)...)
	visited = append(visited, PreorderTraversal(root.Right)...)
	return visited
}

