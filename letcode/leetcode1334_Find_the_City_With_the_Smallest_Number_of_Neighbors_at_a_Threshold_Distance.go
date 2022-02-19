package letcode

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	min := func(a int, b int) int {
		if a >= b {
			return b
		}
		return a
	}

	// 重新计算邻接矩阵
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			if i != j {
				distance[i][j] = math.MaxInt
			}
		}
	}
	for _, edge := range edges {
		distance[edge[0]][edge[1]] = edge[2]
		distance[edge[1]][edge[0]] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distance[i][k] != math.MaxInt && distance[k][j] != math.MaxInt {
					distance[i][j] = min(distance[i][j], distance[i][k] + distance[k][j])
				}
			}
		}
	}

	minIdx := 0
	minCnt := math.MaxInt
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i !=j && distance[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= minCnt {
			minCnt = cnt
			minIdx = i
		}
	}

	return minIdx
}
