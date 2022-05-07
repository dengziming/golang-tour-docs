package letcode

import "sort"

func generateParenthesis2(n int) []string {
	// 方案1 dfs
	result := make([]string, 0)
	var dfs func(path string, left int, right int)
	dfs = func(path string, left int, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		if left < n {
			dfs(path + "(" , left + 1, right)
		}
		if right < left {
			dfs(path + ")", left, right + 1)
		}

	}
	dfs("", 0 ,0)
	return result
}

func generateParenthesis3(n int) []string {
	// 方案2 dfs
	result := make([]string, 0)
	path := ""
	var backtrack func(left int, right int)
	backtrack = func(left int, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		if left < n {
			path = path + "("
			backtrack(left + 1, right)
			path = path[:len(path) - 1]
		}
		if right < left {
			path = path + ")"
			backtrack(left, right + 1)
			path = path[:len(path) - 1]
		}

	}
	backtrack(0 ,0)
	return result
}

func maxTrailingZeros(grid [][]int) int {
	type tuple struct {
		num2 int
		num5 int
	}

	m := len(grid)
	n := len(grid[0])

	min := func(i int, j int) int {
		if i < j {
			return i
		}
		return j
	}

	cntNum := func(v int, num int) int {
		cnt := 0
		for v % num == 0 {
			cnt += 1
			v /= num
		}
		return cnt
	}

	zeroCnt := func(t1 *tuple, t2 *tuple, v int) int {
		return min(t1.num2 + t2.num2 - cntNum(v,2), t1.num5 + t2.num5-cntNum(v,5))
	}

	// product1 product2 表示横着左右 2 5 的个数
	product1 := make([][]tuple, m)
	product2 := make([][]tuple, m)
	for i := 0; i < m; i++ {
		product1[i] = make([]tuple, n)
		product1[i][0] = tuple{cntNum(grid[i][0] , 2), cntNum(grid[i][0] , 5)}
		for j := 1; j < n; j++ {
			num2 := cntNum(grid[i][j] , 2)
			num5 := cntNum(grid[i][j] , 5)
			product1[i][j] = tuple{product1[i][j-1].num2 + num2, product1[i][j-1].num5 + num5}
		}

		product2[i] = make([]tuple, n)
		product2[i][0] = tuple{product1[i][n-1].num2, product1[i][n-1].num5}
		for j := 1; j < n; j++ {
			num2 := cntNum(grid[i][j-1] , 2)
			num5 := cntNum(grid[i][j-1] , 5)
			product2[i][j] = tuple{product2[i][j-1].num2 - num2, product2[i][j-1].num5 - num5}
		}
	}

	product3 := make([][]tuple, m)
	product4 := make([][]tuple, m)

	for i := 0; i < m; i++ {
		product3[i] = make([]tuple, n)
		product4[i] = make([]tuple, n)
	}

	for i := 0; i < n; i++ {
		product3[0][i] = tuple{cntNum(grid[0][i] , 2), cntNum(grid[0][i] ,5)}
		for j := 1; j < m; j++ {
			num2 := cntNum(grid[j][i] ,2)
			num5 := cntNum(grid[j][i] ,5)
			product3[j][i] = tuple{product3[j-1][i].num2 + num2, product3[j-1][i].num5 + num5}
		}
		product4[0][i] = tuple{product3[m-1][i].num2, product3[m-1][i].num5}
		for j := 1; j < m; j++ {
			num2 := cntNum(grid[j-1][i] ,2)
			num5 := cntNum(grid[j-1][i] ,5)
			product4[j][i] = tuple{product4[j-1][i].num2 - num2, product4[j-1][i].num5 - num5}
		}
	}
	maxV := 0
	i1, i2 := 0, 0
	// 计算以任何一个点作为拐角的情况即可
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var tmp int

			tmp = zeroCnt(&product1[i][j], &product3[i][j], grid[i][j])
			if  tmp > maxV {
				maxV = tmp
				i1 = i
				i2 = j
			}

			tmp = zeroCnt(&product1[i][j], &product4[i][j], grid[i][j])
			if  tmp > maxV {
				maxV = tmp
				i1 = i
				i2 = j
			}

			tmp = zeroCnt(&product2[i][j], &product3[i][j], grid[i][j])
			if  tmp > maxV {
				maxV = tmp
				i1 = i
				i2 = j
			}

			tmp = zeroCnt(&product2[i][j], &product4[i][j], grid[i][j])
			if  tmp > maxV {
				maxV = tmp
				i1 = i
				i2 = j
			}
		}
	}

	println(i1)
	println(i2)
	return maxV
}

func longestPath(parent []int, s string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	n := len(parent)
	childs := make(map[int][]int)
	for i := 1; i < n; i++ {
		childs[parent[i]] = append(childs[parent[i]], i)
	}
	totalMax := 1
	var dfs func(node int) (with, maxV int)
	dfs = func(node int) (with, maxV int) {
		with, maxV = 1, 1
		withi0, withi1 := 0, 0
		for _, child := range childs[node] {
			withi, maxi := dfs(child)
			maxV = max(maxV, maxi)
			if s[child] != s[node] {
				if withi > withi0 {
					withi0, withi1 = withi, withi0
				} else if withi > withi1 {
					withi1 = withi
				}
			}
		}
		with = withi0 + 1
		maxV = max(maxV, withi0+withi1+1)
		totalMax = max(totalMax, maxV)
		return
	}
	dfs(0)
	return totalMax
}


// 由于高度很低，我们可以先按照 高度过滤
func countRectangles(rectangles [][]int, points [][]int) []int {
	n := len(points)
	result := make([]int, n)
	// 按照 L 排序
	sort.Slice(rectangles, func(i, j int) bool {
		// 大于
		return rectangles[i][0] > rectangles[j][0]
	})

	/*var maxL = 0
	for i := 0; i < len(points); i++ {
		if points[i][0] > maxL {
			maxL = points[i][0]
		}
	}*/

	// 记录每个 H : 如果只看高度条件的话，哪些点满足条件
	// dp[i][j] 代表 0 到 j 中的点有多少个 H 小于 i 的
	dp := make([][]int, 101)
	for i := 0; i < 101; i++ {
		dp[i] = make([]int, len(rectangles))
	}

	for i := 0; i < len(rectangles); i++ {
		hi := rectangles[i][1]
		for j := 0; j <= hi; j++ {
			dp[j][i] += 1
		}
	}

	// 二分查找 L 满足的，即大于等于的
	for i := 0; i < n; i++ {
		xi := points[i][0]
		left, right := 0, len(rectangles)-1
		for left < right {
			middle := (left + right)/2
			if rectangles[middle][0] <= xi {
				// 可以继续
				right = middle
			} else {
				left = middle+1
			}
		}
		// 0 到 right 有多少小于 point[j][1]
		result[i] = dp[points[i][1]][right]
	}

	return result
}
