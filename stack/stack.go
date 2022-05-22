package stack

import "fmt"

type SNode struct {
	Value interface{} `json:"value"`
	Next  *SNode      `json:"next"`
}

type Stack struct {
	Head *SNode `json:"head"`
}

func (s *Stack) IsEmpty() bool {
	if s.Head == nil {
		return true
	}
	return false
}

func createNewNode(v interface{}) SNode {
	n := SNode{
		Value: v,
	}
	return n
}

func (s *Stack) Push(v interface{}) {
	newNode := createNewNode(v)
	if s.IsEmpty() {
		s.Head = &newNode
		return
	}

	newNode.Next = s.Head
	s.Head = &newNode
	return
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.Head.Value
	s.Head = s.Head.Next
	return v
}

func (s *Stack) Traverse() {
	for node := s.Head; node != nil; node = node.Next {
		fmt.Printf("Node: %+v", node.Value)
	}
}
