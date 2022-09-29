package letcode

import "sort"

// 思路:
// 以 x 为中心，向两边扩散
// 一开始各扩散 k/2，得到区间 [a1, b1], 比较两个点谁更近一点，比如 b1 更近，那 [x, b1] 属于结果区间的一部分
// 然后各自扩散 k/4， 得到区间 [x-k/4, b1+k/4]，比较两个点谁更近
// 这种解法复杂度是 logk ，比官方的题解更优
func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	if k == 0 {
		return []int{}
	}
	resultLeft := sort.SearchInts(arr, x)
	resultRight := resultLeft - 1

	for resultLength := 0; resultLength != k && resultLeft != 0 && resultRight != n-1; {
		stepLength := k - resultLength

		left := resultLeft - (stepLength+1)/2
		right := resultRight + (stepLength+1)/2

		// 调整区间
		if right > n-1 {
			right = n - 1
			left = right - k + 1
		}
		if left < 0 {
			left = 0
			right = left + k - 1
		}

		// 比较哪个更近

		// right 更近
		if arr[right]-x < x-arr[left] {
			resultRight = right
		} else { // left 更近
			resultLeft = left
		}
		resultLength = resultRight - resultLeft + 1
	}
	if resultLeft == 0 {
		return arr[:k]
	} else if resultRight == n-1 {
		return arr[n-k:]
	}

	return arr[resultLeft : resultRight+1]
}
