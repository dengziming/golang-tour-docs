package letcode

import "fmt"

func maximumGood(statements [][]int) int {
	n := len(statements)

	markedTrue := make([]bool, n)
	markedFalse := make([]bool, n)
	// 保存所有好人说过的话, good 标识有多少人说他是好人，bad 标识有多少人说他是坏人
	good := make(map[int]int)
	bad := make(map[int]int)

	// 基于回溯
	var dfs func(i int, cnt int)

	maxV := 0
	tmp := make([]bool, n)
	dfs = func(i int, cnt int) {
		if i == n {
			if cnt > maxV {
				maxV = cnt
				tmp = append([]bool{}, markedTrue...)
			}
			return
		}

		// i 是坏人，需要满足：
		// 1. good 里面没有说他是好人
		if good[i] == 0 {
			markedFalse[i] = true
			dfs(i+1, cnt)
			markedFalse[i] = false
		}

		if bad[i] == 0 {
			// i 是好人，需要满足:
			// 1. bad 里面没有 i
			// 2. good 里面没有 i 说的 bad
			// 3. bad 里面没有 i 说的 good
			flag := true
			for j := 0; j < n; j++ {
				if statements[i][j] == 0 && (good[j] > 0 || markedTrue[j]) || statements[i][j] == 1 && (bad[j] > 0 || markedFalse[j]) {
					flag = false
					break
				}
			}
			if flag {
				for j := 0; j < n; j++ {
					if statements[i][j] == 0 {
						bad[j]++
					}
					if statements[i][j] == 1 {
						good[j]++
					}
				}
				markedTrue[i] = true
				dfs(i+1, cnt+1)
				markedTrue[i] = false
				for j := 0; j < n; j++ {
					if statements[i][j] == 0 {
						bad[j]--
					}
					if statements[i][j] == 1 {
						good[j]--
					}
				}
			}
		}
	}
	dfs(0, 0)
	fmt.Printf("%d %v \n", maxV, tmp)
	return maxV
}
