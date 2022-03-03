package letcode

func minOperations2(boxes string) []int {
	n := len(boxes)

	// 左边的都移到 dp[i] 需要多少次操作
	left := make([]int, n)
	// 左边有多少个 1
	cnt := make([]int, n)

	for i := 1; i < len(boxes); i++ {
		cnt[i] = cnt[i-1]
		if boxes[i-1] == '1' {
			cnt[i]++
		}
		left[i] = left[i-1] + cnt[i]
	}

	tatal := cnt[n-1]
	if boxes[n-1] == '1' {
		tatal++
	}
	// 右边的都移到 dp[i] 需要多少次操作
	right := make([]int, n)
	cnt = make([]int, n)
	for i := n-2; i >= 0; i-- {
		cnt[i] = cnt[i+1]
		if boxes[i+1] == '1' {
			cnt[i]++
		}
		right[i] = right[i+1] + cnt[i]
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = left[i] + right[i]
	}
	return result
}
