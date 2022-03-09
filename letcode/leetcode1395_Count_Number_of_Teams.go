package letcode

func numTeams(rating []int) int {
	n := len(rating)
	// 左右比自己小的数量
	left := make([]int,n)
	right := make([]int, n)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				left[i]++
			}
		}
	}
	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			if rating[j] < rating[i] {
				right[i]++
			}
		}
	}

	result := 0

	for i := 0; i < n; i++ {
		result += left[i]*(n-i-1-right[i])
		result += (i-left[i])*right[i]
	}
	return result
}
