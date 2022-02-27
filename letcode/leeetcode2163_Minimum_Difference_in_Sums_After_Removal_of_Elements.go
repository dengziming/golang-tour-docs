package letcode

import (
	"container/heap"
	"sort"
)

func minimumDifference(nums []int) int64 {
	min := func(a, b int) int { if a > b { return b }; return a }

	n := len(nums)/3
	minPQ := minHeap{nums[2*n:]}
	heap.Init(&minPQ)
	sum := 0
	for _, v := range nums[2 * n:] {
		sum += v
	}
	sufMax := make([]int, 2*n+1) // 后缀最大和
	sufMax[2*n] = sum

	for i := 2*n - 1; i >= n; i-- {
		if v := nums[i]; v > minPQ.IntSlice[0] {
			sum += v - minPQ.IntSlice[0]
			minPQ.IntSlice[0] = v
			heap.Fix(&minPQ, 0)
		}
		sufMax[i] = sum
	}

	maxPQ := maxHeap{nums[:n]}
	heap.Init(&maxPQ)
	preMin := 0 // 前缀最小和
	for _, v := range nums[:n] {
		preMin += v
	}
	ans := preMin - sufMax[n]
	for i := n; i < 2*n; i++ {
		if v := nums[i]; v < maxPQ.IntSlice[0] {
			preMin += v - maxPQ.IntSlice[0]
			maxPQ.IntSlice[0] = v
			heap.Fix(&maxPQ, 0)
		}
		ans = min(ans, preMin-sufMax[i+1])
	}
	return int64(ans)
}

type minHeap struct{ sort.IntSlice }
func (minHeap) Push(interface{})     {}
func (minHeap) Pop() (_ interface{}) { return }

type maxHeap struct{ sort.IntSlice }
func (h maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (maxHeap) Push(interface{})     {}
func (maxHeap) Pop() (_ interface{}) { return }
