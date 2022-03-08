package letcode

import "sort"

func maxRunTime(n int, batteries []int) int64 {

	var digui func(n int, batteries []int) int64

	sort.Ints(batteries)
	sum := int64(0)
	for i := range batteries {
		sum+=int64(batteries[i])
	}

	digui = func(n int, batteries []int) int64 {
		length := len(batteries)
		if n == 1 {
			return sum
		}

		ceil := sum/int64(n)

		if batteries[length - 1] > int(ceil) {
			sum -= int64(batteries[length-1])
			return digui(n-1, batteries[:length-1])
		} else {
			return ceil
		}
	}

	return digui(n, batteries)
}
