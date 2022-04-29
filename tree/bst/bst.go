package bst

import (
	"ds/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Bst struct {
	Root *TreeNode
}

func InitBST() *Bst {
	return &Bst{}
}

func makeNewNode(v int) TreeNode {
	return TreeNode{Val: v}
}

func (t *Bst) Insert(v int) {
	newNode := makeNewNode(v)
	if t.Root == nil {
		t.Root = &newNode
		return
	}
	currentNode := t.Root
	for currentNode != nil {
		if v < currentNode.Val {
			if currentNode.Left == nil {
				currentNode.Left = &newNode
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &newNode
				break
			}
			currentNode = currentNode.Right
		}
	}
}

func BFS(root *TreeNode) []int {
	result := []int{}
	currentNode := root
	if currentNode == nil {
		return result
	}
	queue := queue.Queue{}
	result = append(result, currentNode.Val)
	queue.Enqueue(currentNode)

	for !queue.IsEmpty() {
		qnode, _ := queue.Dequeue()
		if leftNode := qnode.(*TreeNode).Left; leftNode != nil {
			result = append(result, leftNode.Val)
			queue.Enqueue(leftNode)
		}
		if rightNode := qnode.(*TreeNode).Right; rightNode != nil {
			result = append(result, rightNode.Val)
			queue.Enqueue(rightNode)
		}

	}

	return result
}

func DFSPreOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)

	if root.Left != nil {
		DFSPreOrder(root.Left, result)
	}

	if root.Right != nil {
		DFSPreOrder(root.Right, result)
	}
}

func DFSInOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if leftNode := root.Left; leftNode != nil {
		DFSInOrder(leftNode, result)
	}

	*result = append(*result, root.Val)

	if rightNode := root.Right; rightNode != nil {
		DFSInOrder(rightNode, result)
	}
}

func DFSPostOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if leftNode := root.Left; leftNode != nil {
		DFSPostOrder(leftNode, result)
	}

	if rightNode := root.Right; rightNode != nil {
		DFSPostOrder(rightNode, result)
	}

	*result = append(*result, root.Val)
}
