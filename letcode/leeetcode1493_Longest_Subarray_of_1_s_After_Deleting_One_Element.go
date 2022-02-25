package letcode

func longestSubarray(nums []int) int {
	// 除非元素都是1，删除的一定不是 1，记录下不是 1 的索引
	indies := make([]int, 0)
	for i, num := range nums {
		if num != 1 {
			indies = append(indies, i)
		}
	}
	if len(indies) == 0 || len(indies) == 1 {
		return len(nums) - 1
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 给两边补一个 -1 和 n
	indies = append([]int{-1}, indies...)
	indies = append(indies, len(nums))
	maxV := 0
	for i := 0; i < len(indies) -1 ; i++ {
		maxV = max(maxV, indies[i+1] - indies[i] - 1)
	}
	// 遍历所有的下标，如果某一个下标，它的下一个下标和它挨着，那删除它没有意义。
	for i := 1; i < len(indies)-1; i++ {
		if indies[i-1]+1 == indies[i] || indies[i+1] +1 == indies[i] {
			continue
		} else {
			tmp := indies[i+1] - indies[i-1] - 1 - 1
			if tmp > maxV {
				maxV = tmp
			}
		}
	}
	return maxV
}
