package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	return true
}

func buildAdjacencyList(prerequisites [][]int) map[int][]int {
	adjacencyList := make(map[int][]int)

	for _, prerequisite := range prerequisites {
		if _, ok := adjacencyList[prerequisite[1]]; !ok {
			adjacencyList[prerequisite[1]] = []int{prerequisite[0]}
		} else {

			adjacencyList[prerequisite[1]] = append(adjacencyList[prerequisite[1]], prerequisite[0])
		}
	}

	return adjacencyList
}
