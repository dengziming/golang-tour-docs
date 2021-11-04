package letcode

import "math"

// 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := -1
	res := 0

	for i, v := range s {
		newStart, ok := m[v]
		if ok {
			start = int(math.Max(float64(start), float64(newStart)))
		}

		m[v] = i
		res = int(math.Max(float64(res), float64(i - start)))
	}

	return res
}
