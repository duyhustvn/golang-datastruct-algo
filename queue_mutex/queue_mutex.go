package lib

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	mutex sync.Mutex
	Front *Node
	Rear  *Node
}

func (q *Queue) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Front == nil && q.Rear == nil {
		return true
	}
	return false
}

func (q *Queue) Enqueue(v interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	newNode := Node{
		Value: v,
		Next:  nil,
	}
	if q.IsEmpty() {
		q.Front = &newNode
		q.Rear = &newNode
	} else {
		q.Rear.Next = &newNode
		q.Rear = &newNode
	}
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	tmp := q.Front
	result := tmp.Value
	q.Front = q.Front.Next

	if q.Front == nil {
		q.Front = nil
		q.Rear = nil
	}

	tmp = nil // free tmp
	return result, nil
}

func (q *Queue) Traverse() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for currentNode := q.Front; currentNode != nil; currentNode = currentNode.Next {
		fmt.Printf("%+v ", currentNode.Value)
	}
	fmt.Println("")
}
