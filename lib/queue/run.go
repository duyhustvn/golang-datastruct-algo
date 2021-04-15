package queue

import (
	"fmt"
)

func RunQueue() {
	var queue Queue

	queue.Enqueue(2)
	queue.Enqueue(4)
	queue.Enqueue(10)
	queue.Traverse()
	var result interface{}
	var err error
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n ", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("dequeue: %+v\n", result)
}
