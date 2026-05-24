/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "slices"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := slices.Index(inorder, root.Val)

	leftInOrder := inorder[:i]
	rigthInOrder := inorder[i+1:]

	leftTreePreOrder := preorder[1 : 1+len(leftInOrder)]
	rightTreePreOrder := preorder[1+len(leftInOrder):]

	root.Left = buildTree(leftTreePreOrder, leftInOrder)
	root.Right = buildTree(rightTreePreOrder, rigthInOrder)

	return root
}
