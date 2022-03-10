package letcode

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	//
	check := func(mid, k int) bool {
		i, j := n - 1, 0
		num := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				num += i + 1
				j++
			} else {
				i--
			}
		}
		return num >= k
	}

	left := matrix[0][0]
	right := matrix[n-1][n-1]

	for left < right {
		mid := left + (right - left) / 2
		if check(mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
