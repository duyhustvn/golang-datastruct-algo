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
	DFS(root, 1, &max)
	return max
}

func DFS(root *TreeNode, i int, max *int) {
	foundLeft := false
	if root.Left != nil {
		foundLeft = true
		i++
		if i > *max {
			*max = i
		}
		DFS(root.Left, i, max)
	}
	if root.Right != nil {
		if !foundLeft {
			i++
			if i > *max {
				*max = i
			}

		}
		DFS(root.Right, i, max)
	}
}
