package letcode

import (
	"fmt"
	"math"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	k := len(fmt.Sprint(n))
	total := atMostNGivenDigitSet2(digits, n, k)
	// 对于小于 k 的位数，每位各自有 m 个选择
	for i := 1; i < k; i++ {
		total += int(math.Pow(float64(len(digits)), float64(i)))
	}
	return total
}

func atMostNGivenDigitSet2(digits []string, n int, k int) int {
	total := 0

	top := n / int(math.Pow(float64(10), float64(k - 1)))
	sub := n % int(math.Pow(float64(10), float64(k - 1)))

	for _, digit := range digits {
		d, _ := strconv.Atoi(digit)
		if d < top {
			total += int(math.Pow(float64(len(digits)), float64(k - 1)))
		}
		if d == top {
			// 已经是一位数了，直接返回了
			if k == 1 {
				total += 1
			} else {
				total += atMostNGivenDigitSet2(digits, sub, k - 1)
			}
		}
	}
	return total
}
