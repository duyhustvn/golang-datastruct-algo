package bst

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
