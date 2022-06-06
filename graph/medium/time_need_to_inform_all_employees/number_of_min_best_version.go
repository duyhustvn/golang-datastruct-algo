package time_need_to_inform_all_employees

func numOfMinutesOptimal(n int, headID int, manager []int, informTime []int) int {
	if n == 0 || n == 1 {
		return 0
	}
	adjacencyList := buildAdjacencyList(headID, manager)
	return dfs(adjacencyList, headID, informTime)
}

func dfs(adjacencyList map[int][]int, currentID int, informTime []int) int {
	subordinates := adjacencyList[currentID]
	max := 0
	for _, subordinate := range subordinates {
		if currentTime := dfs(adjacencyList, subordinate, informTime); currentTime > max {
			max = currentTime
		}
	}
	return max + informTime[currentID]
}
