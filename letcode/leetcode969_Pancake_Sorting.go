package letcode

func pancakeSort(arr []int) []int {
	n := len(arr)

	result := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		// 得到最大元素下标
		target := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[target] {
				target = j
			}
		}

		if target == i {
			continue
		}

		if target == 0 && i == 0 {
			continue
		}

		// 把 target 换到头
		if target != 0 {
			for j := 0; j < (target + 1) / 2; j++ {
				arr[j], arr[target - j] = arr[target - j], arr[j]
			}
		}

		// target 换到尾
		if i != 0 {
			for j := 0; j < (i + 1) / 2; j++ {
				arr[j], arr[i - j] = arr[i - j], arr[j]
			}
		}

		if target != 0 && i != 0 {
			result = append(result, target + 1, i + 1)
		} else if target != 0 {
			result = append(result, target + 1)
		} else if i != 0 {
			result = append(result, i + 1)
		}
	}
	return result
}
