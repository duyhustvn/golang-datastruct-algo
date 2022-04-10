package bst

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	max := 0
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	DFS(root, 0, &max)
	return max
}

func DFS(root *TreeNode, i int, max *int) {
	if root == nil {
		return
	}
	i++
	if i > *max {
		*max = i
	}

	if root.Left != nil {
		DFS(root.Left, i, max)
	}
	if root.Right != nil {
		DFS(root.Right, i, max)
	}
}
