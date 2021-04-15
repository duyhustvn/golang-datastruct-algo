package doublely

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
	Prev  *Node
}

func (linkedList *LinkedList) AddNode(v interface{}) {
	newNode := Node{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}

	if linkedList.Head == nil {
		linkedList.Head = &newNode
		linkedList.Tail = &newNode
	} else {
		newNode.Prev = linkedList.Tail
		linkedList.Tail.Next = &newNode
		linkedList.Tail = &newNode
	}
}

func (linkedList *LinkedList) AddNodeToHead(v interface{}) {
	newNode := Node{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if linkedList.Head == nil {
		linkedList.Head = &newNode
		linkedList.Tail = &newNode
	} else {
		newNode.Next = linkedList.Head
		linkedList.Head.Prev = &newNode
		linkedList.Head = &newNode
	}
}

func (linkedList *LinkedList) AddNodeAtIndex(index int, v interface{}) error {
	// add to head
	if index == 0 {
		linkedList.AddNodeToHead(v)
		return nil
	}

	i := 0
	for currentNode := linkedList.Head; currentNode != nil; currentNode = currentNode.Next {
		if i == index-1 {
			// add to tail
			if currentNode.Next == nil {
				linkedList.AddNode(v)
			} else {
				newNode := Node{
					Value: v,
					Next:  currentNode.Next,
					Prev:  currentNode,
				}
				currentNode.Next.Prev = &newNode
				currentNode.Next = &newNode
			}
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
		return errors.New("empty")
	}

	if index == 0 {
		// if head = tail
		if linkedList.Head == linkedList.Tail {
			linkedList.Head = nil
			linkedList.Tail = nil
			return nil
		}

		tmpNode := linkedList.Head.Next
		linkedList.Head.Next.Prev = nil
		linkedList.Head = nil
		linkedList.Head = tmpNode
		return nil
	}

	i := 0
	for currentNode := linkedList.Head; currentNode != nil; currentNode = currentNode.Next {
		if i == index-1 {
			if currentNode == linkedList.Tail {
				return errors.New("index out of range")
			} else if currentNode.Next == linkedList.Tail {
				linkedList.Tail.Prev.Next = nil
				linkedList.Tail = currentNode
				return nil
			} else {
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
	return Node{}, errors.New("index out of range")
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
