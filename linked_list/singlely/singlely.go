package single

type LinkedList struct {
	head *Node
}

type Node struct {
	value interface{}
	next  *Node
}

func createNewNode(data interface{}) Node {
	return Node{value: data}
}

// Insert node to the end of linked list
func (l *LinkedList) Insert(data interface{}) {
	newNode := createNewNode(data)
	// if the list is empty
	// add node as a head of the list
	if l.head == nil {
		l.head = &newNode
		return
	}
	currentNode := l.head
	// Traverse till the last node
	for currentNode.next != nil {
		currentNode = currentNode.next
		// Add node to the end of the list
	}
	currentNode.next = &newNode
}

func (l *LinkedList) InsertBefore(data interface{}, beforeNode *Node) {
	newNode := createNewNode(data)

	if l.head == nil {
		l.head = &newNode
		return
	}

	if l.head == beforeNode {
		l.head = &newNode
		l.head.next = beforeNode
	}

	for currentNode := l.head; currentNode.next != nil; currentNode = currentNode.next {
		if currentNode.next == beforeNode {
			currentNode.next = &newNode
			newNode.next = beforeNode
			return
		}
	}
}

func (l *LinkedList) InsertAfter(data interface{}, afterNode *Node) {
	newNode := createNewNode(data)
	if l.head == nil {
		l.head = &newNode
		return
	}
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode == afterNode {
			newNode.next = afterNode.next
			afterNode.next = &newNode
			return
		}
	}
}

func (l *LinkedList) Delete(toBeDeletedNode *Node) {
	if l.head == toBeDeletedNode {
		l.head = l.head.next
		return
	}
	for currentNode := l.head; currentNode.next != nil; currentNode = currentNode.next {
		if currentNode.next == toBeDeletedNode {
			currentNode.next = toBeDeletedNode.next
			return
		}
	}
}

func (l *LinkedList) Revert() {
	if l.head == nil || l.head.next == nil {
		return
	}

	var previousNode *Node
	currentNode := l.head

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		l.head = currentNode
		currentNode = nextNode
	}
}

func (l *LinkedList) FindMiddleNode() *Node {
	//if l.head == nil {
	//	return nil
	//}
	//mapIndexToNode := make(map[int]*Node)

	//currentNode := l.head
	//i := 0
	//for currentNode != nil {
	//	i++
	//	mapIndexToNode[i] = currentNode
	//	currentNode = currentNode.next
	//}

	//middleIndex := i/2 + 1
	//return mapIndexToNode[middleIndex]

	if l.head == nil || l.head.next == nil {
		return l.head
	}

	slowNode := l.head
	fastNode := l.head.next
	for fastNode != nil {
		slowNode = slowNode.next
		fastNode = fastNode.next
		if fastNode != nil {
			fastNode = fastNode.next
		}
	}
	return slowNode
}
