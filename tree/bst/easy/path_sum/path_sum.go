package path_sum

import (
	. "ds/tree/bst"
)

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	max := 0
	maxPathSumR(root, sum, &max)
	return max
}

func maxPathSumR(root *TreeNode, sum int, max *int) {
	if root == nil {
		return
	}

	currentNode := root
	sum += currentNode.Val
	if sum > *max {
		*max = sum
	}

	maxPathSumR(currentNode.Left, sum, max)
	maxPathSumR(currentNode.Right, sum, max)
}
