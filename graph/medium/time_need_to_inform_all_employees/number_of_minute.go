package time_need_to_inform_all_employees

import (
	"ds/graph/lib"
	"ds/stack"
)

// https://leetcode.com/problems/time-needed-to-inform-all-employees/

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	numberOfMinute := 0

	adjacencyList := buildAdjacencyList(headID, manager)
	grap, headNode := buildGraphFromAdjacencyList(adjacencyList, headID)
	getDepth(grap, headNode, informTime)
	return numberOfMinute
}

func buildAdjacencyList(headID int, manager []int) map[int][]int {
	mapDirectManagerToSub := make(map[int][]int) // map direct manager to subordination

	for suborId, man := range manager {
		if man == -1 {
			continue
		}
		subors, ok := mapDirectManagerToSub[man]
		if !ok {
			mapDirectManagerToSub[man] = []int{suborId}
			continue
		}
		mapDirectManagerToSub[man] = append(subors, suborId)
	}

	return mapDirectManagerToSub
}

func buildGraphFromAdjacencyList(adjacencyList map[int][]int, headID int) (*lib.ItemGraph, *lib.Node) {
	graph := lib.NewGraph()
	var headNode *lib.Node
	existedNode := make(map[int]*lib.Node)
	for man, subors := range adjacencyList {
		manNode, ok := existedNode[man]
		if !ok {
			manNode = lib.NewNode(man)
		}
		if man == headID {
			headNode = manNode
		}
		graph.AddNode(manNode)
		for _, sub := range subors {
			subNode := lib.NewNode(sub)
			graph.AddNode(subNode)
			graph.AddEdge(manNode, subNode)
		}
	}
	return graph, headNode
}

func getDepth(graph *lib.ItemGraph, startNode *lib.Node, informTime []int) int {
	depth := make(map[*lib.Node]int)
	maxDepth := 0
	visited := make(map[lib.Node]bool)
	s := stack.Stack{}
	s.Push(startNode)
	visited[*startNode] = true
	depth[startNode] = 0
	for !s.IsEmpty() {
		node := s.Pop().(*lib.Node)
		vNode := (*node).Value.(int)
		neighbors := graph.Edges[*node]
		for _, neighbor := range neighbors {
			if _, ok := visited[*neighbor]; !ok {
				preDepth := depth[node]
				if preDepth+informTime[vNode] > maxDepth {
					maxDepth = preDepth + informTime[vNode]
				}
				depth[neighbor] = preDepth + informTime[vNode]
				s.Push(neighbor)
				visited[*neighbor] = true
			}
		}
	}
	return maxDepth
}
