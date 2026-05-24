/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func kthSmallest(root *TreeNode, k int) int {
	var result int

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil || k == 0 {
			return
		}

		dfs(root.Left)
		k--
		if k == 0 {
			result = root.Val
		}
		dfs(root.Right)
	}

	dfs(root)
	
	return result
}
