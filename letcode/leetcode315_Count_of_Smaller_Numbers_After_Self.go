package letcode

func countSmaller(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)

	// 保存 nums 在原数组中的 index，实际上这里可以用 map，但是数可能重复
	index := make([]int, n)
	for i := 0; i < n; i++ {
		// index 保存原始数据的 nums[i] 当前所在的位置。
		index[i] = i
	}

	var mergeSort func(left int, right int)
	mergeSort = func(left int, right int) {
		if left == right {
			return
		}

		mid := (left + right) / 2
		mergeSort(left, mid)
		mergeSort(mid + 1, right)

		i := left
		j := mid + 1

		tmp := make([]int, 0)
		tmpIndex := make([]int, 0)
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				tmp = append(tmp, nums[i])
				tmpIndex = append(tmpIndex, index[i])
				// 计算 nums[i] 右边的小弟
				counts[index[i]] += j - (mid + 1)
				i++
			} else {
				tmp = append(tmp, nums[j])
				tmpIndex = append(tmpIndex, index[j])
				j++
			}
		}

		for i <= mid {
			tmp = append(tmp, nums[i])
			tmpIndex = append(tmpIndex, index[i])
			// 计算右边的小弟
			counts[index[i]] += j - (mid + 1)
			i++
		}

		for j <= right {
			tmp = append(tmp, nums[j])
			tmpIndex = append(tmpIndex, index[j])
			j++
		}

		for k := left; k <= right; k++ {
			nums[k] = tmp[k - left]
			index[k] = tmpIndex[k - left]
		}
	}
	mergeSort(0 ,n - 1)
	return counts
}
