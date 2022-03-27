package letcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}
	diff1 := make([]int,0)
	diff2 := make([]int,0)
	for num := range set1 {
		if !set2[num] {
			diff1 = append(diff1, num)
		}
	}
	for num := range set2 {
		if !set1[num] {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}

func minDeletion(nums []int) int {
	deletes := 0
	stack := make([]int,0)
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		} else {
			if stack[0] == nums[i] {
				deletes++
			} else {
				stack = make([]int,0)
			}
		}
	}
	return deletes + len(stack)
}

func kthPalindrome(queries []int, intLength int) []int64 {
	// 如何得到长度为 intLength 的所有回文数
	// 第一位分别是 1 到 9,最后一位分别是 1 到 9

	var allPalindrome func(intLength int) []int64
	allPalindrome = func(intLength int) []int64 {
		if intLength == 1 {
			return []int64{0,1,2,3,4,5,6,7,8,9}
		}
		if intLength == 2 {
			return []int64{00,11,22,33,44,55,66,77,88,99}
		}
		result := make([]int64, 0)
		pres := allPalindrome(intLength - 2)
		var prefix int64 = 1
		for i := 1; i < intLength; i++ {
			prefix *= 10
		}
		for i := 0; i < 10; i++ {
			for _, pre := range pres {
				result = append(result, prefix * int64(i) + pre*10 + int64(i))
			}
		}
		return result
	}

	all := make([]int64, 0)
	if intLength == 1 {
		all = []int64{1,2,3,4,5,6,7,8,9}
	}
	if intLength == 2 {
		all = []int64{11,22,33,44,55,66,77,88,99}
	}
	if intLength > 2 {
		all = allPalindrome(intLength)
		// 以 0 开头的
		invalid := len(allPalindrome(intLength - 2))
		all = all[invalid:]
	}
	result := make([]int64, len(queries))
	for i, query := range queries {
		if query > len(all) {
			result[i] = -1
		} else {
			result[i] = all[query-1]
		}
	}
	return result
}
