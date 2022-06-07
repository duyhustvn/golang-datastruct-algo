package time_need_to_inform_all_employees

func NumOfMinutesOptimal(n int, headID int, manager []int, informTime []int) int {
	if n == 0 || n == 1 {
		return 0
	}
	adjacencyList := buildAdjacencyList(headID, manager)
	return dfs(adjacencyList, headID, informTime)
}

func dfs(adjacencyList map[int][]int, currentID int, informTime []int) int {
	if len(adjacencyList[currentID]) == 0 {
		return 0
	}

	subordinates := adjacencyList[currentID]
	max := 0
	for _, subordinate := range subordinates {
		currentTime := dfs(adjacencyList, subordinate, informTime)
		if currentTime > max {
			max = currentTime
		}
	}
	return max + informTime[currentID]
}
