package simplestackvm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value interface{} `json:"value"`
	Next  *Node       `json:"next"`
}

type Stack struct {
	Head *Node `json:"head"`
}

func (s *Stack) IsEmpty() bool {
	if s.Head == nil {
		return true
	}
	return false
}

func createNewNode(v interface{}) Node {
	n := Node{
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

func Solution(a []string) []int {
	stack := Stack{}

	for _, element := range a {
		var num int
		e := strings.Split(strings.Trim(element, " "), " ")
		operator := e[0]
		if len(e) > 1 {
			num, _ = strconv.Atoi(e[1])
		}

		fmt.Printf("operator: %s num: %d\n", operator, num)

		switch operator {
		case "PUSH":
			pushNumberToStack(&stack, num)
			break
		case "MUL":
			if err := multiply(&stack); err != nil {
				return []int{}
			}
			break
		case "ADD":
			break
		case "SWAP":
			break
		default:
			break
		}
	}

	result := []int{}
	for !stack.IsEmpty() {
		e := stack.Pop().(int)
		result = append(result, e)
	}
	return result
}

func pushNumberToStack(stack *Stack, num int) {
	stack.Push(num)
}

func multiply(stack *Stack) error {
	num1 := stack.Pop()
	if num1 == nil {
		return errors.New("Require at least 2 element")
	}
	num2 := stack.Pop()
	if num2 == nil {
		return errors.New("Require at least 2 element")
	}

	stack.Push(num1.(int) * num2.(int))
	return nil
}
