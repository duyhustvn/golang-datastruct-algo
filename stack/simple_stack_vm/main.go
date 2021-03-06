package simplestackvm

import (
	"errors"
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
		var err error
		e := strings.Split(strings.Trim(element, " "), " ")
		if len(e) < 1 {
			return []int{}
		}
		operator := e[0]
		if len(e) > 1 {
			num, err = strconv.Atoi(e[1])
			if err != nil {
				return []int{}
			}
		}

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
			if err := add(&stack); err != nil {
				return []int{}
			}
			break
		case "SWAP":
			if err := swapTop2(&stack); err != nil {
				return []int{}
			}

			break
		default:
			return []int{}
		}
	}

	result := []int{}
	for !stack.IsEmpty() {
		e := stack.Pop().(int)
		result = append(result, e)
	}
	return result
}

func pop2Element(stack *Stack) (int, int, error) {
	num1 := stack.Pop()
	if num1 == nil {
		return 0, 0, errors.New("Require at least 2 element")
	}
	num2 := stack.Pop()
	if num2 == nil {
		return 0, 0, errors.New("Require at least 2 element")
	}
	return num1.(int), num2.(int), nil
}

func pushNumberToStack(stack *Stack, num int) {
	stack.Push(num)
}

func multiply(stack *Stack) error {
	num1, num2, err := pop2Element(stack)
	if err != nil {
		return err
	}
	stack.Push(num1 * num2)
	return nil
}

func add(stack *Stack) error {
	num1, num2, err := pop2Element(stack)
	if err != nil {
		return err
	}
	stack.Push(num1 + num2)
	return nil
}

func swapTop2(stack *Stack) error {
	num1, num2, err := pop2Element(stack)
	if err != nil {
		return err
	}
	stack.Push(num1)
	stack.Push(num2)
	return nil
}
