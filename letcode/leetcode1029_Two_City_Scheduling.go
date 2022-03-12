package letcode

import "sort"

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	diff := make([][2]int, n)

	for i, cost := range costs {
		diff[i] = [2]int{i, cost[0] - cost[1]}
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i][1] < diff[j][1]
	})

	result := 0
	for i := 0; i < n; i++ {
		if i < (n+1)/2 {
			result += costs[diff[i][0]][0]
		} else {
			result += costs[diff[i][0]][1]
		}
	}

	return result
}
