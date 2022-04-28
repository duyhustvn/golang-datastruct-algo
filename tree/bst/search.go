package bst

// https://leetcode.com/problems/search-in-a-binary-search-tree
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	currentNode := root
	for currentNode != nil {
		if val < currentNode.Val {
			if leftNode := currentNode.Left; leftNode != nil {
				currentNode = leftNode
			} else {
				return nil
			}
		} else if val > currentNode.Val {
			if rightNode := currentNode.Right; rightNode != nil {
				currentNode = rightNode
			} else {
				return nil
			}
		} else {
			return currentNode
		}
	}
	return nil
}
