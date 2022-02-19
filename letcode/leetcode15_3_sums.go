package letcode

func threeSum(candidates []int) [][]int {
	result := make([][]int, 0)

	// 计数
	counts := make(map[int]int)
	sets := make(map[int]bool)
	for _, candidate:= range candidates {
		counts[candidate] = counts[candidate] + 1
		sets[candidate] = true
	}

	for first := range sets {
		counts[first]--

		for second := range sets {
			if second>= first && counts[second] > 0 {
				counts[second]--
				third := -(first + second)
				if third >= second && counts[third] > 0 {
					result = append(result, []int{first, second, third})
				}
				counts[second]++
			}
		}
		counts[first]++
	}

	return result
}
