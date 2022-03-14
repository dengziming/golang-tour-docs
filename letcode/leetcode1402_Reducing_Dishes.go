package letcode

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	n := len(satisfaction)

	// 二分法找到 satisfaction 中第一个大于等于0 的数
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if satisfaction[mid] >= 0 {
			right = mid
		} else {
			left = mid+1
		}
	}
	// 这代表二分没找到结果，说明所有的都小于 0
	if satisfaction[left] < 0 {
		return 0
	}
	// left 代表第一个大于等于 0 的菜
	// 先计算出抛弃所有小于 0 的菜能得到的价值，以及每次加一个能得到的价值
	result := 0
	diff := 0
	for i := left; i < n; i++ {
		result += (i-left+1)*satisfaction[i]
		diff += satisfaction[i]
	}

	for i := left-1; i >= 0; i-- {
		// 加入 i - 1 能带来的增量和减量
		increase := diff
		decrease := -satisfaction[i]
		if increase >= decrease {
			result += increase-decrease
			diff -= decrease
		} else {
			break
		}
	}

	return result
}
