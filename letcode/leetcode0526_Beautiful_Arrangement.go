package letcode

func countArrangement(n int) int {
	// 递归, 第 n 个数字，可以是第 1 到 n，所以需要累计 1 到 n 的情况，而对于剩下的 n - 1 个数，转化为一个更小的问题
	// f(n) = ∑f(n-k)

	// 回溯或者直接 dfs 等方法写烂了，今天不写了
	// 动态规划，空间优化版本
	mask := 1 << n
	dp := make([]int, mask)

	getCnt := func(x int) int {
		ans := 0
		for x != 0 {
			x -= (x & -x)
			ans ++
		}
		return ans
	}

	dp[0] = 1
	for state := 1; state < mask; state++ {
		cnt := getCnt(state)
		for i := 0; i < n; i++ {
			if (((state >> i) & 1) == 0) {
				continue
			}
			if ((i + 1) % cnt != 0 && cnt % (i + 1) != 0) {
				continue
			}
			dp[state] += dp[state & (^(1 << i))]
		}
	}

	return dp[mask-1]
}

func countArrangement_dp(n int) int {
	// 递归, 第 n 个数字，可以是第 1 到 n，所以需要累计 1 到 n 的情况，而对于剩下的 n - 1 个数，转化为一个更小的问题
	// f(n) = ∑f(n-k)

	// 回溯或者直接 dfs 等方法写烂了，今天不写了
	// 动态规划，dp[i][j] 表示前面 i 个数字，选择的方案为 j (二进制拆分)
	dp := make([][]int, n+1)
	mask := 1 << n
	dp[0] = make([]int, mask)
	dp[0][0] = 1

	for i := 1; i < n+1; i++ {
		dp[i] = make([]int, mask)
		for state := 0; state < mask; state++ {
			for k := 1; k < n+1; k++ {
				// k 在 state 中必须是 1
				if ((state>>(k-1)))&1 == 0 {
					continue
				}
				if (k % i != 0 && i % k != 0) {
					continue
				}
				dp[i][state] += dp[i-1][state & (^(1 << (k - 1)))]
			}
		}
	}
	return dp[n][mask-1]
}

func countArrangement_dfs(n int) int {
	if n == 1 {
		return 1
	}
	result := 0
	used := make(map[int]bool)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n+1 {
			result++
			return
		}
		for j := 1; j <= n; j++ {
			if used[j] == false && (j % i == 0 || i % j ==0) {
				used[j] = true
				dfs(i+1)
				used[j] = false
			}
		}
	}
	// 动态规划
	dfs(1)
	return result
}
