package single

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func (linkedList *LinkedList) AddNode(v interface{}) {
	newNode := Node{
		Value: v,
		Next:  nil,
	}
	if linkedList.Head == nil {
		linkedList.Head = &newNode
		linkedList.Tail = &newNode
	} else {
		linkedList.Tail.Next = &newNode
		linkedList.Tail = &newNode
	}
}

func (linkedList *LinkedList) AddNodeToHead(v interface{}) {
	newNode := Node{
		Value: v,
		Next:  nil,
	}
	if linkedList.Head == nil {
		linkedList.Head = &newNode
		linkedList.Tail = &newNode
	} else {
		newNode.Next = linkedList.Head
		linkedList.Head = &newNode
	}
}

func (linkedList *LinkedList) AddNodeAtIndex(index int, v interface{}) error {
	if index == 0 {
		linkedList.AddNodeToHead(v)
		return nil
	}

	i := 0
	for currentNode := linkedList.Head; currentNode != nil; currentNode = currentNode.Next {
		if i == index-1 {
			newNode := Node{
				Value: v,
				Next:  currentNode.Next,
			}
			currentNode.Next = &newNode
			return nil
		}
		i++
	}

	return errors.New("index out of range")
}

func (linkedList *LinkedList) FindNodeByValue(v interface{}) []Node {
	if linkedList.Head == nil {
		return []Node{}
	}

	var foundValues []Node
	for currentNode := linkedList.Head; currentNode != nil; currentNode = currentNode.Next {
		if currentNode.Value == v {
			foundValues = append(foundValues, *currentNode)
		}
	}
	return foundValues
}

func (linkedList *LinkedList) RemoveNodeByIndex(index int) error {
	if linkedList.Head == nil {
		return errors.New("index out of range")
	}

	if index == 0 {
		tmpNode := linkedList.Head.Next
		linkedList.Head = nil
		linkedList.Head = tmpNode
		return nil
	}

	i := 0
	for currentNode := linkedList.Head; currentNode != nil; currentNode = currentNode.Next {
		if i == index-1 {
			if currentNode.Next != nil {
				tmp := currentNode.Next.Next
				currentNode.Next.Next = nil
				currentNode.Next = tmp
				return nil
			}
		}
		i++
	}

	return errors.New("index out of range")
}

func (linkedlist *LinkedList) FindNodeByIndex(index int) (Node, error) {
	if index < 0 {
		return Node{}, errors.New("index out of range")
	}

	i := 0
	for currentNode := linkedlist.Head; currentNode != nil; currentNode = currentNode.Next {
		if index == i {
			return *currentNode, nil
		}
		i++
	}
	return Node{}, errors.New("not found")
}

func (linkedList *LinkedList) Traverse() {
	if linkedList.Head == nil {
		return
	}

	for p := linkedList.Head; p != nil; p = p.Next {
		fmt.Print(" ", p.Value)
	}
	fmt.Println()
}
