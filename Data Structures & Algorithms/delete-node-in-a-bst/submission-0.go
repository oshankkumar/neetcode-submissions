/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	root.Val = findInOrderSuccessor(root, key)
	root.Right = deleteNode(root.Right, root.Val)
	return root
}

func findInOrderSuccessor(root *TreeNode, key int) int {
	curr := root.Right
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Val
}
