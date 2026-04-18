func isValidBST(root *TreeNode) bool {
	return isValidBSTUtil(root, math.MinInt, math.MaxInt)
}

func isValidBSTUtil(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}

	return minVal < root.Val && root.Val < maxVal &&
		isValidBSTUtil(root.Left, minVal, root.Val) &&
		isValidBSTUtil(root.Right, root.Val, maxVal)
}