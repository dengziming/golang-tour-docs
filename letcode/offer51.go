package letcode

func reversePairs(nums []int) int {
	var mergeSort func(left int, right int) int

	mergeSort = func(left int, right int) int {
		if left >= right {
			return 0
		}
		mid := (left + right) / 2
		// 经过这里能保证两段都已经排好序了
		cnt := mergeSort(left, mid) + mergeSort(mid + 1, right)

		i := left
		j := mid + 1
		var merge []int
		for i <= mid && j <= right  {
			if nums[i] <= nums[j] {
				// 计算此时 i 前面有多少个来自后半段的元素
				cnt += j - (mid + 1)
				merge = append(merge, nums[i])
				i ++
			} else {
				// 没有贡献直接交换
				merge = append(merge, nums[j])
				j++
			}
		}
		for i <= mid {
			cnt += j - (mid + 1)
			merge = append(merge, nums[i])
			i ++
		}
		for j <= right {
			merge = append(merge, nums[j])
			j++
		}
		for k := left; k <= right; k++ {
			nums[k] = merge[k - left]
		}
		return cnt
	}

	return mergeSort(0, len(nums) - 1)
}
