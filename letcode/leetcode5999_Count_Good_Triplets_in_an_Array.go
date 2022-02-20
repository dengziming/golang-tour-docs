package letcode

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	index1 := make([]int, n)
	nums3 := make([]int, n)

	for i, num := range nums1 {
		index1[num] = i
	}
	// index1 可以把 nums1 变成 [0,1,2,..n-1]， 用它映射 nums2
	for i, num := range nums2 {
		nums3[i] = index1[num]
	}

	// 只需要计算 nums3 的三元组
	var count int64 = 0
	tree := make([]int, n+1)
	for i := 1; i < n-1; i++ {
		for j := nums3[i - 1] + 1; j <= n; j += j & -j { // 将 p[nums2[i-1]]+1 加入树状数组
			tree[j]++
		}
		y, less := nums3[i], 0
		for j := y; j > 0; j &= j - 1 { // 计算 less
			less += tree[j]
		}
		count += int64(less) * int64(n-1-y-(i-less))
	}
	return count
}

/**
// n^2 超时
func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	index1 := make([]int, n)
	nums3 := make([]int, n)

	for i, num := range nums1 {
		index1[num] = i
	}
	// index1 可以把 nums1 变成 [0,1,2,..n-1]， 用它映射 nums2
	for i, num := range nums2 {
		nums3[i] = index1[num]
	}

	// 只需要计算 nums3 的逆序数

	pre := make([]int64, n)
	post := make([]int64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums3[j] < nums3[i] {
				pre[i]++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums3[j] > nums3[i] {
				post[i]++
			}
		}
	}

	var count int64 = 0
	for i := 0; i < n; i++ {
		count += pre[i] * post[i]
	}

	return count
}
 */


/**
// n^3 超时
func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	index1, index2 := make([]int, n), make([]int, n)

	for i, num := range nums1 {
		index1[num] = i
	}
	for i, num := range nums2 {
		index2[num] = i
	}

	var count int64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (index1[j] - index1[i]) * (index2[j] - index2[i]) > 0 && (index1[k] - index1[i]) * (index2[k] - index2[i]) > 0 && (index1[k] - index1[j]) * (index2[k] - index2[j]) > 0 {
					count++
				}
			}
		}
	}
	return count
}
*/