package bst

func min(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minV := root.Val
	minR(root, &minV)
	return minV
}

func minR(root *TreeNode, minV *int) {
	if root == nil {
		return
	}

	currentNode := root
	if *minV > currentNode.Val {
		*minV = currentNode.Val
	}

	minR(currentNode.Left, minV)
	minR(currentNode.Right, minV)
}
