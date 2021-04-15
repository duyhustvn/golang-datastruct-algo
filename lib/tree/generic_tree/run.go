package generic_tree

import (
	queue "ds/lib/queue"
	"fmt"
)

type Node struct {
	tag      string
	text     string
	class    string
	id       string
	src      string
	alt      string
	children []*Node
}

func (node *Node) PrintNodeInfo() {
	fmt.Printf("id: %s\ntag: %s\n", node.id, node.tag)
}

func DFS(node *Node) {
	fmt.Println(node.tag)
	for _, child := range node.children {
		DFS(child)
	}
}

func BFS(node *Node) {
	var q queue.Queue
	q.Enqueue(node)
	for !q.IsEmpty() {
		tmpNode, _ := q.Dequeue()
		fmt.Println(tmpNode.(*Node).tag)
		for _, child := range tmpNode.(*Node).children {
			q.Enqueue(child)
		}
	}
}

func FindByIdBFS(node *Node, id string) *Node {
	var q queue.Queue
	q.Enqueue(node)
	for !q.IsEmpty() {
		tmpNode, _ := q.Dequeue()
		if tmpNode.(*Node).id == id {
			return tmpNode.(*Node)
		}
		for _, child := range tmpNode.(*Node).children {
			q.Enqueue(child)
		}
	}
	return nil
}

func FindByIdDFS(node *Node, id string) *Node {
	if node.id == id {
		return node
	}

	for _, child := range node.children {
		return FindByIdDFS(child, id)
	}
	return nil
}

func RunGenericTree() {
	image := Node{
		tag: "img",
		src: "http://example.com/logo.svg",
		alt: "Example's log",
	}

	p := Node{
		tag:      "p",
		text:     "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{&image},
	}

	h1 := Node{
		tag:  "h1",
		text: "This is a H1",
	}

	span := Node{
		tag:  "span",
		text: "2019 &copy; Ilija Eftimov",
		id:   "copyright",
	}

	div := Node{
		tag:      "div",
		class:    "footer",
		text:     "This is the footer of the page",
		children: []*Node{&span},
	}

	body := Node{
		tag:      "body",
		children: []*Node{&div, &h1, &p},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}

	// DFS
	// DFS(&html)

	// FIND BY ID USING DFS
	// fmt.Println()
	// id := "copyright"
	// found := FindByIdDFS(&html, id)
	// if found != nil {
	// 	fmt.Println("Found node")
	// 	found.PrintNodeInfo()
	// } else {
	// 	fmt.Printf("Not found node with id %s \n", id)
	// }

	// BFS
	BFS(&html)

	// FIND BY ID USING DFS
	fmt.Println()
	id := "copyright"
	found := FindByIdBFS(&html, id)
	if found != nil {
		fmt.Println("Found node")
		found.PrintNodeInfo()
	} else {
		fmt.Printf("Not found node with id %s \n", id)
	}
}
