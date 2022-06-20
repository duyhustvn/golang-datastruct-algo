package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList, indegreeList := buildAdjacencyList(prerequisites)
	zeroIndegreeList := getZeroIndegreeNode(numCourses, indegreeList)
	indegreeList = topologicalSort(adjacencyList, indegreeList, zeroIndegreeList)
	return len(indegreeList) == 0
	// return !isCyclist(adjacencyList)
}

func buildAdjacencyList(prerequisites [][]int) (map[int][]int, map[int]int) {
	adjacencyList := make(map[int][]int)
	indegreeList := make(map[int]int)

	for _, prerequisite := range prerequisites {
		if _, ok := adjacencyList[prerequisite[1]]; !ok {
			adjacencyList[prerequisite[1]] = []int{prerequisite[0]}
		} else {
			adjacencyList[prerequisite[1]] = append(adjacencyList[prerequisite[1]], prerequisite[0])
		}

		if _, ok := indegreeList[prerequisite[0]]; !ok {
			indegreeList[prerequisite[0]] = 1
		} else {
			indegreeList[prerequisite[0]]++
		}
	}

	return adjacencyList, indegreeList
}

func isCyclist(adjacencyList map[int][]int) bool {
	for node, _ := range adjacencyList {
		if bfs(node, adjacencyList) {
			return true
		}
	}
	return false
}

type Queue struct {
	v []int
}

func newQueue() *Queue {
	return &Queue{v: make([]int, 0)}
}

func (q *Queue) Push(v int) {
	q.v = append(q.v, v)
}

func (q *Queue) IsEmpty() bool {
	return len(q.v) == 0
}

func (q *Queue) Pop() (int, bool) {
	if !q.IsEmpty() {
		queueLength := len(q.v)
		popValue := q.v[0]
		q.v = q.v[1:queueLength]
		return popValue, true
	}
	return 0, false
}

func bfs(node int, adjacencyList map[int][]int) bool {
	q := newQueue()
	q.Push(node)

	count := 0
	for !q.IsEmpty() {
		curNode, _ := q.Pop()
		if count != 0 {
			if curNode == node {
				return true
			}
		}
		count++

		neighbors := adjacencyList[curNode]
		for _, neighbor := range neighbors {
			q.Push(neighbor)
		}
	}
	return false
}

/*
   Get missising key in range from 0 -> n-1 in indegreeList
  **/
func getZeroIndegreeNode(n int, indegreeList map[int]int) []int {
	zeroIndegree := []int{}
	for i := 0; i < n; i++ {
		if _, ok := indegreeList[i]; !ok {
			zeroIndegree = append(zeroIndegree, i)
		}
	}
	return zeroIndegree
}

func topologicalSort(adjacencyList map[int][]int, indegreeList map[int]int, zeroIndegreeList []int) map[int]int {
	for len(zeroIndegreeList) > 0 {
		curNode := zeroIndegreeList[0]
		zeroIndegreeList = zeroIndegreeList[1:]

		if pointedNodes, ok := adjacencyList[curNode]; ok {
			for _, node := range pointedNodes {
				indegreeLevel := indegreeList[node]
				if indegreeLevel-1 == 0 {
					zeroIndegreeList = append(zeroIndegreeList, node)
					delete(indegreeList, node)
				} else {
					indegreeList[node]--
				}
			}
		}
	}
	return indegreeList
}
