/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func kthSmallest(root *TreeNode, k int) int {
	elements := kthSmallestUtil(root, nil)
	return elements[k-1]
}

func kthSmallestUtil(root *TreeNode, acc []int) []int {
	if root == nil {
		return acc
	}

	acc = kthSmallestUtil(root.Left, acc)
	acc = append(acc, root.Val)
	acc = kthSmallestUtil(root.Right, acc)
	return acc
}
