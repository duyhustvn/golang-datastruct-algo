package single

// Given the head of a singly linked list and two integers left and right where left <= right,
// reverse the nodes of the list from position left to position right, and return the reversed list.

func (l *LinkedList) reverseBetween(left int, right int) {
	if l.head == nil || l.head.next == nil {
		return
	}
	if left == 1 {
		l.reverseBetweenHead(right)
		return
	}

	currentNode := l.head
	var beforeLeft *Node
	i := 1
	for currentNode != nil {
		if i+1 == left {
			beforeLeft = currentNode
			i++
			currentNode = currentNode.next
			continue
		} else if left <= i && i < right {
			tmpNode := beforeLeft.next
			beforeLeft.next = currentNode.next
			if currentNode.next != nil {
				currentNode.next = currentNode.next.next
			}
			beforeLeft.next.next = tmpNode
		} else {
			currentNode = currentNode.next
		}
		i++
	}
}

func (l *LinkedList) reverseBetweenHead(right int) {
	i := 1
	currentNode := l.head
	for currentNode != nil {
		if i < right {
			tmpNode := currentNode.next
			if tmpNode != nil {
				currentNode.next = tmpNode.next
			}
			tmpHead := l.head
			l.head = tmpNode
			tmpNode.next = tmpHead
		} else {
			currentNode = currentNode.next
		}
		i++
	}

}
