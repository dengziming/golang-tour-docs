package letcode

// https://leetcode-cn.com/problems/minimum-operations-to-make-array-equal/
func minOperations3(n int) int {
	// 最后一定都等于 n，所以直接 用公式求和
	return n * n / 4
}