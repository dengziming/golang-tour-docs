package letcode

func maximumRequests(n int, requests [][]int) int {

	// 代表每一个楼的流量
	flow := make([]int, n)

	var dfs func(i int, cnt int)

	output := func() bool {
		for _, f := range flow {
			if f != 0 {
				return false
			}
		}
		return true
	}

	mavV := 0
	dfs = func(i int, cnt int) {
		if cnt > mavV && output() {
			mavV = cnt
		}
		if i == len(requests) {
			return
		}

		// 不选它
		dfs(i+1, cnt)

		// 选它
		flow[requests[i][0]]--
		flow[requests[i][1]]++
		dfs(i+1, cnt+1)
		flow[requests[i][0]]++
		flow[requests[i][1]]--
	}
	dfs(0, 0)
	return mavV
}
