package letcode

func findKDistantIndices(nums []int, key int, k int) []int {
	idxs := make(map[int]bool)

	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, num := range nums {
		if num == key {
			for j := max(i-k, 0); j <= min(i+k, len(nums)); j++ {
				idxs[j] = true
			}
		}
	}

	result := make([]int, 0)
	for i := range nums {
		if idxs[i] {
			result = append(result, i)
		}
	}
	return result
}
