package letcode

import "sort"

// https://leetcode-cn.com/problems/reveal-cards-in-increasing-order/
func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		queue = append(queue, i)
	}
	ans := make([]int, n)

	for i := range deck {
		// 知道了下一个被取的数的位置
		ans[queue[0]] = deck[i]
		queue = queue[1:]
		if len(queue) != 0 {
			queue = append(queue[1:], queue[0])
		}
	}

	return ans
}
