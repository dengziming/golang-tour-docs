package letcode

import (
	"strconv"
)

func minDistance(word1 string, word2 string) int {
	// s1 添加
	// s2 添加
	// s1 修改

	min := func(a,b,c int) int {
		if a < b {
			if a < c {
				return a
			} else {
				return c
			}
		} else {
			if b < c {
				return b
			} else {
				return c
			}
		}
	}

	key := func(i,j int) string {
		return strconv.Itoa(i) + ":" + strconv.Itoa(j)
	}

	cache := make(map[string]int)

	var divide func(i, j int) int

	// 动态规划写多了，这次换递归
	divide = func(i, j int) int {
		if i == -1 {
			return j+1
		}
		if j == -1 {
			return i+1
		}
		k := key(i,j)
		if _,ok := cache[k]; ok {
			return cache[k]
		}
		// 三种情况
		// 1. s1 添加
		v1 := divide(i-1, j)+1
		v2 := divide(i, j-1)+1
		v3 := divide(i-1, j-1)
		if word1[i]!=word2[j] {
			v3++
		}

		cache[k] = min(v1, v2, v3)
		return cache[k]
	}

	return divide(len(word1)-1, len(word2)-1)
}
