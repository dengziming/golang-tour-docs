package letcode

import "math"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := -1
	res := 0

	for i, v := range s {
		start = int(math.Max(float64(start), float64(m[v] - 1)))
		m[v] = i + 1
		res = int(math.Max(float64(res), float64(i - start)))
	}

	return res
}
