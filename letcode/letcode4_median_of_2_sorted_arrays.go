package letcode

// 整体思路 1
// 先用两个指针分别指向两个数组中间的元素，比较大小
//   两个相等就输出
//   如果不等，不妨设 A 的更大，那就往回移动 A（同时往前移动 B）
// 这个边界处理很麻烦
//
// 整体思路 2
// 每次删掉一半
//

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tuple := TwoSortedArray{nums1: nums1, nums2: nums2}

	if (len(nums1) + len(nums2)) % 2 == 1 {
		return float64(tuple.findK(0, 0, (len(nums1)+len(nums2))/2+1))
	} else {
		return (float64(tuple.findK(0, 0, (len(nums1) + len(nums2)) / 2)) + float64(tuple.findK(0, 0, (len(nums1) + len(nums2)) / 2 + 1))) / 2
	}

}

type TwoSortedArray struct {
	nums1 []int // 数组1
	nums2 []int // 数组2
}

// 找到两个数组中第k小的元素
// 一次删除 k/2 个元素，直到 k == 1
// 如果
func (tuple *TwoSortedArray) findK(start1 int, start2 int, k int) int {
	// 递归机
	if start1 == len(tuple.nums1) {
		return tuple.nums2[start2 + k - 1]
	}
	if start2 == len(tuple.nums2) {
		return tuple.nums1[start1 + k - 1]
	}
	if k == 1 {
		return min(tuple.nums1[start1], tuple.nums2[start2])
	}

	// 下一次递归开始的位置，如果为 len(tuple)
	newStart1 := min(start1 + k / 2, len(tuple.nums1))
	newStart2 := min(start2 + k / 2, len(tuple.nums2))

	// 两种分而治之的情况，一次性删掉 k / 2 个数据
	if tuple.nums1[newStart1 - 1] < tuple.nums2[newStart2 - 1] {
		return tuple.findK(newStart1, start2, k - (newStart1 - start1))
	} else {
		return tuple.findK(start1, newStart2, k - (newStart2 - start2))
	}
}
