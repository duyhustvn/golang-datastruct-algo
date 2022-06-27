package priorityqueue

import "fmt"

type Node struct {
	Value    interface{}
	Priority int
	Next     *Node
}

type PriorityQueue struct {
	Front *Node
	Rear  *Node
	IsMin bool
}

func NewPriorityQueue(isMin bool) *PriorityQueue {
	return &PriorityQueue{IsMin: isMin}
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Front == nil && pq.Rear == nil
}

func newNode(v interface{}, priority int) *Node {
	return &Node{Value: v, Priority: priority}
}

func (pq *PriorityQueue) Enqueue(v interface{}, priority int) {
	node := newNode(v, priority)

	if pq.Front == nil && pq.Rear == nil {
		pq.Front = node
		pq.Rear = node
		return
	}

	// Add to the head if priority greater than priority of front node
	// If there is one node in queue, assign
	curNode := pq.Front
	if (node.Priority > curNode.Priority && !pq.IsMin) || (node.Priority <= curNode.Priority && pq.IsMin) {
		node.Next = curNode
		pq.Front = node

		if pq.Front == pq.Rear {
			pq.Rear = curNode
		}

		return
	}

	var previousNode *Node
	for curNode != nil {
		if (node.Priority > curNode.Priority && !pq.IsMin) || (node.Priority <= curNode.Priority && pq.IsMin) {
			node.Next = curNode
			previousNode.Next = node
			return
		}
		previousNode = curNode
		curNode = curNode.Next
	}

	pq.Rear.Next = node
	pq.Rear = node
}

func (pq *PriorityQueue) Dequeue() (*Node, bool) {
	if pq.IsEmpty() {
		return nil, false
	}
	if pq.Front == pq.Rear {
		node := pq.Front
		pq.Front = nil
		pq.Rear = nil
		return node, true
	}

	node := pq.Front
	pq.Front = pq.Front.Next

	return node, true
}

func (pq *PriorityQueue) ToString(typ string) string {
	str := ""
	curNode := pq.Front
	for curNode != nil {
		switch typ {
		case "int":
			str += fmt.Sprintf("{%d %d} ->", curNode.Value.(int), curNode.Priority)
			break
		case "string":
			str += fmt.Sprintf("{%s %d} ->", curNode.Value.(string), curNode.Priority)
			break
		}

		curNode = curNode.Next
	}
	return str
}
