package bst

import "ds/queue"

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func InitBST() *bst {
	return &bst{}
}

func makeNewNode(v int) node {
	return node{value: v}
}

func (t *bst) Insert(v int) {
	newNode := makeNewNode(v)
	if t.root == nil {
		t.root = &newNode
		return
	}
	currentNode := t.root
	for currentNode != nil {
		if v < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &newNode
				break
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = &newNode
				break
			}
			currentNode = currentNode.right
		}
	}
}

func (t *bst) BFS() []int {
	result := []int{}
	currentNode := t.root
	if currentNode == nil {
		return result
	}
	queue := queue.Queue{}
	result = append(result, currentNode.value)
	queue.Enqueue(currentNode)

	for !queue.IsEmpty() {
		qnode, _ := queue.Dequeue()
		if leftNode := qnode.(*node).left; leftNode != nil {
			result = append(result, leftNode.value)
			queue.Enqueue(leftNode)
		}
		if rightNode := qnode.(*node).right; rightNode != nil {
			result = append(result, rightNode.value)
			queue.Enqueue(rightNode)
		}

	}

	return result
}
