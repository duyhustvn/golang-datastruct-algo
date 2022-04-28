package bst

func sumBST(root *TreeNode) int {
	sum := 0
	sumBSTR(root, &sum)
	return sum
}

func sumBSTR(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	currentNode := root
	*sum += currentNode.Val
	if leftNode := currentNode.Left; leftNode != nil {
		sumBSTR(leftNode, sum)
	}
	if rightNode := currentNode.Right; rightNode != nil {
		sumBSTR(rightNode, sum)
	}
}
