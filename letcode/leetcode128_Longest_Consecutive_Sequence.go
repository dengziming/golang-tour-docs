package letcode

import "math"
func longestConsecutive(nums []int) int {
	headquarter := make(map[int]int)
	for _, num := range nums {
		headquarter[num] = num
	}

	var union func(x int, y int)
	var find func(x int) int

	find = func(x int) int {
		if _, ok := headquarter[x]; ok {
			if headquarter[x] != x {
				headquarter[x] = find(headquarter[x])
			}
			return headquarter[x]
		}
		return math.MinInt
	}

	union = func(x int, y int) {
		headquarter[find(x)] = find(y)
	}

	for  _, num := range nums {
		if find(num + 1) != math.MinInt {
			union(num, num + 1)
		}
	}

	max := 0
	for _, num := range nums {
		head := find(num)
		if head - num + 1 > max {
			max = head - num + 1
		}
	}

	return max
}
