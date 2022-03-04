package letcode

func findLength(nums1 []int, nums2 []int) int {

	max := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		} else {
			return v2
		}
	}

	commonDp := func(path1, path2 []int) int {
		maxV := 0
		m, n := len(path1), len(path2)
		dp := make([][]int, m)
		for i := 0; i < m; i++ {
			dp[i] = make([]int, n)
			for j := 0; j < n; j++ {
				if path1[i] == path2[j] {
					if i == 0 || j == 0 {
						dp[i][j] = 1
					} else {
						dp[i][j] = dp[i-1][j-1]+1
					}
					maxV = max(maxV, dp[i][j])
				}
			}
		}
		return maxV
	}

	return commonDp(nums1, nums2)
}
