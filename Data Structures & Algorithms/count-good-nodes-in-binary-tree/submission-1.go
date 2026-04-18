func goodNodes(root *TreeNode) int {
	return gooNodesUtil(root, root.Val)

}

func gooNodesUtil(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	maxVal = max(root.Val, maxVal)

	c1 := gooNodesUtil(root.Left, maxVal)
	c2 := gooNodesUtil(root.Right, maxVal)

	if root.Val >= maxVal {
		return 1 + c1 + c2
	}
	return c1 + c2
}
