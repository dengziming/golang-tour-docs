package letcode

func checkArithmeticSubarrays(nums []int, l []int, r []int) (result []bool) {
	check := func(left, right int) bool {
		if right-left < 2 {
			return true
		}
		target := nums[left: right+1]
		min := target[0]
		max := target[0]
		for i := range target {
			if target[i] < min {
				min = target[i]
			}
			if target[i] > max {
				max = target[i]
			}
		}
		derivative := float64(max - min)/float64(right-left)
		if derivative == 0.0 {
			return true
		}

		counts := make(map[float64]int)
		for _, num := range target {
			counts[float64(num-min)/derivative]++
		}
		for i := range target {
			if counts[float64(i)] != 1 {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(l); i++ {
		result = append(result, check(l[i], r[i]))
	}
	return
}
