package stack

import "testing"

func TestPushInt(t *testing.T) {
	stack := Stack{}

	if stack.Head != nil {
		t.Error("Wrong value")
	}

	stack.Push(1)
	stack.Push(2)

	if stack.Head.Value.(int) != 2 {
		t.Error("Wrong value")
	}

	if stack.Head.Next.Value.(int) != 1 {
		t.Error("Wrong value")
	}
}

func TestPushString(t *testing.T) {
	stack := Stack{}

	if stack.Head != nil {
		t.Error("Wrong value")
	}

	stack.Push("one")

	if stack.Head.Value.(string) != "one" {
		t.Error("Wrong value")
	}

	stack.Push("two")

	if stack.Head.Value.(string) != "two" {
		t.Error("Wrong value")
	}

	if stack.Head.Next.Value.(string) != "one" {
		t.Error("Wrong value")
	}
}

func TestPushStruct(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	stack := Stack{}

	if stack.Head != nil {
		t.Error("Wrong value")
	}

	person1 := Person{
		Name: "duy",
		Age:  28,
	}

	stack.Push(person1)
	if stack.Head.Value.(Person).Name != "duy" || stack.Head.Value.(Person).Age != 28 {
		t.Error("Wrong value")
	}

	person2 := Person{
		Name: "hien",
		Age:  27,
	}
	stack.Push(person2)

	if stack.Head.Value.(Person).Name != "hien" || stack.Head.Value.(Person).Age != 27 {
		t.Error("Wrong value")
	}

	if stack.Head.Next.Value.(Person).Name != "duy" || stack.Head.Next.Value.(Person).Age != 28 {
		t.Error("Wrong value")
	}
}

func TestPopInt(t *testing.T) {
	stack := Stack{}
	if v := stack.Pop(); v != nil {
		t.Error("Wrong value")
	}
	stack.Push(1)
	if v := stack.Pop(); v.(int) != 1 {
		t.Error("Wrong value")
	}

	stack.Push(1)
	stack.Push(2)
	if v := stack.Pop(); v.(int) != 2 {
		t.Error("Wrong value")
	}
	if v := stack.Pop(); v.(int) != 1 {
		t.Error("Wrong value")
	}
}

func TestPopString(t *testing.T) {
	stack := Stack{}
	if v := stack.Pop(); v != nil {
		t.Error("Wrong value")
	}
	stack.Push("one")
	if v := stack.Pop(); v.(string) != "one" {
		t.Error("Wrong value")
	}

	stack.Push("one")
	stack.Push("two")
	if v := stack.Pop(); v.(string) != "two" {
		t.Error("Wrong value")
	}
	if v := stack.Pop(); v.(string) != "one" {
		t.Error("Wrong value")
	}
}

func TestPopStruct(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	stack := Stack{}

	if v := stack.Pop(); v != nil {
		t.Error("Wrong value")
	}

	person1 := Person{
		Name: "duy",
		Age:  28,
	}
	stack.Push(person1)
	if v := stack.Pop(); v.(Person).Name != "duy" || v.(Person).Age != 28 {
		t.Error("Wrong value")
	}

	person2 := Person{
		Name: "duy",
		Age:  28,
	}
	person3 := Person{
		Name: "hien",
		Age:  27,
	}

	stack.Push(person2)
	stack.Push(person3)
	if v := stack.Pop(); v.(Person).Name != "hien" || v.(Person).Age != 27 {
		t.Error("Wrong value")
	}

	if v := stack.Pop(); v.(Person).Name != "duy" || v.(Person).Age != 28 {
		t.Error("Wrong value")
	}
}
